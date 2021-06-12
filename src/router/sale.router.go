package router

import (
	"github.com/Leonardo-Antonio/api.driving-school/src/handler"
	"github.com/Leonardo-Antonio/api.driving-school/src/middleware"
	"github.com/Leonardo-Antonio/api.driving-school/src/model"
	"github.com/Leonardo-Antonio/api.driving-school/src/utils"
	"github.com/gofiber/fiber/v2"
)

func Sale(storage model.ISale, app *fiber.App) {
	handler := handler.NewSale(storage)
	group := app.Group(
		utils.Config().BaseUri+"/sales",
		middleware.Authorization,
		new(middleware.Unauthorized).Client,
	)
	group.Post("/", handler.Buy)
}
