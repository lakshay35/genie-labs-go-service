package environment

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if os.Getenv("GIN_MODE") != "release" {
		loadDotenv()
	}
}

func loadDotenv() {
	err := godotenv.Load()
	if err != nil {
		log.Print("Error loading .env file")
	}
}

// GetEnvVariable ...
// Get environment variable
func GetEnvVariable(name string) string {
	return os.Getenv(name)
}
