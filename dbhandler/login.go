package dbhandler

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type credentials struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type loginUser struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name" binding:"required"`
	Email     string             `json:"email" bson:"email" binding:"required"`
	Phone     *string            `json:"phone,omitempty" bson:"phone,omitempty"`
	Username  string             `json:"username" bson:"username" binding:"required,min=1,max=30"`
	Password  string             `json:"password" bson:"password" binding:"required,min=15"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

func LoginUser(c *gin.Context) {
	var credentials credentials
	var loginUser loginUser

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// bind user data from gin context to credentials struct
	if err := c.BindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userCollection := GetCollection(UserCollection)

	// find user in db
	err := userCollection.FindOne(ctx, bson.M{"username": credentials.Username}).Decode(&loginUser)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User name couldn't be found"})
		return
	}

	// check password
	if !CheckPasswordHash(loginUser.Password, credentials.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Password is not correct"})
		return
	}

	// success message
	c.JSON(http.StatusOK, gin.H{
		"message": "Erfolgreich eingeloggt",
		"user": gin.H{
			"id":       loginUser.ID,
			"name":     loginUser.Name,
			"Email":    loginUser.Email,
			"username": loginUser.Username,
		},
	})
}
