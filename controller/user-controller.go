package controller

import (
	"log"

	"github.com/PBB-api/models"
	"golang.org/x/crypto/bcrypt"
)

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

func CheckUser(username string, passwd string) (bool, string, error) {
	var User models.User

	err := dbConnect.Model(&User).Where("username = ?", username).Select()
	if err != nil {
		log.Println(err)
		return false, "", &LoginOrPasswdErr{}
	}

	err = bcrypt.CompareHashAndPassword([]byte(User.Passwd), []byte(passwd))
	if err != nil {
		log.Println(err)
		return false, "", &IncorrectPasswd{}
	}

	return true, User.Uuid, nil
}

type LoginNotExist struct{}
type IncorrectPasswd struct{}
type LoginOrPasswdErr struct{}

func (m *LoginNotExist) Error() string {
	return "The user with this name does not exist"
}

func (m *IncorrectPasswd) Error() string {
	return "Password is not correct"
}

func (m *LoginOrPasswdErr) Error() string {
	return "Login or password is not correct"
}
