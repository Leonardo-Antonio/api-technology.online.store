package handler

import (
	"github.com/Leonardo-Antonio/api-technology.online.store/src/entity"
	"github.com/Leonardo-Antonio/api-technology.online.store/src/model"
	"github.com/Leonardo-Antonio/api-technology.online.store/src/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type category struct {
	storage model.ICategory
}

func NewCategory(storage model.ICategory) *category {
	return &category{storage}
}

func (ct *category) Create(c *fiber.Ctx) error {
	dataCategory := new(entity.Category)
	if err := c.BodyParser(dataCategory); err != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.ResponseError(err))
	}

	if err := utils.CategoryValidation(*dataCategory); err != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.ResponseError(err))
	}
	result, err := ct.storage.Create(*dataCategory)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.ResponseError(err))
	}

	if result.InsertedID == nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.ResponseError(err))
	}

	response := utils.ResponseSatisfactory("category created successfully", result)
	return c.Status(http.StatusCreated).JSON(response)
}

func (ct *category) GetAll(c *fiber.Ctx) error {
	categories, err := ct.storage.FindAll()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.ResponseError(err))
	}
	response := utils.ResponseSatisfactory("OK", categories)
	return c.Status(http.StatusOK).JSON(response)
}

func (ct *category) GetByID(c *fiber.Ctx) error {
	ID, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.ResponseError(err))
	}
	categoryData, err := ct.storage.FindOne(ID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.ResponseError(err))
	}
	response := utils.ResponseSatisfactory("Ok", categoryData)
	return c.Status(http.StatusOK).JSON(response)
}

func (ct *category) Update(c *fiber.Ctx) error {
	dataCategory := new(entity.Category)
	if err := c.BodyParser(&dataCategory); err != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.ResponseError(err))
	}
	if dataCategory.ID.IsZero() {
		return c.Status(http.StatusBadRequest).JSON(utils.ResponseError(utils.ErrIDRequired))
	}
	if err := utils.CategoryValidation(*dataCategory); err != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.ResponseError(err))
	}
	result, err := ct.storage.Update(*dataCategory)
	if result.ModifiedCount != 1 {
		return c.Status(http.StatusBadRequest).JSON(utils.ResponseError(utils.ErrCategoryNotExists))
	}
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.ResponseError(err))
	}
	response := utils.ResponseSatisfactory("category updated successfully", result)
	return c.Status(http.StatusOK).JSON(response)
}

func (ct *category) RemoveID(c *fiber.Ctx) error {
	ID, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.ResponseError(err))
	}

	result, err := ct.storage.Delete(ID)
	if result.ModifiedCount != 1 {
		return c.Status(http.StatusBadRequest).JSON(utils.ResponseError(utils.ErrCategoryNotExists))
	}
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.ResponseError(err))
	}

	response := utils.ResponseSatisfactory("category deleted successfully", result)
	return c.Status(http.StatusOK).JSON(response)
}
