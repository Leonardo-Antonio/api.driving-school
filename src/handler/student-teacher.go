package handler

import (
	"net/http"
	"strings"

	"github.com/Leonardo-Antonio/api.driving-school/src/entity"
	"github.com/Leonardo-Antonio/api.driving-school/src/model"
	"github.com/Leonardo-Antonio/api.driving-school/src/utils"
	"github.com/Leonardo-Antonio/api.driving-school/src/utils/validate"
	"github.com/gofiber/fiber/v2"
)

type studentTeacher struct {
	storage model.IStudentTeacher
}

func NewStudentTeacher(storage model.IStudentTeacher) *studentTeacher {
	return &studentTeacher{storage}
}

func (st *studentTeacher) GetByTurn(ctx *fiber.Ctx) error {
	turn := strings.Title(strings.ToLower(ctx.Params("turn")))

	studentTeacher, err := st.storage.FindByTurn(turn)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(utils.ResponseErr(err.Error(), nil))
	}

	if &studentTeacher == nil {
		return ctx.Status(http.StatusNoContent).
			JSON(utils.ResponseErr(err.Error(), nil))
	}

	return ctx.Status(http.StatusOK).
		JSON(utils.ResponseSatisfactory("ok", studentTeacher))
}

func (st *studentTeacher) AssignStudentToTeacher(ctx *fiber.Ctx) error {
	assignStudentToTeacher := new(entity.AssignStudentTeacher)
	if err := ctx.BodyParser(&assignStudentToTeacher); err != nil {
		return ctx.Status(http.StatusBadRequest).
			JSON(utils.ResponseErr(err.Error(), nil))
	}

	if err := validate.StudentTeacher(*assignStudentToTeacher); err != nil {
		return ctx.Status(http.StatusBadRequest).
			JSON(utils.ResponseErr(err.Error(), nil))
	}

	result, err := st.storage.AssingStudentToTeacher(*assignStudentToTeacher)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(utils.ResponseErr(err.Error(), nil))
	}

	return ctx.Status(http.StatusCreated).
		JSON(
			utils.ResponseSatisfactory(
				"the student was correctly assigned before the teacher",
				result,
			),
		)
}
