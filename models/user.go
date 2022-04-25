package models

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	tableName struct{} `pg:"auth"`
	ID        string   `json:"id" pg:"id"`
	Username  string   `json:"username" pg:"username" validate:"required, max=30"`
	Passwd    string   `json:"passwd" pg:"passwd" validate:"required, min=6, max=100"`
}

func (user *User) HashPasswd(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		log.Printf("Error during hashing password: %v", err)
		return err
	}
	user.Passwd = string(bytes)
	return nil
}

func (user *User) CheckPasswd(providedPasswd string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Passwd), []byte(providedPasswd))
	if err != nil {
		log.Printf("Error during checking the password: %v", err)
		return err
	}
	return nil
}
