package models

type History struct {
	tableName     struct{} `pg:"history"`
	Parcel_uuid   string   `json:"parcel_uuid" pg:"parcel_uuid" binding:"required"`
	Parcel_status string   `json:"parcel_status" pg:"parcel_status"`
	sp_id         int      `json:"sp_id" pg:"sp_id"`
}
