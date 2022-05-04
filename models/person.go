package models

type Person struct {
	tableName struct{} `pg:"person"`
	//ID        uint     `json:"id" pg:"id,pk"`
	Name    string `json:"name" pg:"name,notnull"`
	Surname string `json:"surname" pg:"surname,notnull"`
	Phone   string `json:"phone" pg:"phone,notnull"`
	Email   string `json:"email" pg:"email,notnull,unique"`
}
