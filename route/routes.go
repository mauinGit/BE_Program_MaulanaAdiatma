package route

import "github.com/gofiber/fiber/v2"

func RouteList(app *fiber.App) {
	api := app.Group("/api")
	EventRoutes(api)
}