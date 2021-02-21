package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "Ok!",
		})
	})

	r.Run(":8020") // listen and serve on 0.0.0.0:8080
}
