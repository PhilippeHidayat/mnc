package handlers

import (
	"mnc/testApi/db"
	"mnc/testApi/models"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type LoginRequest struct {
	PhoneNumber string `json:"phone_number" binding:"required"`
	Pin         string `json:"pin" binding:"required"`
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func LoginHandler(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	err := db.DB.Model(&user).Where("phone_number = ? AND pin = ?", req.PhoneNumber, req.Pin).Select()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid phone number or pin"})
		return
	}

	accessToken, refreshToken, err := generateTokens(user.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate tokens"})
		return
	}

	// Return the tokens
	c.JSON(http.StatusOK, TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

func generateTokens(userID uuid.UUID) (string, string, error) {

	accessTokenExpiry := time.Now().Add(15 * time.Minute)    // 15 minutes
	refreshTokenExpiry := time.Now().Add(7 * 24 * time.Hour) // 7 days

	accessTokenClaims := jwt.MapClaims{
		"user_id":    userID,
		"token_type": "access",
		"exp":        accessTokenExpiry.Unix(),
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenString, err := accessToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", "", err
	}

	refreshTokenClaims := jwt.MapClaims{
		"user_id":    userID,
		"token_type": "refresh",
		"exp":        refreshTokenExpiry.Unix(),
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}
