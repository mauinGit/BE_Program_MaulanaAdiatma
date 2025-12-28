package route

import (
	"GDGBatch2026/controller"
	"GDGBatch2026/middleware"

	"github.com/gofiber/fiber/v2"
)

func EventRoutes(api fiber.Router) {
	event := api.Group("/event")
	event.Post("/", middleware.AuthRequired, middleware.AdminOnly, controller.CreateEvent)
	event.Get("/", middleware.AuthRequired, controller.GetEvent)
	event.Get("/:id", middleware.AuthRequired, controller.GetEventByID)
	event.Put("/:id", middleware.AuthRequired, middleware.AdminOnly, controller.UpdateEvent)
	event.Delete("/:id", middleware.AuthRequired, middleware.AdminOnly, controller.DeleteEvent)
}