package models

type Parcel_history struct {
	tableName     struct{} `pg:"parcel_history"`
	ID            uint     `json:"ID" pg:"id"`
	Achieved      bool     `json:"achieved" gp:"achieved,type:boolean"`
	Parcel_status string   `json:"parcel_status" pg:"parcel_status"`
	City          string   `json:"city" pg:"city"`
	Point_date    string   `json:"point_date" pg:"point_date"`
	Point_time    string   `json:"point_time" pg:"point_time"`
}
