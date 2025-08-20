package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("", func(c *gin.Context) {
			c.String(http.StatusOK, "Hello from Gin!")
		})

	}
}
