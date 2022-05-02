package service

import (
	"log"

	"github.com/PBB-api/models"
	"golang.org/x/crypto/bcrypt"
)

type RegisterService interface {
	Save(models.RegisterForm) models.RegisterForm
	HashPasswd(passwd string) string
	ComparePasswd(hashPasswd string, passwd string) error
}

type registerService struct {
	registerForm models.RegisterForm
}

// func New() RegisterService {
// 	return &registerService{}
// }

// func (service *registerService) Save(form models.RegisterForm) models.RegisterForm {
// 	service.registerForm = append(service.registerForm, form)
// 	return form
// }

func HashPasswd(passwd string) string {
	bytePW := []byte(passwd)
	hashPW, err := bcrypt.GenerateFromPassword(bytePW, bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hash: %v", err)
		return ""
	}
	return string(hashPW)
}

func ComparePasswd(hashPasswd string, passwd string) error {
	PW := []byte(passwd)
	hashPW := []byte(hashPasswd)
	err := bcrypt.CompareHashAndPassword(hashPW, PW)
	return err
}
