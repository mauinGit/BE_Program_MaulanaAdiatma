package controller

import (
	"GDGBatch2026/database"
	"GDGBatch2026/model"
	"GDGBatch2026/util"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	username := c.FormValue("username")
	email := c.FormValue("email")
	password := c.FormValue("password")

	if username == "" || email == "" || password == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Username, email, dan password wajib diisi",
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to hash password",
		})
	}

	user := model.User{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
		Role:     "user",
	}

	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Email sudah terdaftar",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Register berhasil",
	})
}

func Login(c *fiber.Ctx) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	if email == "" || password == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Email dan password wajib diisi",
		})
	}

	var user model.User
	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return c.Status(401).JSON(fiber.Map{
			"error": "Email atau password salah",
		})
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	); err != nil {
		return c.Status(401).JSON(fiber.Map{
			"error": "Email atau password salah",
		})
	}

	token, err := util.GenerateJWT(user.ID, user.Role)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to generate token",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Login berhasil",
		"token":   token,
	})
}

func Logout(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Logout berhasil",
	})
}
