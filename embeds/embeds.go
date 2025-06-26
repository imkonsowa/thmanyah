package embeds

import (
	_ "embed"

	"thmanyah/third_party/swaggerui"
)

//go:embed multipart_openapi.yaml
var MultipartOpenAPI []byte

//go:embed openapi.yaml
var OpenAPI []byte

var OpenAPIList = map[string][]byte{
	"Thmanyah-API": swaggerui.MergeOpenAPIJSON(OpenAPI, MultipartOpenAPI),
}
