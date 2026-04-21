package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Key string
}

func New() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Ошибка при загрузке .env, %v", err)
	}
	key := os.Getenv("KEY")
	return &Config{
		Key: key,
	}
}
