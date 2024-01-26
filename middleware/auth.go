package middleware

import (
	"seangjr/kehilah/handlers"
	"seangjr/kehilah/model"

	"github.com/gofiber/fiber/v2"
)

func Authenticated(c *fiber.Ctx) error {
	json := new(model.Session)
	if err := c.BodyParser(json); err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid Session Format",
		})
	}
	user, err := handlers.GetUser(json.Sessionid)
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    404,
			"message": "404: not found",
		})
	}
	c.Locals("user", user)
	return c.Next()
}

func IsAdmin(c *fiber.Ctx) error {
	user, err := handlers.GetUser(c.Locals("user").(*model.User).ID)
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    404,
			"message": "404: not found",
		})
	}
	if user.Role != "admin" {
		return c.JSON(fiber.Map{
			"code":    401,
			"message": "401: unauthorized",
		})
	}
	return c.Next()
}