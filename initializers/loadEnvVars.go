package initializers

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadSecrets() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("env not found")
		return
	}
}
