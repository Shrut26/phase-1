package routes

import (
	"github.com/Shrut26/gingorm/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.GET("/", controllers.Getusers)
	router.POST("/create", controllers.Createuser)
	router.PUT("/update/:id", controllers.UpdateUser)
	router.DELETE("/delete/:id", controllers.DeleteUser)
}
