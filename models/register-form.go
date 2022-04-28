package models

type RegisterForm struct {
	Login      string `json:"login" binding:"required"`
	Passwd     string `json:"password" binding:"required"`
	RePasswd   string `json:"repasswd" binding:"required"`
	PlanNum    string `json:"plannum" binding:"required,MinSize(2),MaxSize(3)"`
	Phone      string `json:"phone" binding:"required"`
	Email      string `json:"email" binding:"required"`
	Street     string `json:"street" binding:"required"`
	BuldingNum string `json:"buldingnum"`
	LocalNum   string `json:"localnum"`
	ZipCode    string `json:"zipcode" binding:"required,size(5)"`
	City       string `json:"city" binding:"required"`
}
