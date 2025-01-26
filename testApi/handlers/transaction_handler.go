package handlers

import (
	"fmt"
	"mnc/testApi/db"
	"mnc/testApi/models"
	"mnc/testApi/scripts"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
)

type TopUpRequest struct {
	Amount float64 `json:"amount" binding:"required"`
}

type PaymentRequest struct {
	Amount  float64 `json:"amount" binding:"required"`
	Remarks string  `json:"remarks" binding:"required"`
}

type TransferRequest struct {
	TargetUser uuid.UUID `json:"target_user" binding:"required"`
	Amount     float64   `json:"amount" binding:"required"`
	Remarks    string    `json:"remarks" binding:"required"`
}

func TopUpHandler(c *gin.Context) {
	var req TopUpRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "User ID not found"})
		return
	}

	var lastTransaction models.Transaction
	err := db.DB.Model(&lastTransaction).Where("user_id = ?", userID).Order("created_date DESC").Limit(1).Select()

	if err != nil {
		lastTransaction.BalanceAfter = 0.00
	}
	newBalance := lastTransaction.BalanceAfter + req.Amount

	var userIDUUID, _ = uuid.Parse(userID.(string))

	newTransaction := &models.Transaction{
		UserID:          userIDUUID,
		Status:          "SUCCESS",
		TransactionType: "CREDIT",
		Amount:          req.Amount,
		Remarks:         "",
		BalanceBefore:   lastTransaction.BalanceAfter,
		BalanceAfter:    newBalance,
	}
	newTransaction.BeforeInsert()
	fmt.Println(newTransaction)
	go func(db *pg.DB, newTransaction *models.Transaction) {
		err := scripts.InsertTransaction(db, newTransaction)
		if err != nil {
			fmt.Printf("Failed to insert transaction: %v", err)
		}
	}(db.DB, newTransaction)

	c.JSON(http.StatusCreated, gin.H{
		"status": "SUCCESS",
		"result": gin.H{
			"transaction_id": newTransaction.TransactionID,
			"amount_top_up":  newTransaction.Amount,
			"balance_before": newTransaction.BalanceBefore,
			"balance_after":  newTransaction.BalanceAfter,
			"created_date":   newTransaction.CreatedDate,
		},
	})
}

func PaymentHandler(c *gin.Context) {
	var req PaymentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "User ID not found"})
		return
	}

	var lastTransaction models.Transaction
	err := db.DB.Model(&lastTransaction).Where("user_id = ?", userID).Order("created_date DESC").Limit(1).Select()
	if err != nil {
		lastTransaction.BalanceAfter = 0
	}

	if lastTransaction.BalanceAfter < req.Amount {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Balance is not enough"})
		return
	}

	newBalance := lastTransaction.BalanceAfter - req.Amount
	var userIDUUID, _ = uuid.Parse(userID.(string))
	newTransaction := &models.Transaction{
		UserID:          userIDUUID,
		Status:          "SUCCESS",
		TransactionType: "DEBIT",
		Amount:          req.Amount,
		Remarks:         req.Remarks,
		BalanceBefore:   lastTransaction.BalanceAfter,
		BalanceAfter:    newBalance,
	}
	newTransaction.BeforeInsert()

	go func(db *pg.DB, newTransaction *models.Transaction) {
		err := scripts.InsertTransaction(db, newTransaction)
		if err != nil {
			fmt.Printf("Failed to insert payment transaction: %v", err)
		}
	}(db.DB, newTransaction)

	c.JSON(http.StatusCreated, gin.H{
		"status": "SUCCESS",
		"result": gin.H{
			"transaction_id": newTransaction.TransactionID,
			"amount":         newTransaction.Amount,
			"remarks":        newTransaction.Remarks,
			"balance_before": newTransaction.BalanceBefore,
			"balance_after":  newTransaction.BalanceAfter,
			"created_date":   newTransaction.CreatedDate,
		},
	})
}

func TransferHandler(c *gin.Context) {
	var req TransferRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "User ID not found"})
		return
	}

	var lastTransaction models.Transaction
	err := db.DB.Model(&lastTransaction).Where("user_id = ?", userID).Order("created_date DESC").Limit(1).Select()
	if err != nil {
		lastTransaction.BalanceAfter = 0
	}

	if lastTransaction.BalanceAfter < req.Amount {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Balance is not enough"})
		return
	}

	newBalance := lastTransaction.BalanceAfter - req.Amount

	var userIDUUID, _ = uuid.Parse(userID.(string))

	sourceTransaction := &models.Transaction{
		UserID:          userIDUUID,
		Status:          "SUCCESS",
		TransactionType: "DEBIT",
		Amount:          req.Amount,
		Remarks:         req.Remarks,
		BalanceBefore:   lastTransaction.BalanceAfter,
		BalanceAfter:    newBalance,
	}
	sourceTransaction.BeforeInsert()

	var lastTargetTransaction models.Transaction
	err = db.DB.Model(&lastTargetTransaction).Where("user_id = ?", req.TargetUser).Order("created_date DESC").Limit(1).Select()
	if err != nil {
		fmt.Println("No prior transaction exists")
		lastTargetTransaction.BalanceAfter = 0
	}
	var targetUserIDUUID, _ = uuid.Parse(req.TargetUser.String())
	targetTransaction := &models.Transaction{
		UserID:          targetUserIDUUID,
		Status:          "SUCCESS",
		TransactionType: "CREDIT",
		Amount:          req.Amount,
		Remarks:         req.Remarks,
		BalanceBefore:   lastTargetTransaction.BalanceAfter,
		BalanceAfter:    lastTargetTransaction.BalanceAfter + req.Amount,
	}
	targetTransaction.BeforeInsert()
	fmt.Println(targetTransaction)

	go func(db *pg.DB, sourceTransaction *models.Transaction, targetTransaction *models.Transaction) {
		err := scripts.InsertTransaction(db, sourceTransaction)
		if err != nil {
			fmt.Printf("Failed to insert source transaction: %v", err)
		} else {
			fmt.Println("Source transaction inserted successfully")
			err = scripts.InsertTransaction(db, targetTransaction)
			if err != nil {
				fmt.Printf("Failed to insert target transaction: %v", err)
			}
		}
	}(db.DB, sourceTransaction, targetTransaction)

	c.JSON(http.StatusCreated, gin.H{
		"status": "SUCCESS",
		"result": gin.H{
			"transaction_id": sourceTransaction.TransactionID,
			"amount":         sourceTransaction.Amount,
			"remarks":        sourceTransaction.Remarks,
			"balance_before": sourceTransaction.BalanceBefore,
			"balance_after":  sourceTransaction.BalanceAfter,
			"created_date":   sourceTransaction.CreatedDate,
		},
	})
}

func TransactionReportHandler(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "User ID not found"})
		return
	}

	var transactions []models.Transaction
	err := db.DB.Model(&transactions).
		Where("user_id = ?", userID).
		Order("created_date ASC").
		Select()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve transactions"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "SUCCESS",
		"result": transactions,
	})
}
