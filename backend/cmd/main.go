package main

import (
	"log"
	"net/http"

	"github.com/KornutaKM/fullstack-dashboard/internal/config"
	"github.com/KornutaKM/fullstack-dashboard/internal/handlers"
	"github.com/KornutaKM/fullstack-dashboard/internal/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// 1. Загружаем конфиг
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 2. Создаём сервисы
	weatherService := services.NewWeatherService(cfg.WeatherAPIKey)
	newsService := services.NewNewsService(cfg.NewsAPIKey)
	exchangeService := services.NewExchangeService(cfg.ExchangeAPIKey)

	// 3. Создаём хендлеры
	weatherHandler := handlers.NewWeatherHandler(weatherService)
	newsHandler := handlers.NewNewsHandler(newsService)
	exchangeHandler := handlers.NewExchangeHandler(exchangeService)

	// 4. Настраиваем роутер
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// 5. Роуты
	r.Get("/api/weather", weatherHandler.GetWeather)
	r.Get("/api/news", newsHandler.GetNews)
	r.Get("/api/exchange", exchangeHandler.GetExchange)

	// 6. Запускаем сервер
	log.Println("🚀 Server starting on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Server failed: %v", err)
	}

}
