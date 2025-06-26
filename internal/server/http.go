package server

import (
	"context"
	http2 "net/http"
	"net/url"
	"os"
	"strings"
	"time"

	v1 "thmanyah/api/grpc/v1"
	"thmanyah/embeds"
	"thmanyah/internal/conf"
	"thmanyah/internal/modules/cms/service"
	discover "thmanyah/internal/modules/discover/service"
	"thmanyah/internal/utils"
	"thmanyah/keys"
	"thmanyah/third_party/swaggerui"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/ratelimit"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport"
	jwt2 "github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/handlers"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

var jsonMarshalOptions = protojson.MarshalOptions{
	EmitUnpopulated: false,
	UseEnumNumbers:  false,
	UseProtoNames:   false,
}

func CustomResponseEncoder(w http.ResponseWriter, r *http.Request, v interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	if pb, ok := v.(proto.Message); ok {
		data, err := jsonMarshalOptions.Marshal(pb)
		if err != nil {
			return err
		}
		_, err = w.Write(data)
		return err
	}
	return nil
}

func JWTMiddleware(keysStore *keys.Store) middleware.Middleware {
	return selector.Server(
		jwt.Server(
			func(token *jwt2.Token) (interface{}, error) {
				claims, ok := token.Claims.(jwt2.MapClaims)
				if !ok {
					return nil, errors.Unauthorized("unauthorized", "Invalid token")
				}

				if claims["user_id"] == "" {
					return nil, errors.Unauthorized("unauthorized", "Invalid token")
				}

				return keysStore.PublicKey(), nil
			},
			jwt.WithSigningMethod(jwt2.SigningMethodRS256),
		),
	).
		Match(JWTWhiteListMatcher()).
		Build()

}

func NewHTTPServer(
	c *conf.Server,
	keysStore *keys.Store,
	authService *service.AuthService,
	cmsservice *service.CmsService,
	discoverService *discover.DiscoverService,
	logger log.Logger,
) *http.Server {
	h := log.NewHelper(logger)

	var opts = []http.ServerOption{
		http.Middleware(
			ratelimit.Server(),
			NewCookieAuthMiddleware(h),
			NewWebLoginMiddleware(
				keysStore,
				WithCookieName("jwt"),
				WithCookieMaxAge(86400),
			),
			recovery.Recovery(),
			validate.Validator(),
			JWTMiddleware(keysStore),
			func(handler middleware.Handler) middleware.Handler {
				return func(ctx context.Context, req any) (any, error) {
					res, err := handler(ctx, req)
					if err != nil {
						return nil, toNetworkError(err, h)
					}

					return res, nil
				}
			},
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}

	opts = append(opts, http.Filter(
		handlers.CORS(
			handlers.AllowCredentials(),
			handlers.AllowedOriginValidator(func(s string) bool {
				if s == "http://localhost:3000" || s == "https://localhost:3000" {
					return true
				}

				if strings.HasSuffix(s, "geeks.local") || strings.HasSuffix(s, "geeks.quest") {
					return true
				}

				return false
			}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
			handlers.AllowedHeaders([]string{
				"Content-Type",
				"Content-Length",
				"Accept-Encoding",
				"X-CSRF-Token",
				"Authorization",
				"accept",
				"origin",
				"Cache-Control",
				"X-Requested-With",
				"X-Platform",
			}),
			handlers.ExposedHeaders([]string{"Content-Length"}),
			handlers.MaxAge(12*3600), // Set preflight cache duration (in seconds)
		),
	))

	opts = append(opts, http.ResponseEncoder(CustomResponseEncoder))
	srv := http.NewServer(
		opts...,
	)

	openAPIHandler := swaggerui.Handler(embeds.OpenAPIList)

	srv.HandlePrefix("/q/", openAPIHandler)

	r := srv.Route("/")

	r.PUT("/api/v1/cms/episodes/upload", func(outerContext http.Context) error {
		h := outerContext.Middleware(func(ctx context.Context, req any) (any, error) {
			userId, err := utils.GetUserID(ctx)
			if err != nil {
				return nil, err
			}

			return authService.UpdateAvatar(outerContext, userId)
		})

		res, err := h(outerContext, nil)
		if err != nil {
			return err
		}

		response, ok := res.(*v1.EpisodeFileUpdateResponse)
		if !ok {
			return errors.New(http2.StatusInternalServerError, "invalid response type", "invalid response type")
		}

		return outerContext.JSON(http2.StatusOK, response)
	})

	v1.RegisterAuthServiceHTTPServer(srv, authService)
	v1.RegisterCmsServiceHTTPServer(srv, cmsservice)
	v1.RegisterDiscoverServiceHTTPServer(srv, discoverService)

	return srv
}

func JWTWhiteListMatcher() selector.MatchFunc {
	whiteList := make(map[string]bool)
	whiteList["/thmanyah.v1.AuthService/Register"] = true
	whiteList["/thmanyah.v1.AuthService/Login"] = true
	whiteList["/thmanyah.v1.DiscoverService/Search"] = true
	whiteList["/thmanyah.v1.DiscoverService/Featured"] = true

	return func(ctx context.Context, operation string) bool {
		return !whiteList[operation]
	}
}

func NewCookieAuthMiddleware(logger *log.Helper) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {

			tr, ok := transport.FromServerContext(ctx)
			if !ok {
				return nil, errors.InternalServer("invalid transport", "invalid transport")
			}

			cookie := tr.RequestHeader().Get("Cookie")
			if cookie == "" {
				return handler(ctx, req)
			}

			cookies := utils.ParseCookies(cookie)
			jwtCookie, exists := cookies["jwt"]
			if !exists {
				return handler(ctx, req)
			}

			tr.RequestHeader().Set("Authorization", "Bearer "+jwtCookie)

			return handler(ctx, req)
		}
	}
}

func NewWebLoginMiddleware(keyStore *keys.Store, opts ...Option) middleware.Middleware {
	options := &Options{
		cookieName:   "jwt",
		cookieMaxAge: 86400,
		isProduction: os.Getenv("ENV") == "production",
		secureCookie: true,
		cookieDomain: "",
		cookiePath:   "/",
	}

	for _, opt := range opts {
		opt(options)
	}

	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			res, err := handler(ctx, req)
			if err != nil {
				return res, err
			}

			switch req.(type) {
			case *v1.LoginRequest:
				tr, ok := transport.FromServerContext(ctx)
				if !ok {
					return res, errors.Unauthorized("invalid transport", "invalid transport")
				}

				response, ok := res.(*v1.LoginResponse)
				if !ok {
					return res, errors.New(http2.StatusUnauthorized, "invalid response type", "invalid response type")
				}

				origin := tr.RequestHeader().Get("origin")
				if origin == "" {
					return res, nil
				}

				claimsMap := utils.NewClaimsBuilder().
					WithUserID(response.User.Id).
					WithExpiry(time.Now().Add(time.Hour * 72).Unix()).
					Build()

				claims := jwt2.NewWithClaims(jwt2.SigningMethodRS256, claimsMap)

				signedString, err := claims.SignedString(keyStore.PrivateKey())
				if err != nil {
					return res, errors.InternalServer("generate token failed", err.Error())
				}

				host := tr.RequestHeader().Get("host")
				if host == "" {
					// Try to parse host from origin
					parsedOrigin, err := url.Parse(origin)
					if err != nil {
						return res, errors.Unauthorized("invalid origin", "invalid origin")
					}
					host = parsedOrigin.Host

				}
				// rootDomain := extractRootDomain(host)

				cookie := &http2.Cookie{
					Name:  options.cookieName,
					Value: signedString,
					// Domain:   rootDomain,
					Path:     options.cookiePath,
					MaxAge:   options.cookieMaxAge,
					HttpOnly: true,
					Secure:   options.isProduction || getSameSiteMode(options.isProduction) == http2.SameSiteNoneMode,
					SameSite: getSameSiteMode(options.isProduction),
				}

				tr.ReplyHeader().Set("Set-Cookie", cookie.String())
				tr.ReplyHeader().Set("X-Content-Type-Options", "nosniff")
				tr.ReplyHeader().Set("X-Frame-Options", "DENY")
				tr.ReplyHeader().Set("Cache-Control", "no-store")

			}

			return res, err
		}
	}
}

type Options struct {
	cookieName   string
	cookieMaxAge int
	isProduction bool
	secureCookie bool
	cookieDomain string
	cookiePath   string
}

type Option func(*Options)

func WithCookieName(name string) Option {
	return func(o *Options) {
		o.cookieName = name
	}
}

func WithCookieMaxAge(maxAge int) Option {
	return func(o *Options) {
		o.cookieMaxAge = maxAge
	}
}

func setSecurityHeaders(hw http.ResponseWriter) {
	hw.Header().Set("X-Content-Type-Options", "nosniff")
	hw.Header().Set("X-Frame-Options", "DENY")
	hw.Header().Set("Cache-Control", "no-store")
}

func getSameSiteMode(isProduction bool) http2.SameSite {
	if isProduction {
		return http2.SameSiteStrictMode
	}
	return http2.SameSiteNoneMode
}

func extractRootDomain(hostname string) string {
	parts := strings.Split(hostname, ".")
	if len(parts) <= 2 {
		// If there are only two or fewer parts, it's likely already the root domain
		return hostname
	}

	// Return the last two parts (e.g., "geeks.quest" from "malik.geeks.quest")
	return strings.Join(parts[len(parts)-2:], ".")
}
