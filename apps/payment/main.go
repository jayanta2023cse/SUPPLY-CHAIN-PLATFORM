package main

import (
	"net/http"
	"supply_chain_platform/apps/payment/routes"
	"supply_chain_platform/config"
	"supply_chain_platform/logger"
	"supply_chain_platform/middlewares"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	log := logger.GetLogger("payment")

	gin.SetMode(config.AppConfig.GinMode)

	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.Use(middlewares.ThrottleGuard())
	routes.SetupRoutes(router)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "healthy", "service": "Payment-service", "version": "1.0.0"})
	})

	log.Printf("Starting Payment Service on port %s", config.AppConfig.PaymentPort)
	if err := router.Run(":" + config.AppConfig.PaymentPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
