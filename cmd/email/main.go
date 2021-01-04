package main

import (
	"fmt"
	"log"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func main() {
	from := mail.NewEmail("ToDanni", "admin@todanni.com")
	subject := "Thank you for signing up for ToDanni!"
	to := mail.NewEmail("Danni Popova", "dannii.popova@gmail.com")
	plainTextContent := "Please verify your account by clicking http://www.todanni.com"
	htmlContent := "<strong>Please verify your account by clicking https://www.todanni.com</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient("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")
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(response.StatusCode)
	fmt.Println(response.Body)
	fmt.Println(response.Headers)
}
