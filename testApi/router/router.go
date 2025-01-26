package router

import (
	"mnc/testApi/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hi",
		})
	})

	router.POST("/register", handlers.RegisterHandler)
	router.POST("/login", handlers.LoginHandler)

	// Protected routes
	protected := router.Group("/")
	protected.Use(handlers.AuthMiddleware())
	{
		protected.POST("/topup", handlers.TopUpHandler)
		protected.POST("/payment", handlers.PaymentHandler)
		protected.POST("/transfer", handlers.TransferHandler)
		protected.GET("/transaction-report", handlers.TransactionReportHandler)
		protected.PUT("/profile", handlers.UpdateProfileHandler)
	}

	return router
}
