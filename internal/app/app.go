package app

import (
	"context"
	"orders-service/config"
	_ "orders-service/docs"
	v1 "orders-service/internal/controller/http/v1"
	dbOrder "orders-service/internal/domain/order/mongodb"
	dbUser "orders-service/internal/domain/user/mongodb"
	orderService "orders-service/internal/services/order"
	userService "orders-service/internal/services/user"
	"orders-service/pkg/httpserver"
	"orders-service/pkg/logger"
	"orders-service/pkg/mongodb"
	"os"
	"os/signal"
	"syscall"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

// Запуск сервиса
func Start(cfg config.Config, log *logger.Logger) {
	// Создание mongodb клиента
	mongoDBClient, err := mongodb.NewClient(context.Background(), cfg.Db.Host, cfg.Db.Port, cfg.Db.Username,
		cfg.Db.Password, cfg.Db.Database, cfg.Db.AuthDB)
	if err != nil {
		panic(err)
	}

	//Репозиторий пользователей
	userRepos := dbUser.NewUserRepository(mongoDBClient, cfg.Db.CollectionUser, log)
	// Сервис пользователей
	userServices := userService.NewUserService(userRepos)
	// Handlers пользователей
	userHandler := v1.NewUserHandler(log, userServices)

	// Репозиторий заказов
	orderRepos := dbOrder.NewOrderRepository(mongoDBClient, cfg.Db.CollectionOrder, log)
	// Сервис заказов
	orderServices := orderService.NewOrderService(orderRepos)
	// Handlers заказов
	orderHandler := v1.NewOrdersHandler(log, orderServices)

	// Роутер
	router := gin.New()
	// Инициализация swagger docs
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// Регистрация user handlers
	userHandler.Register(router)
	// Регистрация orders handlers
	orderHandler.Register(router)

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
	log.Infof("Shutdown server, %s", oscall)

	if err := server.Shutdown(context.Background()); err != nil {
		log.Errorf("Error occured on server shutting down: %v", err)
	}
}
