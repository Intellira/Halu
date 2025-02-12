package router

import (
	"github.com/Kelniit/Halu/controller"
	"github.com/gin-gonic/gin"
)

func UserRouter(route *gin.Engine) {
	// User Route
	users := route.Group("/users")
	{
		users.GET("/", controller.GetAllUsers)
		users.GET("/:UID", controller.GetUserID)
		users.POST("/", controller.MoreUsers)
		users.DELETE("/:UID", controller.DeleteUser)
	}
}
