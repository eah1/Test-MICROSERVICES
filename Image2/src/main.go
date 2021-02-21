package main

import (
	_ "Test-MICROSERVICES/Image2/src/docs"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)

// @title User API
// @description User microservice server.
// @schemes http https
// @securityDefinitions.basic BasicAuth
func main() {

	r := gin.Default()

	// ListUser is the handler of list user endpoint
	// @Summary Health
	// @Description Endpoint health
	// @Tags health
	// @Produce  json
	// @Success 200 type {string}
	// @Router /health [get]
	r.GET("/health", checkHandler())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8020")
}

// CheckHandler returns an HTTP handler to perform health checks.
// ListUser is the handler of list user endpoint
// @Summary Health
// @Description Endpoint health
// @Tags health
// @Produce  json
// @Success 200 type {string}
// @Router /health [get]
func checkHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "Ok!",
		})
	}
}
