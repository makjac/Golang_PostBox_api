package mails

import (
	"fmt"
	"log"

	"github.com/mailjet/mailjet-apiv3-go/v3"
)

func VerficationMail(to string, login string, href string) {
	mailjetClient := MJConnect()
	messagesInfo := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Email: "support@makjac.pl",
				Name:  "Post Box",
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: to,
				},
			},
			TemplateLanguage: true,
			Subject:          "Account activation",
			Variables: map[string]interface{}{
				"login": string(login),
				"href":  string(href),
			},
			TextPart: "Mail Box Hi {{var:login}}! Thanks for setting up the account ;) Click on the link below to activate your account:  {{var:href}} click me",
			HTMLPart: "<center><h1>Mail Box</h1><h3>Hi {{var:login}}! Thanks for setting up the account ;)</h3><h3>Click on the link below to activate your account:</h3><a href=\"{{var:href}}\">click me</a></center>",
		}}
	messages := mailjet.MessagesV31{Info: messagesInfo}
	res, err := mailjetClient.SendMailV31(&messages)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data: %+v", res)
}
