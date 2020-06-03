package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"{{.Appname}}/internal/services"
	"{{.Appname}}/pkg/utils"
)

var userService *services.UserService

func init() {
	userService = services.NewUserService()
}

func GetUserList(ctx *gin.Context) {
	data, err := userService.GetUserList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data":    data,
	})
	return
}

func UserInfo(ctx *gin.Context) {
	id := ctx.Param("id")
	id_ := utils.Str2Uint(id)

	userInfo, err := userService.GetUserInfo(id_)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data":    userInfo,
	})
	return
}

func CreateUser(ctx *gin.Context) {

	username := ctx.Request.FormValue("username")
	password := ctx.Request.FormValue("password")

	err := userService.CreateUser(username, password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data":    "",
	})
	return
}

func UpdateUser(ctx *gin.Context) {

	id := ctx.Request.FormValue("id")
	password := ctx.Request.FormValue("password")
	id_ := utils.Str2Uint(id)

	err := userService.UpdateUserInfo(id_, password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
	})
	return
}

func DeleteUser(ctx *gin.Context) {

	id_ := ctx.Request.FormValue("id")
	id := utils.Str2Uint(id_)
	err := userService.DeleteUser(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
	})
	return
}
