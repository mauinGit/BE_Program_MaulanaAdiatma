package main

import (
	"GDGBatch2026/config"
	"GDGBatch2026/database"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.ENVInit()
	database.DBInit()
	app := fiber.New()

	log.Fatal(app.Listen(":3000"))
}
