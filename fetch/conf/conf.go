package conf

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort  string
	DebugPort   string
	SecretKey   string
	CurrencyURL string
	ResourceURL string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Print(".env notfound")
	}

	return &Config{
		ServerPort:  os.Getenv("SERVER_PORT"),
		DebugPort:   os.Getenv("DEBUG_PORT"),
		SecretKey:   os.Getenv("JWT_SECRET_KEY"),
		CurrencyURL: os.Getenv("CURRENCY_URL"),
		ResourceURL: os.Getenv("RESOURCE_URL"),
	}
}
