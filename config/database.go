package database

import (
	"context"
	"log"
	"os"

	"github.com/PBB-api/controller"
	"github.com/go-pg/pg/v10"
	"github.com/joho/godotenv"
)

var db *pg.DB

func SetupDB() {

	//Loading .env file
	err := godotenv.Load(".database.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")

	//Connecting to the postgresql server database
	db = pg.Connect(&pg.Options{
		Addr:     dbHost + ":" + dbPort,
		User:     dbUser,
		Password: dbPass,
		Database: dbName,
	})

	ctx := context.Background()

	if err := db.Ping(ctx); err != nil {
		panic(err)
	}

	controller.InitiazeDB(db)
}
