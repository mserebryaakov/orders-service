package app

import (
	"context"
	"net/http"
	"orders-service/config"
	"orders-service/pkg/httpserver"
	"orders-service/pkg/logger"
	"os"
	"os/signal"
	"syscall"

	"github.com/julienschmidt/httprouter"
)

// Запуск сервиса
func Start(cfg config.Config, log *logger.Logger) {
	// Создание объекта сервера
	server := new(httpserver.Server)

	// Создание роутера
	router := httprouter.New()

	// Регистрация эндпоинтов
	router.HandlerFunc(http.MethodGet, "/hello", func(w http.ResponseWriter, r *http.Request) {
		log.Debug("hello, debug")
	})

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
