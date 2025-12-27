package route

import (
	"GDGBatch2026/controller"

	"github.com/gofiber/fiber/v2"
)

func EventRoutes(api fiber.Router) {
	event := api.Group("/event")
	event.Post("/", controller.CreateEvent)
	event.Get("/", controller.GetEvent)
}