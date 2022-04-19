package database

import (
	"context"

	"github.com/go-pg/pg/v10"
)

var DB *pg.DB

func SetupDB() {
	DB = pg.Connect(&pg.Options{
		Addr:     "35.226.117.80" + ":" + "5432",
		User:     "dbuser",
		Password: "masterkey",
		Database: "pgFlutter",
	})

	ctx := context.Background()

	if err := DB.Ping(ctx); err != nil {
		panic(err)
	}
}
