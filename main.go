package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/alai04/greenlog/api"

	_ "github.com/alai04/greenlog/docs"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8000
func main() {
	r := gin.New()

	r.GET("/ping", api.Pong)
	r.POST("/log", api.LogPost)

	port := ":8000"
	url := ginSwagger.URL(fmt.Sprintf("http://localhost%v/swagger/doc.json", port)) // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.Run(port)
}
