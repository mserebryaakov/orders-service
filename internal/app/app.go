package app

import (
	"context"
	"orders-service/config"
	v1 "orders-service/internal/controller/http/v1"
	db "orders-service/internal/domain/order/mongodb"
	service "orders-service/internal/services/order"
	"orders-service/pkg/httpserver"
	"orders-service/pkg/logger"
	"orders-service/pkg/mongodb"
	"os"
	"os/signal"
	"syscall"

	"github.com/julienschmidt/httprouter"
)

// Запуск сервиса
func Start(cfg config.Config, log *logger.Logger) {
	// Создание mongodb клиента
	mongoDBClient, err := mongodb.NewClient(context.Background(), cfg.Db.Host, cfg.Db.Port, cfg.Db.Username,
		cfg.Db.Password, cfg.Db.Database, cfg.Db.AuthDB)
	if err != nil {
		panic(err)
	}

	// Repository
	repos := db.New(mongoDBClient, cfg.Db.Collection, log)
	// Use case
	services := service.New(repos)

	// Создание роутера и регистрация эндпоинтов
	router := httprouter.New()
	handler := v1.NewHandler(log, services)
	handler.Register(router)

	// Создание объекта сервера
	server := new(httpserver.Server)

	// Запуск сервера
	go func() {
		if err := server.Run(cfg.Server.Port, router); err != nil {
			log.Fatal("Failed running server %v", err)
		}
	}()

	// Shutdown
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	oscall := <-interrupt
	log.Infof("app.Start() - signal, %s", oscall)

	if err := server.Shutdown(context.Background()); err != nil {
		log.Errorf("Error occured on server shutting down: %v", err)
	}
}
