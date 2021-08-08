package send

import (
	"net/smtp"

	"github.com/Leonardo-Antonio/api.lyabook/src/utils/env"
	"github.com/jordan-wright/email"
)

func send(to, tpl string) error {
	e := email.NewEmail()
	e.From = "Leonardo Antonio Nolasco Leyva <" + env.Data.Email + ">"
	e.To = []string{to}
	e.Subject = "Bienvenid@ a LyaBook"
	e.HTML = []byte(tpl)
	if err := e.Send("smtp.gmail.com:587", smtp.PlainAuth(
		"", env.Data.Email,
		env.Data.SecretKeyApplicationEmail,
		"smtp.gmail.com",
	)); err != nil {
		return err
	}

	return nil
}
