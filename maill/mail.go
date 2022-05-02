package mails

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mailjet/mailjet-apiv3-go/v3"
)

func MJConnect() *mailjet.Client {
	err := godotenv.Load(".mail.env")
	if err != nil {
		log.Fatal("Error loading .mail.env file")
	}

	publicKey := os.Getenv("MJ_PUBK")
	secretKey := os.Getenv("MJ_SECK")

	mjClient := mailjet.NewMailjetClient(publicKey, secretKey)

	return mjClient
}
