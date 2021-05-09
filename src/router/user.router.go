package router

import (
	"github.com/Leonardo-Antonio/api-technology.online.store/src/autorization"
	"github.com/Leonardo-Antonio/api-technology.online.store/src/handler"
	"github.com/Leonardo-Antonio/api-technology.online.store/src/model"
	"github.com/Leonardo-Antonio/api-technology.online.store/src/utils"
	"github.com/gofiber/fiber/v2"
)

func User(storage model.IUser, app *fiber.App) {
	controller := handler.NewUser(storage)
	userGroup := app.Group(utils.Config().UrlMain + "users")
	userGroup.Post("/sign-in", controller.SignIn)
	userGroup.Delete("/:id", autorization.Middleware(), controller.RemoveByID)
	userGroup.Post("/log-in", controller.LogIn)
	userGroup.Put("/", controller.Update)
	userGroup.Get("", autorization.Middleware(), controller.GetAll)
}
