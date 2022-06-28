package models

type ParcelShowcase struct {
	tableName     struct{} `pg:"parcel_showcase"`
	Parcel_num    string   `json:"parcel_num" pg:"parcel_num"`
	Parcel_name   string   `json:"parcel_name" pg:"parcel_name"`
	Sender        string   `json:"sender" pg:"sender"`
	Sender_name   string   `json:"sender_uuid" pg:"sender_uuid"`
	Receiver      string   `json:"receiver" pg:"receiver"`
	Receiver_uuid string   `json:"receiver_uuid" pg:"receiver_uuid"`
}
