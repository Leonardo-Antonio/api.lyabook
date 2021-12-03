package reniec

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Leonardo-Antonio/api.lyabook/src/entity"
	"github.com/Leonardo-Antonio/api.lyabook/src/utils/env"
	"github.com/Leonardo-Antonio/api.lyabook/src/utils/errores"
)

func GetDniReniec(user *entity.User) error {
	resp, err := http.Get(env.Data.ApiReniecDni + user.Dni)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errores.ErrFindDniApiReniec
	}

	bodyJson, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var dniData *entity.Dni
	if err := json.Unmarshal(bodyJson, &dniData); err != nil {
		return err
	}

	user.Name = dniData.Nombres
	user.LastName = dniData.ApellidoPaterno + " " + dniData.ApellidoMaterno

	return nil
}
