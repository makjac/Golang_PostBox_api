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

func GetAllUsers(c *gin.Context) {
	var User []models.User
	err := dbConnect.Model(&User).Select()
	if err != nil {
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"staus":   http.StatusOK,
		"message": "all users",
		"data":    User,
	})
	return
}

func FindUserByID(uuid string) (bool, error) {
	var User []models.User
	Exists, err := dbConnect.Model(&User).Where("uuid = ?", uuid).Exists()
	if err != nil {
		log.Println(err)
		return Exists, err
	}
	return Exists, nil
}

func FindUserByLogin(login string) (bool, error) {
	var User []models.User
	Exists, err := dbConnect.Model(&User).Where("username = ?", login).Exists()
	if err != nil {
		log.Println(err)
		return false, err
	}
	if Exists {
		return true, nil
	} else {
		return false, &LoginNotExist{}
	}
}

func CheckUser(username string, password string) (bool, error) {
	var User []models.User
	Exists, err := FindUserByLogin(username)
	if err != nil {
		log.Printf("Login error, user " + username + " does not exist")
		return Exists, err
	}

	//err := dbConnect.Model(&User).Where("username = ?", username).Select()

	if Exists {
		err := dbConnect.Model(&User).Where("username = ?", username).Select()
		if err != nil {
			log.Println(err)
			return false, err
		}

		//err := bcrypt.CompareHashAndPassword([]byte(User["Passwd"]), []byte(password))
	}
	return false, nil
}

type LoginNotExist struct{}

func (m *LoginNotExist) Error() string {
	return "The user with this name does not exist"
}
