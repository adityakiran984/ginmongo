package controllers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"example/learnginmongo/configs"
	"example/learnginmongo/models"
	"example/learnginmongo/responses"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var collection *mongo.Collection = configs.GetCollections(configs.Client, os.Getenv("COLLECTION_NAME"))
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
			Id:       bson.NewObjectID(),
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

		objId, err := bson.ObjectIDFromHex(id)
		if err != nil {
			fmt.Println("in ObjectIDFromHex error block")
			c.IndentedJSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "objId for requested user is not valid", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		filter := bson.M{"_id": objId}

		findErr := collection.FindOne(ctx, filter).Decode(&user)
		// if findErr != nil {
		// 	fmt.Println(user)
		// 	fmt.Println("in FindOne error block")
		// 	c.IndentedJSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "failure", Data: map[string]interface{}{"data": findErr.Error()}})
		// 	// c.IndentedJSON(http.StatusInternalServerError, gin.H{"data" : })
		// 	return
		// }
		if errors.Is(findErr, mongo.ErrNoDocuments) {
			c.JSON(http.StatusNotFound, gin.H{"message": "User with the given id does not exist"})
			return
		}

		fmt.Println(user)
		if findErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": findErr.Error()})
			return
		}

		c.IndentedJSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": user}})
	}
}
