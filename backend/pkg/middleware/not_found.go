package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func SetupNotFoundMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"errors":  "route/not-found",
		})
	}
}
