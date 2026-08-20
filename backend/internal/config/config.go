package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	WeatherAPIKey string
	NewsAPIKey    string
	ExchangeAPIKey string
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("failed to load .env: %w", err)
	}

	return &Config{
		WeatherAPIKey: os.Getenv("WEATHER_API_KEY"),
		NewsAPIKey:    os.Getenv("NEWS_API_KEY"),
		ExchangeAPIKey:    os.Getenv("EXCHANGE_API_KEY"),
	}, nil
}
