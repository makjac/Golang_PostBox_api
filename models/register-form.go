package models

type RegisterForm struct {
	tableName struct{} `pg:"register_form"`
	// auth table
	Login  string `json:"login" pg:"_username" validate:"required,unique"`
	Passwd string `json:"password" pg:"_passwd" validate:"required,min=6,max=100"`
	// end auth table / person table
	Name    string `json:"name" pg:"_name" validate:"required,excludesall=0x7C"`
	Surname string `json:"surname" pg:"_surname" validate:"required,excludesall=0x7C"`
	Phone   string `json:"phone" pg:"_phone" validate:"required,numeric,unique"`
	Email   string `json:"email" pg:"_email" validate:"required,email,unique"`
}
