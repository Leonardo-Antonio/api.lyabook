package formatting

import (
	"strings"

	"github.com/Leonardo-Antonio/api.lyabook/src/entity"
)

func User(user *entity.User) {
	user.Name = strings.Title(strings.ToLower(user.Name))
	user.LastName = strings.Title(strings.ToLower(user.LastName))
	user.Rol = strings.Title(strings.ToLower(user.Rol))
}
