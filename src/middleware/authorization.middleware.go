package middleware

import (
	"net/http"

	"github.com/Leonardo-Antonio/api.driving-school/src/autorization"
	"github.com/Leonardo-Antonio/api.driving-school/src/utils"
	"github.com/gofiber/fiber/v2"
)

func Authorization(ctx *fiber.Ctx) error {
	token := ctx.Get("Authorization")
	_, err := autorization.ValidateToken(token)
	if err != nil {
		return ctx.Status(http.StatusForbidden).
			JSON(utils.ResponseErr(err.Error(), nil))
	}
	return ctx.Next()
}
