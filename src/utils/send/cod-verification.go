package send

import "github.com/Leonardo-Antonio/api.lyabook/src/entity"

func CodVerification(user entity.User) error {
	tpl, err := readTemplate("cod-verification.html", "./template/cod-verification.html", user)
	if err != nil {
		return err
	}
	if err := send(user.Email, tpl); err != nil {
		return err
	}
	return nil
}
