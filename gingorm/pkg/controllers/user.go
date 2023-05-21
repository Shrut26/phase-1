package controllers

import (
	"net/http"

	"github.com/Shrut26/gingorm/config"
	"github.com/Shrut26/gingorm/models"
	"github.com/gin-gonic/gin"
)

func Getusers(ctx *gin.Context) {
	users := []models.User{}
	config.DB.Find(&users)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"users":   users,
	})
}

func Createuser(ctx *gin.Context) {
	var user models.User
	ctx.ShouldBindJSON(&user)
	err := config.DB.Create(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": "Unable to Create user",
			"err":     err,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"user":    user,
	})
}

func UpdateUser(ctx *gin.Context) {
	var user models.User
	id := ctx.Param("id")
	config.DB.Where("id=?", id).First(&user)
	ctx.BindJSON(&user)
	res := config.DB.Save(&user)
	if res.RowsAffected == 0 {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": "Unable to update user",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"user":    user,
	})
}

func DeleteUser(ctx *gin.Context) {
	var user models.User
	id := ctx.Param("id")
	res := config.DB.Where("id=?", id).Delete(&user)

	if res.RowsAffected == 0 {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": "Unable to delete",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success deleted",
	})
}
