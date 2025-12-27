package route

import (
	"GDGBatch2026/controller"
	"GDGBatch2026/middleware"

	"github.com/gofiber/fiber/v2"
)

func BookingRoutes(api fiber.Router) {
	booking := api.Group("/booking")
	booking.Post("/", middleware.AuthRequired, controller.CreateBooking)
}