package swaggerui

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gorilla/mux"
)

//go:embed embed
var swagFs embed.FS

type ConfigUrl struct {
	Url  string `json:"url"`
	Name string `json:"name"`
}

func byteHandler(b []byte) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write(b)
	}
}

func configHandler(config map[string][]byte) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {

		var conf []ConfigUrl

		for k := range config {
			conf = append(conf, ConfigUrl{Url: "/q/service/" + k, Name: k})
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		err := json.NewEncoder(w).Encode(conf)
		if err != nil {
			return
		}
	}
}

func Handler(config map[string][]byte) http.Handler {
	static, _ := fs.Sub(swagFs, "embed")
	router := mux.NewRouter()
	router.HandleFunc("/q/services", configHandler(config))

	for k, file := range config {
		router.HandleFunc("/q/service/"+k, byteHandler(file))
	}
	sh := http.StripPrefix("/q/swagger-ui", http.FileServer(http.FS(static)))
	router.PathPrefix("/q/swagger-ui").Handler(sh)
	return router
}

func MergeOpenAPIJSON(files ...[]byte) []byte {
	if len(files) < 2 {
		if len(files) < 1 {
			return nil
		}
		return files[0]
	}

	// Parse the first OpenAPI JSON string
	swagger1, err := openapi3.NewLoader().LoadFromData(files[0])
	if err != nil {
		panic(fmt.Errorf("failed to parse first JSON: %v", err))
	}

	for i := 1; i < len(files); i++ {

		// Parse the second OpenAPI JSON string
		swagger2, err := openapi3.NewLoader().LoadFromData(files[i])
		if err != nil {
			panic(fmt.Errorf("failed to parse second JSON: %v", err))
		}

		if swagger2.Paths != nil {
			for name, pathItem := range swagger2.Paths.Map() {
				swagger1.Paths.Set(name, pathItem)
			}
		}

		if len(swagger2.Components.Schemas) > 0 {
			for name, pathItem := range swagger2.Components.Schemas {
				swagger1.Components.Schemas[name] = pathItem
			}
		}
	}

	jsString, err := swagger1.MarshalJSON()
	if err != nil {
		panic(fmt.Errorf("failed to marshal JSON: %v", err))
	}
	return jsString
}
