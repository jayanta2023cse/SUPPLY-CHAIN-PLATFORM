package main

import (
	"net/http"
	"supply_chain_platform/apps/admin/routes"
	"supply_chain_platform/config"
	"supply_chain_platform/logger"
	"supply_chain_platform/middlewares"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	log := logger.GetLogger("admin")

	// Set Gin mode from config
	gin.SetMode(config.AppConfig.GinMode)

	// Create Gin router
	router := gin.Default()
	router.SetTrustedProxies(nil)

	// middlewares
	router.Use(middlewares.ThrottleGuard())

	// Use separated routes
	routes.SetupRoutes(router)

	// Serve Swagger UI at /swagger/*any
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "healthy", "service": "Admin-service", "version": "1.0.0"})
	})

	log.Printf("Starting Admin Service on port %s", config.AppConfig.AdminServicePort)
	if err := router.Run(":" + config.AppConfig.AdminServicePort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
