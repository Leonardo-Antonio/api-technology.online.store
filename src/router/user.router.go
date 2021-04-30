package router

import (
	"github.com/Leonardo-Antonio/api-technology.online.store/src/handler"
	"github.com/Leonardo-Antonio/api-technology.online.store/src/model"
	"github.com/Leonardo-Antonio/api-technology.online.store/src/utils"
	"github.com/gofiber/fiber/v2"
)

func User(storage model.IUser, app *fiber.App) {
	controller := handler.NewUser(storage)
	userGroup := app.Group(utils.Config().UrlMain + "users")
	userGroup.Post("", controller.Create)
	userGroup.Delete("/:id", controller.RemoveByID)
}
