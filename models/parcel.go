package models

type Parcel struct {
	ID     uint   `jason: "id" grom:"primary_key"`
	Name   string `json:"name"`
	Owner  string `json:"owner"`
	Status string `json:"status"`
}
