package controllers

import (
	"blogapi/config"
	"blogapi/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// Register — naya account banao
func Register(c *gin.Context) {
	var input models.RegisterInput

	// Input lo aur check karo
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Sab fields bharo",
		})
		return
	}

	// Password hash karo — plain mat save karo!
	hashed, _ := bcrypt.GenerateFromPassword(
		[]byte(input.Password), bcrypt.DefaultCost,
	)

	// User banao
	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashed),
	}

	// Database mein save karo
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email already exist karta hai",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "✅ Account ban gaya!",
		"name":    user.Name,
		"email":   user.Email,
	})
}

// Login — token lo
func Login(c *gin.Context) {
	var input models.LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email aur password do",
		})
		return
	}

	// User dhundho
	var user models.User
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Galat email ya password",
		})
		return
	}

	// Password check karo
	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password), []byte(input.Password),
	); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Galat email ya password",
		})
		return
	}

	// JWT Token banao
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})

	secret := os.Getenv("JWT_SECRET")
	tokenString, _ := token.SignedString([]byte(secret))

	c.JSON(http.StatusOK, gin.H{
		"message": "✅ Login successful!",
		"token":   tokenString,
	})
}