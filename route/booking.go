package route

import (
	"GDGBatch2026/controller"
	"GDGBatch2026/middleware"

	"github.com/gofiber/fiber/v2"
)

func BookingRoutes(api fiber.Router) {
	booking := api.Group("/booking")
	booking.Post("/", middleware.AuthRequired, middleware.UserOnly, controller.CreateBooking)
	booking.Get("/me", middleware.AuthRequired, middleware.UserOnly, controller.GetMyBookings)
	booking.Get("/", middleware.AuthRequired, middleware.AdminOnly, controller.GetAllBookings)
	booking.Get("/:event_id", middleware.AuthRequired, middleware.AdminOnly, controller.GetBookingsByEventID)
	booking.Delete("/:id", middleware.AuthRequired, middleware.UserOnly, controller.CancelMyBooking)
}