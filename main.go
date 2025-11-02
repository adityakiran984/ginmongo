package main

import (
	"example/learnginmongo/configs"
	"example/learnginmongo/routes"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println(".env file not found.")
	}
	
	router := gin.Default()

	configs.ConnectMongoDB()
	routes.UserRoute(router)

	router.Run("localhost:8080")
}
