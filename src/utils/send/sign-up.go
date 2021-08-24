package send

import (
	"log"

	"github.com/Leonardo-Antonio/api.lyabook/src/entity"
)

func EmailSignUp(user entity.User) error {
	tpl, err := readTemplate("sign-up.html", "template/sign-up.html", user)
	if err != nil {
		return err
	}
	log.Println(tpl)
	if err := send(user.Email, tpl); err != nil {
		return err
	}
	return nil
}
