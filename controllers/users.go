package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Rohanrevanth/chat-demo-go/auth"
	"github.com/Rohanrevanth/chat-demo-go/database"
	"github.com/Rohanrevanth/chat-demo-go/models"

	"github.com/gin-gonic/gin"
)

// GetAlbums responds with the list of all albums as JSON from the database.
func GetAllUsers(c *gin.Context) {
	users, err := database.GetAllUsers()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}
	c.IndentedJSON(http.StatusOK, users)
}

func GetPublicUsers(c *gin.Context) {
	users, err := database.GetPublicUsers()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}
	c.IndentedJSON(http.StatusOK, users)
}

func RegisterUser(c *gin.Context) {
	// var newUser database.User
	var newUser models.User
	if err := c.BindJSON(&newUser); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind user"})
		return
	}

	// Hash the password
	if err := newUser.HashPassword(newUser.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	err := database.SignupUser(newUser)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	user, err := database.GetUserByEmail(newUser.Email)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "User registered successfully", "user": user})
}

// Login authenticates a user and returns a JWT token
func Login(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	// if err := database.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
	// 	return
	// }
	user, err := database.GetUserByEmail(input.Email)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	input_, _ := json.Marshal(input)
	fmt.Println(string(input_))
	user_, _ := json.Marshal(user)
	fmt.Println(string(user_))
	// Check if the password is correct
	if err := user.CheckPassword(input.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Generate JWT token
	token, err := auth.GenerateJWT(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token, "user": user})
}
