package routes

import (
	"example/learnginmongo/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.POST("/user", controllers.CreateUser())
	router.GET("/user/:id", controllers.GetUser())
}
