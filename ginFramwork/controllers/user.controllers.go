package controllers

import (
	"net/http"

	"github.com/Shrut26/ginFramework/models"
	"github.com/Shrut26/ginFramework/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService services.Userservices
}

func New(userservice *services.Userservices) UserController {
	return UserController{
		UserService: *userservice,
	}
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": "Invalid User",
			"error":   err.Error(),
		})
		return
	}

	res, err := uc.UserService.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": "Unable to create user",
			"error":   err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"User":    user,
		"ID":      res.InsertedID,
	})
}

func (uc *UserController) GetUser(ctx *gin.Context) {
	username := ctx.Param("name")
	result, err := uc.UserService.GetUser(&username)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": "Unable to get the user",
			"error":   err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Successfully retrieved the user",
		"User":    result,
	})
}

func (uc *UserController) GetAll(ctx *gin.Context) {
	var users []*models.User
	users, err := uc.UserService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": "Unable to get the users",
			"error":   err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Successful retrived all",
		"Users":   users,
	})
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
	}

	err := uc.UserService.UpdateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": "Failed to update the user",
			"error":   err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":      "Successfully updated user",
		"Updated user": user,
	})
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {
	username := ctx.Param("name")
	err := uc.UserService.DeleteUser(&username)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Unable to delete the user",
			"error":   err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Successfully delete account",
	})
}

func (uc *UserController) RegisterUserRoutes(rg *gin.RouterGroup) {
	userRoute := rg.Group("/user")
	userRoute.POST("/create", uc.CreateUser)
	userRoute.GET("/get/:name", uc.GetUser)
	userRoute.GET("/getAll", uc.GetAll)
	userRoute.PATCH("/update", uc.UpdateUser)
	userRoute.DELETE("/delete/:name", uc.DeleteUser)
}
