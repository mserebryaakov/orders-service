package main

import (
	"orders-service/config"
	"orders-service/internal/app"
	"orders-service/pkg/logger"
)

// @title Order API
// @version 1.0
// @description API for Order service

// @host localhost:8000
// @BasePath /v1/order

func main() {
	// Получение логгера
	log := logger.GetLogger()

	// Получение конфига
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configs: %v", err)
	}

	// Запуск сервиса
	app.Start(cfg, log)
}
