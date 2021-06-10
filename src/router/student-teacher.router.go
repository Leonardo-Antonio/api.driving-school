package router

import (
	"github.com/Leonardo-Antonio/api.driving-school/src/handler"
	"github.com/Leonardo-Antonio/api.driving-school/src/middleware"
	"github.com/Leonardo-Antonio/api.driving-school/src/model"
	"github.com/Leonardo-Antonio/api.driving-school/src/utils"
	"github.com/gofiber/fiber/v2"
)

func StudentTeacher(storage model.IStudentTeacher, app *fiber.App) {
	handler := handler.NewStudentTeacher(storage)
	group := app.Group(
		utils.Config().BaseUri+"/student-teacher",
		middleware.Authorization,
	)

	group.Get("/:turn", handler.GetByTurn)
	group.Post("/", handler.AssignStudentToTeacher)
}
