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
	router.User(model.NewUser(dbutil.GetConnection()), server.app)
}

func (server *server) Listeing() {
	log.Fatalln(server.app.Listen(utils.Config().Port))
}
