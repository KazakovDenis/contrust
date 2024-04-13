package contrad

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerPort   string
	DatabaseURI  string
	DatabaseName string
}

var Config *AppConfig

func init() {
	Config = NewConfig()
}

func NewConfig() *AppConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return &AppConfig{
		ServerPort:   getEnv("SERVER_PORT", "8080"),
		DatabaseURI:  buildDbURI(),
		DatabaseName: getEnv("MONGO_DATABASE", "contra"),
	}
}

func getEnv(key string, defaultVal string) string {
	val := os.Getenv(key)
	if len(val) == 0 {
		return defaultVal
	}
	return val
}

func buildDbURI() string {
	user := getEnv("MONGO_INITDB_ROOT_USERNAME", "contra")
	password := getEnv("MONGO_INITDB_ROOT_PASSWORD", "contra")
	host := getEnv("MONGO_HOST", "localhost")
	port := getEnv("MONGO_PORT", "27017")
	return fmt.Sprintf(`mongodb://%s:%s@%s:%s/?authSource=admin`, user, password, host, port)
}
