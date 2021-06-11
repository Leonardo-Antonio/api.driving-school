package router

import (
	"github.com/Leonardo-Antonio/api.driving-school/src/model"
	"github.com/Leonardo-Antonio/api.driving-school/src/utils"
	"github.com/gofiber/fiber/v2"
)

func Sale(storage model.ISale, app *fiber.App) {
	group := app.Group(utils.Config().BaseUri + "/sales")
	group.Post("/")
}
