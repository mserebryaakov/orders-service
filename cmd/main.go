package main

import (
	"orders-service/config"
	"orders-service/internal/app"
	"orders-service/pkg/logger"
)

// Swagger документация
// @title           Order service
// @version         1.0
// @description     API Server for orders usecase

// @host      localhost:8000
// @BasePath  /
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
