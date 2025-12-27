package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthRequired(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(401).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	tokenStr := authHeader[len("Bearer "):]

	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		return c.Status(401).JSON(fiber.Map{
			"error": "Invalid token",
		})
	}

	claims := token.Claims.(jwt.MapClaims)

	c.Locals("user_id", uint(claims["user_id"].(float64)))
	c.Locals("role", claims["role"].(string))

	return c.Next()
}

func AdminOnly(c *fiber.Ctx) error {
	if c.Locals("role") != "admin" {
		return c.Status(403).JSON(fiber.Map{
			"error": "Forbidden ",
		})
	}
	return c.Next()
}