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

func (u *user) SignIn(c *fiber.Ctx) error {
	dataUser := new(entity.User)
	if err := c.BodyParser(dataUser); err != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.ResponseError(err))
	}

	if err := utils.UserValidate(*dataUser); err != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.ResponseError(err))
	}

	dataUser.Password = base64.StdEncoding.EncodeToString([]byte(dataUser.Password))

	result, err := u.storage.Create(*dataUser)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.ResponseError(err))
	}
	response := utils.ResponseSatisfactory("user created successfully", result)
	return c.Status(http.StatusCreated).JSON(response)
}

func (u *user) RemoveByID(c *fiber.Ctx) error {
	objectID, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.ResponseError(err))
	}

	result, err := u.storage.RemoveByID(objectID)
	if result.MatchedCount != 1 {
		return c.Status(http.StatusBadRequest).JSON(utils.ResponseError(utils.ErrUserNotExists))
	}
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(utils.ResponseError(err))
	}
	response := utils.ResponseSatisfactory("user deleted successfully", result)
	return c.Status(http.StatusOK).JSON(response)
}

func (u *user) LogIn(c *fiber.Ctx) error {
	userData := new(entity.User)
	if err := c.BodyParser(userData); err != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.ResponseError(err))
	}

	userData.Password = base64.StdEncoding.EncodeToString([]byte(userData.Password))

	dataUser, err := u.storage.FindOne(*userData)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(utils.ResponseError(err))
	}

	token, err := autorization.GenerateToken(dataUser)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.ResponseError(err))
	}

	userDataAndToken := map[string]interface{}{
		"token": token,
		"user":  dataUser,
	}
	response := utils.ResponseSatisfactory("user found successfully", userDataAndToken)
	return c.Status(http.StatusOK).JSON(response)
}

func (u *user) Update(c *fiber.Ctx) error {
	dataUser := new(entity.User)
	if err := c.BodyParser(&dataUser); err != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.ResponseError(err))
	}

	if dataUser.ID.IsZero() {
		return c.Status(http.StatusBadRequest).JSON(utils.ResponseError(utils.ErrIDRequired))
	}
	if err := utils.UserValidate(*dataUser); err != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.ResponseError(err))
	}
	dataUser.Password = base64.StdEncoding.EncodeToString([]byte(dataUser.Password))
	result, err := u.storage.UpdateByID(*dataUser)
	if result.MatchedCount != 1 {
		return c.Status(http.StatusBadRequest).JSON(utils.ResponseError(utils.ErrUserNotExists))
	}
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.ResponseError(err))
	}

	response := utils.ResponseSatisfactory("user updated successfully", result)
	return c.Status(http.StatusOK).JSON(response)
}

func (u *user) GetAll(c *fiber.Ctx) error {
	users, err := u.storage.FindAll()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.ResponseError(err))
	}
	response := utils.ResponseSatisfactory("Ok", users)
	return c.Status(http.StatusOK).JSON(response)
}
