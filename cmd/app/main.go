package main

import (
	"orders-service/config"
	"orders-service/internal/app"
	"orders-service/pkg/logger"
)

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
