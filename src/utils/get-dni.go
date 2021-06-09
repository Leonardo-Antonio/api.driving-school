package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/Leonardo-Antonio/api.driving-school/src/entity"
)

func GetDniReniec(user *entity.User) error {
	if len(user.DNI) != 8 {
		return ErrDniInvalid
	}
	response, err := http.Get(
		Config().ApiReniecDni +
			"/" + user.DNI + "?token=" +
			Config().TokenReniec,
	)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return ErrDniNotFound
	}
	bodyJson, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	var dniData entity.Dni
	if err := json.Unmarshal(bodyJson, &dniData); err != nil {
		return err
	}

	user.Names = strings.Title(strings.ToLower(dniData.Names))
	user.LastNames = fmt.Sprintf("%s %s",
		strings.Title(strings.ToLower(dniData.FathersLastName)),
		strings.Title(strings.ToLower(dniData.MothersLastName)),
	)

	return nil
}
