package models

type RegisterForm struct {
	// auth table
	Login         string `json:"login" validate:"required,unique"`
	Passwd        string `json:"password" validate:"required,min=6,max=100"`
	ConfirmPasswd string `json:"confirm_password" validate:"eqfield=Passwd,required"`
	// end auth table / person table
	Name    string `json:"name" validate:"required,excludesall=0x7C"`
	Surname string `json:"surname" validate:"required,excludesall=0x7C"`
	Phone   string `json:"phone" validate:"required,numeric,unique"`
	Email   string `json:"email" validate:"required,email,unique"`
	// end person table / adres table
	Street     string `json:"street"`
	Buldingnum string `json:"buldingnum"`
	Localnum   string `json:"localnum"`
	Zipcode    string `json:"zipcode" validate:"numeric"`
	City       string `json:"city"`
	Country    string `json:"country" validate:"country_code"`
	// end adress table
}
