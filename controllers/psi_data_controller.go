package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

// User model
type User struct {
    ID   string `json:"id"`
    Name string `json:"name"`
}

// GetUserHandler retrieves a user by ID
func GetUserHandler(c *gin.Context) {
    userID := c.Param("id")

    user := User{ID: userID, Name: "John Doe"}
    c.JSON(http.StatusOK, user)
}

// CreateUserHandler creates a new user
func CreateUserHandler(c *gin.Context) {
    c.JSON(http.StatusCreated, gin.H{"message": "User created"})
}
