package utils

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Protected(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Missing authorization token",
		})
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	// Validate the JWT token
	claims, err := ValidateJWT(tokenString)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid or expired token",
		})
	}

	c.Locals("username", claims["username"])

	return c.Next()
}
