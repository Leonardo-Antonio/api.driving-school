package middleware

import (
	"errors"
	"net/http"

	"github.com/Leonardo-Antonio/api.driving-school/src/autorization"
	"github.com/Leonardo-Antonio/api.driving-school/src/utils"
	"github.com/Leonardo-Antonio/api.driving-school/src/utils/const/roles"
	"github.com/gofiber/fiber/v2"
)

var (
	errUnauthorized = errors.New("you do not have the necessary permissions to perform this operation")
)

type Unauthorized struct{}

func (u *Unauthorized) Instructor(ctx *fiber.Ctx) error {
	token := ctx.Get("Authorization")
	claims, err := autorization.ValidateToken(token)
	if err != nil {
		return ctx.Status(http.StatusForbidden).
			JSON(utils.ResponseErr(err.Error(), nil))
	}

	if claims.Rol != roles.INSTRUCTOR {
		return ctx.Status(http.StatusUnauthorized).
			JSON(utils.ResponseErr(errUnauthorized.Error(), nil))
	}

	return ctx.Next()
}

func (u *Unauthorized) Admin(ctx *fiber.Ctx) error {
	token := ctx.Get("Authorization")
	claims, err := autorization.ValidateToken(token)
	if err != nil {
		return ctx.Status(http.StatusForbidden).
			JSON(utils.ResponseErr(err.Error(), nil))
	}

	if claims.Rol != roles.ADMIN {
		return ctx.Status(http.StatusUnauthorized).
			JSON(utils.ResponseErr(errUnauthorized.Error(), nil))
	}

	return ctx.Next()
}

func (u *Unauthorized) Client(ctx *fiber.Ctx) error {
	token := ctx.Get("Authorization")
	claims, err := autorization.ValidateToken(token)
	if err != nil {
		return ctx.Status(http.StatusForbidden).
			JSON(utils.ResponseErr(err.Error(), nil))
	}

	if claims.Rol != roles.CLIENT {
		return ctx.Status(http.StatusUnauthorized).
			JSON(utils.ResponseErr(errUnauthorized.Error(), nil))
	}

	return ctx.Next()
}
