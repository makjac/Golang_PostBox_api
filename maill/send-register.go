package mails

import (
	"fmt"
	"log"

	"github.com/mailjet/mailjet-apiv3-go/v3"
)

func SendRegisterActive(to string, login string, href string) {
	mj := MJConnect()
	messagesInfo := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Email: "support@makjac.pl",
				Name:  "Post Box",
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: to,
					Name:  login,
				},
			},
			Variables: map[string]interface{}{
				"login": login,
				"href":  href},
			TemplateID:       3905246,
			TemplateLanguage: true,
			Subject:          "Account activation Post Box",
		},
	}
	messages := mailjet.MessagesV31{Info: messagesInfo}
	res, err := mj.SendMailV31(&messages)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data: %+v\n", res)
}
