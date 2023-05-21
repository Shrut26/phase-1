package routes

import (
	"github.com/Shrut26/Postmongo/controller"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.GET("/", controller.UserControllers)
}
