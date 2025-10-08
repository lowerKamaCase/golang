package mymail

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/jordan-wright/email"
)

func SendEmail(to string, text string) error {
	e := email.NewEmail()

	email := os.Getenv("EMAIL")

	password := os.Getenv("EMAIL_PASSWORD")

	e.From = email

	fmt.Println("From: ", e.From)

	e.To = []string{to}

	e.Subject = "Golang is awesome!!!"

	e.Text = []byte(text)

	err := e.Send(
		"smtp.yandex.ru:587",
		smtp.PlainAuth("", email, password, "smtp.yandex.ru"),
	)

	return err

}
