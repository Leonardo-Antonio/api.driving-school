package main

import (
	"log"

	"github.com/Leonardo-Antonio/api.driving-school/src/app"
	"github.com/Leonardo-Antonio/api.driving-school/src/autorization"
	"github.com/joho/godotenv"
)

func main() {
	if err := autorization.LoadFiles(); err != nil {
		log.Fatalln(err)
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}

	app := app.NewAppServer()
	app.Middlewares()
	app.Routers()
	app.Listeing()
}
