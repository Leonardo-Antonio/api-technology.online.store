package server

import (
	"github.com/Leonardo-Antonio/api-technology.online.store/src/dbutil"
	"github.com/Leonardo-Antonio/api-technology.online.store/src/model"
	"github.com/Leonardo-Antonio/api-technology.online.store/src/router"
	"github.com/Leonardo-Antonio/api-technology.online.store/src/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
)

type app struct {
	fiberApp *fiber.App
}

func NewApp() *app {
	return &app{
		fiberApp: fiber.New(),
	}
}

func (a *app) Middlewares()  {
	a.fiberApp.Use(logger.New())
}

func (a *app) Routes() {
	db := dbutil.Connection()
	router.User(model.NewUser(db), a.fiberApp)
}

func (a *app) Listening() {
	log.Fatalln(a.fiberApp.Listen(utils.Config().Port))
}

