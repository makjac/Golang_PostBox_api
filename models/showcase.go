package models

import "time"

type Showcase struct {
	tableName struct{}  `pg:"showcase"`
	User_id   int       `json:"user_id" pg:"user_id"`
	User_uuid string    `json:"user_uuid" pg:"user_uuid"`
	Username  string    `json:"username" pg:"username"`
	Name      string    `json:"name" pg:"name"`
	Surname   string    `json:"surname" pg:"surname"`
	Email     string    `json:"email" pg:"email"`
	Phone     string    `json:"phone" pg:"phone"`
	Created   time.Time `json:"created" pg:"created"`
}
