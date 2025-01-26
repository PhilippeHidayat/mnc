package handlers

import (
	"mnc/testApi/db"
	"mnc/testApi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateProfileHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "User profile updated successfully",
	})
}

type RegisterRequest struct {
	FirstName   string `json:"first_name" binding:"required"`
	LastName    string `json:"last_name" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Address     string `json:"address" binding:"required"`
	Pin         string `json:"pin" binding:"required"`
}

func RegisterHandler(c *gin.Context) {
	var req RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var existingUser models.User
	err := db.DB.Model(&existingUser).Where("phone_number = ?", req.PhoneNumber).Select()
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"message": "Phone number already registered"})
		return
	}

	newUser := &models.User{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		PhoneNumber: req.PhoneNumber,
		Address:     req.Address,
		Pin:         req.Pin,
	}
	newUser.BeforeInsert()

	_, err = db.DB.Model(newUser).Insert()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to register user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"result": gin.H{
			"user_id":      newUser.UserID,
			"first_name":   newUser.FirstName,
			"last_name":    newUser.LastName,
			"phone_number": newUser.PhoneNumber,
			"address":      newUser.Address,
			"created_at":   newUser.CreatedAt,
		},
		"status": "SUCCESS",
	})
}
