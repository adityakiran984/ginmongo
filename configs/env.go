package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func SetEnvVariables() string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(".env file not found.")
	}
	return os.Getenv("MONGODB_CONNECTION_STRING")
}
