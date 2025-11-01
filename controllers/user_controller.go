package controllers

import (
	"context"
	"example/learnginmongo/configs"
	"example/learnginmongo/models"
	"example/learnginmongo/responses"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var collection *mongo.Collection = configs.GetCollections(configs.Client, "test")
var validate *validator.Validate = validator.New()

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var user models.User
		defer cancel()

		if err := c.BindJSON(&user); err != nil {
			c.IndentedJSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "failure", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if validationError := validate.Struct(&user); validationError != nil {
			c.IndentedJSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "failure", Data: map[string]interface{}{"data": validationError.Error()}})
			return
		}

		newUser := models.User{
			Id:       primitive.NewObjectID(),
			Name:     user.Name,
			Location: user.Location,
			Title:    user.Title,
		}
		result, err := collection.InsertOne(ctx, newUser)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "failure", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.IndentedJSON(http.StatusOK, responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		id := c.Param("id")
		var user models.User

		userId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "failure", Data: map[string]interface{}{"data" : err.Error()}})
		}

		filter := bson.D{{Key: "_id", Value: userId}}
		
		findErr := collection.FindOne(ctx, filter).Decode(&user)
		if err!= nil {
			c.IndentedJSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "failure", Data: map[string]interface{}{"data" : findErr.Error()}})
			return
		}

		c.IndentedJSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data" : user}})
	}
}
