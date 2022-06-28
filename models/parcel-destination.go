package models

type ParcelDestination struct {
	tableName     struct{} `pg:"parcel_destination"`
	Parcel_uuid   string   `json:"parcel_uuid" pg:"parcel_uuid"`
	Pm_street     string   `json:"pm_street" pg:"pm_street"`
	Pm_buldingnum string   `json:"pm_buldingnum" pg:"pm_buldingnum"`
	Pm_zipcode    string   `json:"pm_zipcode" pg:"pm_zipcode"`
	Pm_city       string   `json:"pm_city" pg:"pm_city"`
}
