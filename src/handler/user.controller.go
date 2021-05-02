package handler

import (
	"encoding/base64"
	"github.com/Leonardo-Antonio/api-technology.online.store/src/autorization"
	"github.com/Leonardo-Antonio/api-technology.online.store/src/entity"
	"github.com/Leonardo-Antonio/api-technology.online.store/src/model"
	"github.com/Leonardo-Antonio/api-technology.online.store/src/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type user struct {
	storage model.IUser
}

func NewUser(storage model.IUser) *user {
	return &user{storage}
}

func (u *user) Create(c *fiber.Ctx) error {
	dataUser := new(entity.User)
	if err := c.BodyParser(dataUser); err != nil {
		response := utils.ResponseError(err)
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	if err := utils.UserValidate(*dataUser); err != nil {
		response := utils.ResponseError(err)
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	dataUser.Password = base64.StdEncoding.EncodeToString([]byte(dataUser.Password))

	result, err := u.storage.Create(*dataUser)
	if err != nil {
		response := utils.ResponseError(err)
		return c.Status(http.StatusBadRequest).JSON(response)
	}
	response := utils.ResponseSatisfactory("user created successfully", result)
	return c.Status(http.StatusCreated).JSON(response)
}

func (u *user) RemoveByID(c *fiber.Ctx) error {
	objectID, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		response := utils.ResponseError(err)
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	result, err := u.storage.RemoveByID(objectID)
	if err != nil {
		response := utils.ResponseError(err)
		return c.Status(http.StatusNotFound).JSON(response)
	}
	response := utils.ResponseSatisfactory("user deleted successfully", result)
	return c.Status(http.StatusOK).JSON(response)
}

func (u *user) Login(c *fiber.Ctx) error {
	userData := new(entity.User)
	if err := c.BodyParser(userData); err != nil {
		response := utils.ResponseError(err)
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	userData.Password = base64.StdEncoding.EncodeToString([]byte(userData.Password))

	dataUser, err := u.storage.FindOne(*userData)
	if err != nil {
		response := utils.ResponseError(err)
		return c.Status(http.StatusUnauthorized).JSON(response)
	}

	token, err := autorization.GenerateToken(dataUser)
	if err != nil {
		response := utils.ResponseError(err)
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	userDataAndToken := map[string]interface{}{
		"token": token,
		"user": userData,
	}
	response := utils.ResponseSatisfactory("user found successfully", userDataAndToken)
	return c.Status(http.StatusOK).JSON(response)
}

func (u *user) GetAll(c *fiber.Ctx) error {
	users, err := u.storage.FindAll()
	if err != nil {
		response := utils.ResponseError(err)
		return c.Status(http.StatusInternalServerError).JSON(response)
	}
	response := utils.ResponseSatisfactory("Ok", users)
	return c.Status(http.StatusOK).JSON(response)
}