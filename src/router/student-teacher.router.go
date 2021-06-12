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

	group.Get(
		"/:turn",
		new(middleware.Unauthorized).Admin,
		handler.GetByTurn,
	)
	group.Post(
		"/",
		new(middleware.Unauthorized).Admin,
		handler.AssignStudentToTeacher,
	)
	group.Get(
		"/teacher/:id",
		new(middleware.Unauthorized).Instructor,
		handler.StudentsByTeacher,
	)
}
