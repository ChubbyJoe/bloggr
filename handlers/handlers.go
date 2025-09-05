package handlers

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ChubbyJoe/bloggr/models"
	"github.com/ChubbyJoe/bloggr/utility"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	_ "github.com/joho/godotenv/autoload"
	"golang.org/x/crypto/bcrypt"
)

func SignUpHandler(c *gin.Context) {
	var input utility.SignUpInput

	err := c.BindJSON(&input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	var existingUser models.User

	result := models.DB.First(&existingUser, "username = ?", input.Username)

	if result.Error == nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": "This user already exists please sign in instead",
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to hash password",
		})
		return
	}

	userToSave := models.User{
		Username: input.Username,
		Password: string(hashedPassword),
	}

	err = models.DB.Create(&userToSave).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to created user",
		})
		return
	}

	response := utility.UserResponse{
		ID:        userToSave.ID,
		Username:  userToSave.Username,
		Blogs:     userToSave.Blogs,
		CreatedAt: userToSave.CreatedAt,
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "user created successfully",
		"message": response,
	})
}

func LoginHandler(c *gin.Context) {
	var input utility.SignUpInput

	err := c.BindJSON(&input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var user models.User

	results := models.DB.First(&user, "username = ?", input.Username)

	if results.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid credentials",
		})
		return
	}

	claims := jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtSecret := os.Getenv("JWT_SECRET")

	if jwtSecret == "" {
		log.Println("jwt secret not set")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server error",
		})
		return
	}

	tokenString, err := token.SignedString([]byte(jwtSecret))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
