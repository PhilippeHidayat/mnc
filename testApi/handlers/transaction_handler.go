package handlers

import (
	"github.com/gin-gonic/gin"
)

func TopUpHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Top-up successful",
	})
}

func PaymentHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Payment successful",
	})
}

func TransferHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Transfer successful",
	})
}

func TransactionReportHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Transaction report retrieved successfully",
	})
}
