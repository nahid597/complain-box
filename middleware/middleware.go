package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nahidh597/complain-box/utils"
)

func IsAuthenticated(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	if _, err := utils.ParseJwt(cookie); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unacthorized",
		})
	}

	return c.Next()
}
