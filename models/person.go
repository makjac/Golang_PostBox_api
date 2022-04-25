package models

type Person struct {
	tableName struct{} `pg:"person"`
	ID        uint     `json:"id" pg:"id,pk"`
	Name      string   `json:"name" pg:"name"`
	Surname   string   `json:"surname" pg:"surname"`
	Phone     string   `json:"phone" pg:"phone"`
	Email     string   `json:"email" pg:"email"`
}
