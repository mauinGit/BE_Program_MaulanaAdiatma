package controller

import (
	"GDGBatch2026/model"
	"GDGBatch2026/util"
	"GDGBatch2026/database"

	"github.com/gofiber/fiber/v2"
	"strconv"
	"time"
)

func CreateEvent(c *fiber.Ctx) error {
	var event model.Event

	event.Judul = c.FormValue("judul")
	if event.Judul == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Judul wajib diisi",
		})
	}

	_, err := c.FormFile("cover")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Cover wajib diupload",
		})
	}

	event.Deskripsi = c.FormValue("deskripsi")
	if event.Deskripsi == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Deskripsi wajib diisi",
		})
	}

	if tanggalStr := c.FormValue("tanggal"); tanggalStr != "" {
		t, err := time.Parse(time.RFC3339, tanggalStr)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Invalid tanggal format. Use ISO 8601 (RFC3339)",
			})
		}
		event.Tanggal = t
	}

	capacity, err := strconv.Atoi(c.FormValue("capacity"))
	if err != nil || capacity <= 0 {
		return c.Status(400).JSON(fiber.Map{
			"error": "Capacity harus angka lebih dari 0",
		})
	}
	event.Capacity = capacity
	event.RemainingTicket = capacity

	price, err := strconv.Atoi(c.FormValue("price"))
	if err != nil || price <= 0 {
		return c.Status(400).JSON(fiber.Map{
			"error": "Price harus angka lebih dari 0",
		})
	}
	event.Price = price

	cover, err := util.SaveFile(c, "cover")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to save cover",
		})
	}
	event.Cover = cover

	if err := database.DB.Create(&event).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to save event",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Event created successfully",
		"event":   event,
	})
}

func GetEvent(c *fiber.Ctx) error {
	var event []model.Event

	if err := database.DB.Find(&event).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to get event",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Success get event",
		"data":    event,
	})
}
