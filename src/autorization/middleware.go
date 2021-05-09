package autorization

import (
	"github.com/Leonardo-Antonio/api-technology.online.store/src/utils"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func Middleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("Authorization")
		_, err := ValidateToken(token)
		if err != nil {
			response := utils.ResponseError(err)
			return c.Status(http.StatusForbidden).JSON(response)
		}
		return c.Next()
	}
}
