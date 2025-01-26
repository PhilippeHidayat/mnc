package handlers

import (
	"mnc/testApi/db"
	"mnc/testApi/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UpdateProfileRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Address   string `json:"address" binding:"required"`
}

func UpdateProfileHandler(c *gin.Context) {
	var req UpdateProfileRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "User ID not found"})
		return
	}
	var userIDUUID, _ = uuid.Parse(userID.(string))

	user := &models.User{
		UserID:    userIDUUID,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Address:   req.Address,
	}
	user.BeforeUpdate()

	_, err := db.DB.Model(user).Where("user_id = ?", user.UserID).Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update profile"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "SUCCESS",
		"result": gin.H{
			"user_id":      user.UserID,
			"first_name":   user.FirstName,
			"last_name":    user.LastName,
			"address":      user.Address,
			"updated_date": user.UpdatedAt,
		},
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
