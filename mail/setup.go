package mailer

import (
	"fmt"
	"net/smtp"
	"sync"
)

type SMTPMailSetup struct {
	Server   string
	Port     string
	Address  string
	Sender   string
	Password string
	Auth     smtp.Auth
}

var mailSetup *SMTPMailSetup
var once sync.Once

func getSetup() *SMTPMailSetup {
	once.Do(func() {
		mailSetup = &SMTPMailSetup{
			Server:   "smtp.gmail.com",
			Port:     "587",
			Sender:   "your_email",             // change with your own email
			Password: "generated_app_password", // change with your trimmed generated app password earlier
		}

		mailSetup.Address = fmt.Sprintf("%s:%s", mailSetup.Server, mailSetup.Port)
		mailSetup.Auth = smtp.PlainAuth("", mailSetup.Sender, mailSetup.Password, mailSetup.Server)
	})

	return mailSetup
}
