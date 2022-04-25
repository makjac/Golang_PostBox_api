package controller

import "github.com/go-pg/pg/v10"

var dbConnect *pg.DB

func InitiazeDB(db *pg.DB) {
	dbConnect = db
}
