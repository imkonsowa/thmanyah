package server

import (
	v1 "thmanyah/api/grpc/v1"
	"thmanyah/internal/conf"
	"thmanyah/internal/modules/cms/service"
	discover "thmanyah/internal/modules/discover/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

func NewGRPCServer(
	c *conf.Server,
	authService *service.AuthService,
	cmsService *service.CmsService,
	discoverService *discover.DiscoverService,
	_ log.Logger,
) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterAuthServiceServer(srv, authService)
	v1.RegisterCmsServiceServer(srv, cmsService)
	v1.RegisterDiscoverServiceServer(srv, discoverService)
	return srv
}
