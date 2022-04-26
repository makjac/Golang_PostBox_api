package controller

import (
	"log"
	"net/http"

	"github.com/PBB-api/models"
	"github.com/gin-gonic/gin"
)

func RegisterUser(ctx *gin.Context) {
	var User models.User

	if err := ctx.ShouldBindJSON(&User); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"invalid request": err.Error(),
		})
		ctx.Abort()
		return
	}

	if err := User.HashPasswd(User.Passwd); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}

	//id := guuid.New().String()
}

func GetAllUsers() []models.User {
	var User []models.User
	_, err := dbConnect.QueryOne(&User, "SELECT * FROM users")
	if err != nil {
		log.Println("error")
		return nil
	}
	return User
}
