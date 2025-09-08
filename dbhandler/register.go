package dbhandler

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = GetCollection(UserCollection)

func RegisterUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user User
	// bind JSON Data
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// check if user already exists
	var existingUser User
	err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&existingUser)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Benutzer mit dieser E-Mail existiert bereits"})
		return
	}

	err = userCollection.FindOne(ctx, bson.M{"username": user.Username}).Decode(&existingUser)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Benutzername bereits vergeben"})
		return
	}
	// hash password
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Fehler beim Verschlüsseln des Passworts"})
		return
	}
	// prepare new user object
	newUser := User{
		ID:        primitive.NewObjectID(),
		Name:      user.Name,
		Email:     user.Email,
		Phone:     user.Phone,
		Username:  user.Username,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	// safe user in db
	result, err := userCollection.InsertOne(ctx, newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Fehler beim Speichern des Benutzers"})
		return
	}
	// success message
	c.JSON(http.StatusCreated, gin.H{
		"message": "Benutzer erfolgreich registriert",
		"userId":  result.InsertedID,
		"user": gin.H{
			"id":         newUser.ID,
			"name":       newUser.Name,
			"email":      newUser.Email,
			"phone":      newUser.Phone,
			"username":   newUser.Username,
			"created_at": newUser.CreatedAt,
		},
	})
}
