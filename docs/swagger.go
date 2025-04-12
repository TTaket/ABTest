package docs

import "github.com/swaggo/swag"

var (
	// SwaggerInfo holds exported Swagger Info so clients can modify it
	SwaggerInfo = &swag.Spec{
		Version:          "1.0",
		Title:            "ABTest API",
		Description:      "ABTest服务的HTTP API接口文档",
		Host:             "localhost:8080",
		BasePath:         "/",
		Schemes:          []string{"http"},
	}
)

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}