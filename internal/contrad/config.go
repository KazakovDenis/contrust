package contrad

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	serverPort string
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return &Config{
		serverPort: getEnv("SERVER_PORT", "8080"),
	}
}

func getEnv(key string, defaultVal string) string {
	val := os.Getenv(key)
	if len(val) == 0 {
		return defaultVal
	}
	return val
}
