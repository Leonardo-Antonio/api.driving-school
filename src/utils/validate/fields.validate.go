package validate

import (
	"strings"

	"github.com/Leonardo-Antonio/api.driving-school/src/entity"
)

func FieldsUser(user *entity.User) {
	user.Names = strings.Title(strings.ToLower(user.Names))
	user.LastNames = strings.Title(strings.ToLower(user.LastNames))
	user.Rol = strings.Title(strings.ToLower(user.Rol))
}

func FieldsPackage(pack *entity.Package) {
	pack.Name = strings.Title(strings.ToLower(pack.Name))
	for i, value := range pack.Content {
		pack.Content[i] = strings.Title(strings.ToLower(value))
	}
}
