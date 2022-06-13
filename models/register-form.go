package models

type RegisterForm struct {
	// auth table
	Login  string `json:"login" validate:"required,unique"`
	Passwd string `json:"password" validate:"required,min=6,max=100"`
	// end auth table / person table
	Name    string `json:"name" validate:"required,excludesall=0x7C"`
	Surname string `json:"surname" validate:"required,excludesall=0x7C"`
	Phone   string `json:"phone" validate:"required,numeric,unique"`
	Email   string `json:"email" validate:"required,email,unique"`
}
