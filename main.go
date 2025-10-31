package main

import (
	"example/learnginmongo/configs"
	"example/learnginmongo/routes"

	"github.com/gin-gonic/gin"
)

func main () {
	router := gin.Default()

	configs.ConnectMongoDB()
	routes.UserRoute(router)

	router.Run("localhost:8080")
}