package models

type PostmachineShowcase struct {
	tableName     struct{} `pg:"post_machine_showcase"`
	Pm_id         int      `json:"pm_id" pg:"pm_id"`
	Pm_city       string   `json:"pm_city" pg:"mp_city"`
	Pm_zipcode    string   `json:"pm_zipcode" pg:"pm_zipcode"`
	Pm_street     string   `json:"pm_street" pg:"pm_street"`
	Pm_buldingnum string   `json:"pm_buldingnum" pg:"pm_buldingnum"`
	Pm_localnum   string   `json:"pm_localnum" pg:"pm_localnum"`
	Number        int      `json:"number" pg:"number"`
	Available     int      `json:"available" pg:"available"`
	Box_type      string   `json:"box_type" pg:"box_type"`
}
