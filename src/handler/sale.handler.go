package handler

import (
	"net/http"

	"github.com/Leonardo-Antonio/api.driving-school/src/entity"
	"github.com/Leonardo-Antonio/api.driving-school/src/model"
	"github.com/Leonardo-Antonio/api.driving-school/src/utils"
	"github.com/Leonardo-Antonio/api.driving-school/src/utils/validate"
	"github.com/gofiber/fiber/v2"
)

type sale struct {
	storage model.ISale
}

func NewSale(storage model.ISale) *sale {
	return &sale{storage}
}

func (s *sale) Buy(ctx *fiber.Ctx) error {
	var sale entity.Sale
	if err := ctx.BodyParser(&sale); err != nil {
		return ctx.Status(http.StatusBadRequest).
			JSON(utils.ResponseErr(err.Error(), nil))
	}

	if err := validate.IsZeroObjectIds(sale.IdClient, sale.IdPackage); err != nil {
		return ctx.Status(http.StatusBadRequest).
			JSON(utils.ResponseErr(err.Error(), nil))
	}

	return nil
}
