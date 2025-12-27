package util

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func SaveFile(c *fiber.Ctx, fieldName string) (string, error) {
	file, err := c.FormFile(fieldName)
	if err != nil {
		return "", err
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	allowedExt := map[string]bool{
		".png":  true,
		".jpg":  true,
		".jpeg": true,
		".webp": true,
	}

	if !allowedExt[ext] {
		return "", fmt.Errorf("invalid file type")
	}

	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)

	saveDir := "./assets"
	savePath := filepath.Join(saveDir, filename)

	if err := c.SaveFile(file, savePath); err != nil {
		return "", err
	}

	baseURL := os.Getenv("BASE_URL")
	fileURL := fmt.Sprintf("%s/assets/%s", baseURL, filename)

	return fileURL, nil
}