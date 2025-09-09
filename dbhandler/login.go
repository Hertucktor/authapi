package dbhandler

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func LoginUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var credentials struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	// bind user data from gin context to credentials struct
	if err := c.BindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userCollection := GetCollection(UserCollection)

	// find user in db
	var user User
	err := userCollection.FindOne(ctx, bson.M{"username": credentials.Username}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User name couldn't be found"})
		return
	}

	// check password
	if !CheckPasswordHash(user.Password, credentials.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Password is not correct"})
		return
	}

	// success message
	c.JSON(http.StatusOK, gin.H{
		"message": "Erfolgreich eingeloggt",
		"user": gin.H{
			"id":       user.ID,
			"name":     user.Name,
			"Email":    user.Email,
			"username": user.Username,
		},
	})
}
