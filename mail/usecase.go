package mailer

import (
	"bytes"
	"fmt"
	"net/smtp"
	"text/template"
)

type Mailer struct {
	Body  []byte
	Setup *SMTPMailSetup
}

func NewSMTPMailer() *Mailer {
	return &Mailer{
		Setup: getSetup(),
	}
}

func (m *Mailer) createBody(subject, htmlTemplate string, data any) (err error) {
	var body bytes.Buffer

	const mime = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: %s\n%s\n\n", subject, mime)))
	t, err := template.ParseFiles(htmlTemplate)
	if err != nil {
		return
	}

	err = t.Execute(&body, data)
	if err != nil {
		return
	}

	m.Body = body.Bytes()

	return nil
}

func (m *Mailer) Send(receipent, subject, template string, data any) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s | %v", "error on sending email", r)
		}
	}()

	err = m.createBody(subject, template, data)
	if err != nil {
		return
	}

	err = smtp.SendMail(m.Setup.Address, m.Setup.Auth, m.Setup.Sender, []string{receipent}, m.Body)

	return
}
