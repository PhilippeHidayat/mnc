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
	router.POST("/topup", handlers.TopUpHandler)
	router.POST("/payment", handlers.PaymentHandler)
	router.POST("/transfer", handlers.TransferHandler)
	router.GET("/transaction-report", handlers.TransactionReportHandler)
	router.PUT("/update-profile", handlers.UpdateProfileHandler)

	return router
}
