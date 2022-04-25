package models

type Status_point struct {
	tableName        struct{} `pg:"status_point"`
	ID               uint     `json:"id" pg:"id,pk,notnull"`
	Sortung_plant_ID uint     `json:"Sort_id" pg:"sorting_plant_id"`
	Parcel_status_id uint     `json:"parcStat_id" pg:"parcel_status_id"`
	Achieved         bool     `json:"achieved" pg:"achieved,type:boolean"`
	Point_date       string   `json:"point_date" gp:"point_type,type:date"`
	Point_time       string   `json:"point_time" pg:"point_time,type:time"`
}
