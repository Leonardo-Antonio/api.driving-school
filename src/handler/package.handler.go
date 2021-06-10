package handler

import (
	"net/http"

	"github.com/Leonardo-Antonio/api.driving-school/src/entity"
	"github.com/Leonardo-Antonio/api.driving-school/src/model"
	"github.com/Leonardo-Antonio/api.driving-school/src/utils"
	"github.com/Leonardo-Antonio/api.driving-school/src/utils/validate"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type pack struct {
	storage model.IPackage
}

func NewPackage(storage model.IPackage) *pack {
	return &pack{storage}
}

func (p *pack) Create(ctx *fiber.Ctx) error {
	var pack entity.Package
	if err := ctx.BodyParser(&pack); err != nil {
		return ctx.Status(http.StatusBadRequest).
			JSON(utils.ResponseErr(err.Error(), nil))
	}

	if err := validate.Package(pack); err != nil {
		return ctx.Status(http.StatusBadRequest).
			JSON(utils.ResponseErr(err.Error(), nil))
	}
	validate.FieldsPackage(&pack)

	result, err := p.storage.Insert(pack)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(utils.ResponseErr(err.Error(), nil))
	}

	return ctx.Status(http.StatusCreated).
		JSON(
			utils.ResponseSatisfactory(
				"the package was created successfully",
				result,
			),
		)
}

func (p *pack) GetAll(ctx *fiber.Ctx) error {
	packages, err := p.storage.FindAll()
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(utils.ResponseErr(err.Error(), nil))
	}

	if len(packages) == 0 || packages == nil {
		return ctx.Status(http.StatusNotFound).
			JSON(utils.ResponseErr(err.Error(), nil))
	}

	return ctx.Status(http.StatusOK).
		JSON(
			utils.ResponseSatisfactory(
				"ok",
				packages,
			),
		)
}

func (p *pack) GetById(ctx *fiber.Ctx) error {
	id, err := primitive.ObjectIDFromHex(ctx.Params("id"))
	if err != nil {
		return ctx.Status(http.StatusBadRequest).
			JSON(utils.ResponseErr(utils.ErrIdInvalid.Error(), nil))
	}
	if id.IsZero() {
		return ctx.Status(http.StatusBadRequest).
			JSON(utils.ResponseErr(utils.ErrIdInvalid.Error(), nil))
	}

	pack, err := p.storage.FindById(id)
	if pack.ID.IsZero() {
		return ctx.Status(http.StatusBadRequest).
			JSON(utils.ResponseErr(utils.ErrIdInvalid.Error(), nil))
	}

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(utils.ResponseErr(err.Error(), nil))
	}

	return ctx.Status(http.StatusOK).
		JSON(
			utils.ResponseSatisfactory(
				"ok",
				pack,
			),
		)
}

func (p *pack) GetByName(ctx *fiber.Ctx) error {
	name := ctx.Params("name")
	pack, err := p.storage.FindByName(name)

	if pack.ID.IsZero() {
		return ctx.Status(http.StatusBadRequest).
			JSON(utils.ResponseErr(utils.ErrNamePackageNotExist.Error(), nil))
	}
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(utils.ResponseErr(err.Error(), nil))
	}

	return ctx.Status(http.StatusOK).
		JSON(
			utils.ResponseSatisfactory(
				"ok",
				pack,
			),
		)
}

func (p *pack) Edit(ctx *fiber.Ctx) error {
	var pack entity.Package
	if err := ctx.BodyParser(&pack); err != nil {
		return ctx.Status(http.StatusBadRequest).
			JSON(utils.ResponseErr(err.Error(), nil))
	}

	if pack.ID.IsZero() {
		return ctx.Status(http.StatusBadRequest).
			JSON(utils.ResponseErr(utils.ErrIdInvalid.Error(), nil))
	}

	if err := validate.Package(pack); err != nil {
		return ctx.Status(http.StatusBadRequest).
			JSON(utils.ResponseErr(err.Error(), nil))
	}
	validate.FieldsPackage(&pack)

	result, err := p.storage.Update(pack)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(utils.ResponseErr(err.Error(), nil))
	}

	if result.ModifiedCount == 0 {
		return ctx.Status(http.StatusNotModified).
			JSON(utils.ResponseErr("the package and it is up to date", nil))
	}

	return ctx.Status(http.StatusCreated).
		JSON(
			utils.ResponseSatisfactory(
				"the package was updated successfully",
				result,
			),
		)
}

func (p *pack) Delete(ctx *fiber.Ctx) error {
	id, err := primitive.ObjectIDFromHex(ctx.Params("id"))
	if err != nil {
		return ctx.Status(http.StatusBadRequest).
			JSON(utils.ResponseErr(utils.ErrIdInvalid.Error(), nil))
	}
	if id.IsZero() {
		return ctx.Status(http.StatusBadRequest).
			JSON(utils.ResponseErr(utils.ErrIdInvalid.Error(), nil))
	}

	result, err := p.storage.RemoveById(id)
	if result.ModifiedCount == 0 {
		return ctx.Status(http.StatusNotModified).
			JSON(utils.ResponseErr("the package has already been deactivated", nil))
	}

	return ctx.Status(http.StatusOK).
		JSON(
			utils.ResponseSatisfactory(
				"the package was deleted successfull",
				result,
			),
		)
}
