package configs

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

func SetEnvVariables() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading dotenv")
	}

	return os.Getenv("MONGODB_CONNECTION_STRING")
}
