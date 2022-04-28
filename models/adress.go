package models

type Adress struct {
	tableName  struct{} `pg:"adress"`
	ID         uint     `json:"id" pg:"id,pk"`
	Street     string   `json:"Street" pg:"street" binding:"required"`
	BuldingNum string   `json:"BuldingNum" pg:"buldingnum"`
	LocalNum   string   `json:"LocalNum" pg:"localnum"`
	ZipCode    string   `json:"ZipCode" pg:"zipcode" binding:"required, size=5"`
	City       string   `json:"City" pg:"city" binding:"required"`
	Region     string   `json:"Region" pg:"region"`
	Country    string   `json:"Country" pg:"country"`
}
