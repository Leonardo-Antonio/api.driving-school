package router

import (
	"github.com/Leonardo-Antonio/api.driving-school/src/handler"
	"github.com/Leonardo-Antonio/api.driving-school/src/middleware"
	"github.com/Leonardo-Antonio/api.driving-school/src/model"
	"github.com/Leonardo-Antonio/api.driving-school/src/utils"
	"github.com/gofiber/fiber/v2"
)

func User(storage model.IUser, app *fiber.App) {
	handler := handler.NewUser(storage)

	grorp := app.Group(utils.Config().BaseUri + "/users")
	grorp.Get("/", middleware.Authorization, handler.GetAll)
	grorp.Post("/sign-up/email", handler.SignUpEmail)
	grorp.Post("/sign-up/dni", handler.SignUpDni)
	grorp.Post("/log-in/dni", handler.LogInDNI)
	grorp.Post("/log-in/email", handler.LogInEmail)
	grorp.Delete("/:id", middleware.Authorization, handler.Delete)
	grorp.Put("/", middleware.Authorization, handler.Edit)
}
