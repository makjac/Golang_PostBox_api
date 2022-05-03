package models

type RegisterForm struct {
	Login    string `json:"login" binding:"required"`
	Passwd   string `json:"password" binding:"required"`
	RePasswd string `json:"repasswd" binding:"required"`
	PlanNum  string `json:"plannum" binding:"required,MinSize(2),MaxSize(3)"`
	Phone    string `json:"phone" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Role     string `json:"role" binding:"default=0"`
}
