package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

func ENVInit() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	fmt.Println(".env Succesfully load")
}