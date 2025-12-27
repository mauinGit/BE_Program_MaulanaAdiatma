package controller

import (
	"GDGBatch2026/database"
	"GDGBatch2026/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateBooking(c *fiber.Ctx) error {
	userIDValue := c.Locals("user_id")
	if userIDValue == nil {
		return c.Status(401).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	userID := userIDValue.(uint)

	eventIDInput := c.FormValue("event_id")
	quantityInput := c.FormValue("quantity")

	if eventIDInput == "" || quantityInput == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "event_id dan quantity wajib diisi",
		})
	}

	eventID, err := strconv.Atoi(eventIDInput)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "event_id harus berupa angka",
		})
	}

	quantity, err := strconv.Atoi(quantityInput)
	if err != nil || quantity <= 0 {
		return c.Status(400).JSON(fiber.Map{
			"error": "quantity harus berupa angka lebih dari 0",
		})
	}

	err = database.DB.Transaction(func(transaction *gorm.DB) error {
		var event model.Event

		if err := transaction.
			Clauses(clause.Locking{Strength: "UPDATE"}).
			First(&event, eventID).Error; err != nil {
			return err
		}

		if event.RemainingTicket < quantity {
			return gorm.ErrInvalidData
		}

		totalPrice := quantity * event.Price

		booking := model.Booking{
			UserID:     userID,
			EventID:    event.ID,
			Quantity:   quantity,
			TotalPrice: totalPrice,
		}

		if err := transaction.Create(&booking).Error; err != nil {
			return err
		}

		event.RemainingTicket = event.RemainingTicket - quantity
		if err := transaction.Save(&event).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		if err == gorm.ErrInvalidData {
			return c.Status(400).JSON(fiber.Map{
				"error": "stok tiket tidak cukup",
			})
		}

		return c.Status(400).JSON(fiber.Map{
			"error": "event tidak ditemukan",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Booking berhasil",
	})
}

func GetMyBookings(c *fiber.Ctx) error {
	userIDValue := c.Locals("user_id")
	if userIDValue == nil {
		return c.Status(401).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}
	userID := userIDValue.(uint)

	var bookings []model.Booking

	if err := database.DB.
		Where("user_id = ?", userID).
		Order("created_at desc").
		Find(&bookings).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Gagal mengambil data booking",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Berhasil mengambil riwayat booking",
		"data":    bookings,
	})
}

func GetAllBookings(c *fiber.Ctx) error {
	var bookings []model.Booking

	if err := database.DB.
		Order("created_at desc").
		Find(&bookings).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Gagal mengambil data booking",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Berhasil mengambil semua booking",
		"data":    bookings,
	})
}

func GetBookingsByEventID(c *fiber.Ctx) error {
	eventIDParam := c.Params("event_id")
	eventID, err := strconv.Atoi(eventIDParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "event_id tidak valid",
		})
	}

	var event model.Event
	if err := database.DB.First(&event, eventID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Event tidak ditemukan",
		})
	}

	var bookings []model.Booking

	if err := database.DB.
		Where("event_id = ?", eventID).
		Order("created_at desc").
		Find(&bookings).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Gagal mengambil data booking",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Berhasil mengambil booking untuk event",
		"event": fiber.Map{
			"id":    event.ID,
			"judul": event.Judul,
		},
		"data": bookings,
	})
}

func CancelMyBooking(c *fiber.Ctx) error {
	userIDValue := c.Locals("user_id")
	if userIDValue == nil {
		return c.Status(401).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}
	userID := userIDValue.(uint)

	bookingIDParam := c.Params("id")
	bookingID, err := strconv.Atoi(bookingIDParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "ID booking tidak valid",
		})
	}

	err = database.DB.Transaction(func(transaction *gorm.DB) error {
		var booking model.Booking

		// ambil booking milik user
		if err := transaction.
			Where("id = ? AND user_id = ?", bookingID, userID).
			First(&booking).Error; err != nil {
			return err
		}

		var event model.Event
	
		if err := transaction.
			Clauses(clause.Locking{Strength: "UPDATE"}).
			First(&event, booking.EventID).Error; err != nil {
			return err
		}

		event.RemainingTicket = event.RemainingTicket + booking.Quantity
		if err := transaction.Save(&event).Error; err != nil {
			return err
		}

		if err := transaction.Delete(&booking).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Booking tidak ditemukan atau bukan milik Anda",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Booking berhasil dibatalkan",
	})
}