package router

import (
	"github.com/Leonardo-Antonio/api-technology.online.store/src/autorization"
	"github.com/Leonardo-Antonio/api-technology.online.store/src/handler"
	"github.com/Leonardo-Antonio/api-technology.online.store/src/model"
	"github.com/Leonardo-Antonio/api-technology.online.store/src/utils"
	"github.com/gofiber/fiber/v2"
)

func Category(storage model.ICategory, app *fiber.App) {
	controller := handler.NewCategory(storage)
	categoryGroup := app.Group(utils.Config().UrlMain+"categories", autorization.Middleware())
	categoryGroup.Post("", controller.Create)
	categoryGroup.Get("", controller.GetAll)
	categoryGroup.Get("/:id", controller.GetByID)
	categoryGroup.Put("", controller.Update)
	categoryGroup.Delete("/:id", controller.RemoveID)
}
