package app

import (
	"log"

	"github.com/Leonardo-Antonio/api.driving-school/src/dbutil"
	"github.com/Leonardo-Antonio/api.driving-school/src/model"
	"github.com/Leonardo-Antonio/api.driving-school/src/router"
	"github.com/Leonardo-Antonio/api.driving-school/src/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type server struct {
	app *fiber.App
}

func NewAppServer() *server {
	return &server{
		app: fiber.New(),
	}
}

func (server *server) Middlewares() {
	server.app.Use(logger.New())
}

func (server *server) Routers() {
	db := dbutil.GetConnection()
	router.User(model.NewUser(db), server.app)
	router.Package(model.NewPackage(db), server.app)
	router.StudentTeacher(model.NewStudentTeacher(db), server.app)
	router.Sale(model.NewSale(db), server.app)
}

func (server *server) Listeing() {
	log.Fatalln(server.app.Listen(utils.Config().Port))
}
