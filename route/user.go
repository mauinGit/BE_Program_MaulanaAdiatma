package route

import (
	"GDGBatch2026/controller"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(api fiber.Router) {
	user := api.Group("/user")
	user.Post("/register", controller.Register)
	user.Post("/login", controller.Login)
	user.Post("/logout", controller.Logout)
}