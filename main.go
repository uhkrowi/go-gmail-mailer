package main

import (
	"log"
	mailer "mailer/mail"
	"time"
)

func main() {
	mail := mailer.NewSMTPMailer()

	err := mail.Send("sample_receipent@xemaildomain.com", "Test", "./mail/sample.template", nil)
	if err != nil {
		log.Println(err)
	}

	time.Sleep(30 * time.Second)
}
