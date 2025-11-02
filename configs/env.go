package configs

import "os"

func SetEnvVariables() string {
	return os.Getenv("MONGODB_CONNECTION_STRING")
}
