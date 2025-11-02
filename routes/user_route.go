package routes

import (
	"github.com/gin-gonic/gin"

	"example/learnginmongo/controllers"
)

func UserRoute(router *gin.Engine) {
	router.POST("/user", controllers.CreateUser())
	router.GET("/user/:id", controllers.GetUser())
	router.PUT("/user/:id", controllers.UpdateUser())
	router.DELETE("/user/:id", controllers.DeleteUser())
}
