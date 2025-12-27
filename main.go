package main

import (
	"GDGBatch2026/config"
	"GDGBatch2026/database"
	"GDGBatch2026/route"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.ENVInit()
	database.DBInit()
	database.DBMigrate()
	
	app := fiber.New()

	route.RouteList(app)

	log.Fatal(app.Listen(":3000"))
}
