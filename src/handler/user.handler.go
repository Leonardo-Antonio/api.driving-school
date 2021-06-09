package handler

import (
	"encoding/base64"
	"errors"
	"net/http"
	"strings"

	"github.com/Leonardo-Antonio/api.driving-school/src/autorization"
	"github.com/Leonardo-Antonio/api.driving-school/src/entity"
	"github.com/Leonardo-Antonio/api.driving-school/src/model"
	"github.com/Leonardo-Antonio/api.driving-school/src/utils"
	"github.com/Leonardo-Antonio/api.driving-school/src/utils/validate"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type user struct {
	storage model.IUser
}

func NewUser(storage model.IUser) *user {
	return &user{storage}
}

func (u *user) GetAll(ctx *fiber.Ctx) error {
	users, err := u.storage.FindAll()
	if err != nil {
		response := utils.Response(utils.ERR, err.Error(), true, nil)
		return ctx.Status(http.StatusInternalServerError).JSON(response)
	}
	if len(users) == 0 || users == nil {
		response := utils.Response(utils.MSG, "no user data is available", false, users)
		return ctx.Status(http.StatusNoContent).JSON(response)
	}

	response := utils.ResponseSatisfactory("satifactory", users)
	return ctx.Status(http.StatusOK).JSON(response)
}

func (u *user) LogInDNI(ctx *fiber.Ctx) error {
	var user entity.User
	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(http.StatusBadRequest).
			JSON(utils.ResponseErr(err.Error(), nil))
	}

	if err := validate.UserDni(user.DNI); err != nil {
		return ctx.Status(http.StatusBadRequest).
			JSON(utils.ResponseErr(err.Error(), nil))
	}

	user.Password = base64.StdEncoding.EncodeToString([]byte(user.Password))
	data, err := u.storage.FindByDNI(user)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(utils.ResponseErr(err.Error(), nil))
	}

	if data.ID.IsZero() {
		return ctx.Status(http.StatusBadRequest).
			JSON(utils.ResponseErr("the data that I enter is not correct", nil))
	}

	token, err := autorization.GenerateToken(&data)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(utils.ResponseErr(err.Error(), nil))
	}

	response := make(map[string]interface{}, 2)
	response["token"] = token
	response["user"] = data

	return ctx.Status(http.StatusOK).
		JSON(utils.ResponseSatisfactory("ok", response))
}

func (u *user) LogInEmail(ctx *fiber.Ctx) error {
	var user entity.User
	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(http.StatusBadRequest).
			JSON(utils.ResponseErr(err.Error(), nil))
	}

	if err := validate.UserEmail(user.Email); err != nil {
		return ctx.Status(http.StatusBadRequest).
			JSON(utils.ResponseErr(err.Error(), nil))
	}

	user.Password = base64.StdEncoding.EncodeToString([]byte(user.Password))
	data, err := u.storage.FindByEmail(user)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(utils.ResponseErr(err.Error(), nil))
	}

	if data.ID.IsZero() {
		return ctx.Status(http.StatusBadRequest).
			JSON(utils.ResponseErr("the data that I enter is not correct", nil))
	}

	token, err := autorization.GenerateToken(&data)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(utils.ResponseErr(err.Error(), nil))
	}

	response := make(map[string]interface{}, 2)
	response["token"] = token
	response["user"] = data

	return ctx.Status(http.StatusOK).
		JSON(utils.ResponseSatisfactory("ok", response))
}

func (u *user) SignUpDni(ctx *fiber.Ctx) error {
	var user entity.User
	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(http.StatusBadRequest).
			JSON(utils.ResponseErr(err.Error(), nil))
	}

	if err := utils.GetDniReniec(&user); err != nil {
		response := utils.ResponseErr(err.Error(), nil)
		if errors.Is(err, utils.ErrDniInvalid) ||
			errors.Is(err, utils.ErrDniNotFound) {
			return ctx.Status(http.StatusBadRequest).
				JSON(response)
		}
		return ctx.Status(http.StatusInternalServerError).JSON(response)
	}

	user.Rol = strings.Title(strings.ToLower(user.Rol))
	if err := validate.User(user); err != nil {
		return ctx.Status(http.StatusBadRequest).
			JSON(utils.ResponseErr(err.Error(), nil))
	}

	user.Password = base64.StdEncoding.EncodeToString([]byte(user.Password))
	result, err := u.storage.Insert(user)
	if err != nil {
		response := utils.ResponseErr(err.Error(), nil)
		if mongo.IsDuplicateKeyError(err) {
			return ctx.Status(http.StatusBadRequest).JSON(response)
		}
		return ctx.Status(http.StatusInternalServerError).JSON(response)
	}

	response := utils.ResponseSatisfactory("the user was created successfully", result)
	return ctx.Status(http.StatusCreated).JSON(response)
}

func (u *user) SignUpEmail(ctx *fiber.Ctx) error {
	var user entity.User
	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(http.StatusBadRequest).
			JSON(utils.ResponseErr(err.Error(), nil))
	}

	if err := validate.UserEmail(user.Email); err != nil {
		return ctx.Status(http.StatusBadRequest).
			JSON(utils.ResponseErr(err.Error(), nil))
	}
	if err := validate.User(user); err != nil {
		return ctx.Status(http.StatusBadRequest).
			JSON(utils.ResponseErr(err.Error(), nil))
	}

	validate.FieldsUser(&user)
	user.Password = base64.StdEncoding.EncodeToString([]byte(user.Password))
	result, err := u.storage.Insert(user)
	if err != nil {
		response := utils.ResponseErr(err.Error(), nil)
		if mongo.IsDuplicateKeyError(err) {
			return ctx.Status(http.StatusBadRequest).JSON(response)
		}
		return ctx.Status(http.StatusInternalServerError).JSON(response)
	}

	if err := utils.SendEmailToAccountCreated(&user); err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(utils.ResponseErr(err.Error(), nil))
	}

	response := utils.ResponseSatisfactory("the user was created successfully", result)
	return ctx.Status(http.StatusCreated).JSON(response)
}

func (u *user) Delete(ctx *fiber.Ctx) error {
	ID, err := primitive.ObjectIDFromHex(ctx.Params("id"))
	if err != nil {
		return ctx.Status(http.StatusBadRequest).
			JSON(utils.ResponseErr(err.Error(), nil))
	}

	result, err := u.storage.Remove(ID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(utils.ResponseErr(err.Error(), nil))
	}

	if result.ModifiedCount == 0 {
		return ctx.Status(http.StatusNotModified).
			JSON(utils.ResponseErr("the account has already been deactivated", nil))
	}

	response := utils.ResponseSatisfactory("the user was deleted successfully", result)
	return ctx.Status(http.StatusOK).JSON(response)
}

func (u *user) Edit(ctx *fiber.Ctx) error {
	user := new(entity.User)
	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(http.StatusBadRequest).
			JSON(utils.ResponseErr(err.Error(), nil))
	}

	if err := validate.User(*user); err != nil {
		return ctx.Status(http.StatusBadRequest).
			JSON(utils.ResponseErr(err.Error(), nil))
	}

	user.Password = base64.StdEncoding.EncodeToString([]byte(user.Password))
	result, err := u.storage.Update(*user)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(utils.ResponseErr(err.Error(), nil))
	}

	if result.ModifiedCount == 0 {
		return ctx.Status(http.StatusNotModified).
			JSON(utils.ResponseErr("no field has been modified", nil))
	}

	response := utils.ResponseSatisfactory("the user was updated successfully", result)
	return ctx.Status(http.StatusOK).JSON(response)
}
