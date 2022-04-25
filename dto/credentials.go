package dto

type Credentials struct {
	Username string `form:"username"`
	Passwd   string `form:"passwd"`
}
