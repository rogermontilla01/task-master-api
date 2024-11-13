package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type ApiConfig struct {
	Port string `mapstructure:"PORT"`
}

func NewApiConfig() ApiConfig {
	err := godotenv.Load()
	if err != nil {
		log.Info().Msg(".env file not found")
	}

	return ApiConfig{
		Port: os.Getenv("PORT"),
	}
}
