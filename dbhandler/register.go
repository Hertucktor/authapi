package dbhandler

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type registerUserRequest struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name" binding:"required"`
	Email     string             `json:"email" bson:"email" binding:"required"`
	Phone     *string            `json:"phone,omitempty" bson:"phone,omitempty"`
	Username  string             `json:"username" bson:"username" binding:"required,min=1,max=30"`
	Password  string             `json:"password" bson:"password" binding:"required,min=15"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

type user struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name" binding:"required"`
	Email     string             `json:"email" bson:"email" binding:"required"`
	Phone     *string            `json:"phone,omitempty" bson:"phone,omitempty"`
	Username  string             `json:"username" bson:"username" binding:"required,min=1,max=30"`
	Password  string             `json:"password" bson:"password" binding:"required,min=15"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

func RegisterUser(c *gin.Context) {
	var reqisterUser registerUserRequest

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// bind JSON
	if err := c.BindJSON(&reqisterUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userCollection := GetCollection(UserCollection)

	// check if user already exists
	var existingUser user
	err := userCollection.FindOne(ctx, bson.M{"email": reqisterUser.Email}).Decode(&existingUser)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Benutzer mit dieser E-Mail existiert bereits"})
		return
	}

	err = userCollection.FindOne(ctx, bson.M{"username": reqisterUser.Username}).Decode(&existingUser)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Benutzername bereits vergeben"})
		return
	}
	// hash password
	hashedPassword, err := HashPassword(reqisterUser.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Fehler beim Verschlüsseln des Passworts"})
		return
	}
	// prepare new user object
	newUser := user{
		ID:        primitive.NewObjectID(),
		Name:      reqisterUser.Name,
		Email:     reqisterUser.Email,
		Phone:     reqisterUser.Phone,
		Username:  reqisterUser.Username,
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
