package validate

import (
	"errors"
	"strings"

	"github.com/Leonardo-Antonio/api.driving-school/src/entity"
	"github.com/Leonardo-Antonio/api.driving-school/src/utils/const/roles"
	"github.com/Leonardo-Antonio/api.driving-school/src/utils/const/turn"
)

var (
	errEmailInvalid     = errors.New("the email you entered is invalid")
	errDniInvalid       = errors.New("the dni you entered is invalid")
	errRolInvalid       = errors.New("the role you are trying to enter is not valid")
	errTurnInvalid      = errors.New("the turn you are trying to enter is not valid")
	errInsecurePassword = errors.New("the password entered is not very secure")
	errRolRequired      = errors.New("the role field is required")
	errInstructorTurn   = errors.New("the instructor must have a work schedule")
)

func User(user entity.User) error {
	if len(user.Password) <= 8 {
		return errInsecurePassword
	}

	if len(user.Rol) == 0 {
		return errRolRequired
	}

	if strings.EqualFold(roles.INSTRUCTOR, user.Rol) {
		if len(user.Turn) == 0 || &user.Turn == nil {
			return errInstructorTurn
		}
		if strings.EqualFold(turn.MORNING, user.Turn) ||
			strings.EqualFold(turn.AFTERNOON, user.Turn) ||
			strings.EqualFold(turn.NIGHT, user.Turn) {
			return nil
		}
		return errInstructorTurn
	}

	if strings.EqualFold(roles.ADMIN, user.Rol) ||
		strings.EqualFold(roles.MANAGER, user.Rol) ||
		strings.EqualFold(roles.INSTRUCTOR, user.Rol) ||
		strings.EqualFold(roles.CLIENT, user.Rol) {
		return nil
	}
	return errRolInvalid
}

func UserDni(dni string) error {
	if len(dni) != 8 {
		return errDniInvalid
	}
	return nil
}

func UserEmail(email string) error {
	if !strings.Contains(email, "@") {
		return errEmailInvalid
	}
	return nil
}
