package models

type User struct {
	tableName struct{} `pg:"auth"`
	Username  string   `json:"username" pg:"username,notnull,unique" validate:"required, max=30"`
	Passwd    string   `json:"passwd" pg:"passwd,notnull" validate:"required, min=6, max=100"`
	Uuid      string   `json:"UUID" pg:"uuid,notnull,unique" validate:"uuid4"`
}

type Auth struct {
	tableName struct{} `pg:"auth"`
	Username  string   `json:"username" pg:"username"`
	Active    bool     `json:"active" pg:"active"`
}
