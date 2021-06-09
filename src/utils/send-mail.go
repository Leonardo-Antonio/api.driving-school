package utils

import (
	"fmt"
	"io/ioutil"
	"net/smtp"

	"github.com/Leonardo-Antonio/api.driving-school/src/entity"
	"github.com/jordan-wright/email"
)

func SendEmailToAccountCreated(user *entity.User) error {
	body, err := ioutil.ReadFile("./public/template/account-created.html")
	if err != nil {
		return err
	}
	mail := email.NewEmail()
	mail.From = "Leonardo Antonio Nolasco Leyva <leo2001.nl08@gmail.com>"
	mail.To = []string{user.Email}
	mail.Subject = "Se creo su cuenta exitosamente"
	mail.HTML = []byte(fmt.Sprintf(string(body), user.Names, user.Email, user.Password))
	if err := mail.Send(
		"smtp.gmail.com:587",
		smtp.PlainAuth(
			"", "leo2001.nl08@gmail.com",
			Config().PasswordEmail,
			"smtp.gmail.com",
		),
	); err != nil {
		return err
	}
	return nil
}
