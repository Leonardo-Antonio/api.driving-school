package router

import (
	"github.com/Leonardo-Antonio/api.driving-school/src/handler"
	"github.com/Leonardo-Antonio/api.driving-school/src/middleware"
	"github.com/Leonardo-Antonio/api.driving-school/src/model"
	"github.com/Leonardo-Antonio/api.driving-school/src/utils"
	"github.com/gofiber/fiber/v2"
)

func Package(storage model.IPackage, app *fiber.App) {
	handler := handler.NewPackage(storage)
	group := app.Group(
		utils.Config().BaseUri + "/packages",
	)

	group.Post(
		"/",
		middleware.Authorization,
		new(middleware.Unauthorized).Admin,
		handler.Create,
	)
	group.Get("/", handler.GetAll)
	group.Get("/:id", handler.GetById)
	group.Get("/name/:name", handler.GetByName)
	group.Put(
		"/",
		middleware.Authorization,
		new(middleware.Unauthorized).Admin,
		handler.Edit,
	)
	group.Delete(
		"/:id",
		middleware.Authorization,
		new(middleware.Unauthorized).Admin,
		handler.Delete,
	)
}
