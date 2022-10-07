package v1

import (
	handlers "orders-service/internal/controller/http"
	service "orders-service/internal/services/order"
	"orders-service/pkg/logger"

	"github.com/go-chi/chi"
)

const (
	orderURL = "/v1/order"
)

type handler struct {
	log          *logger.Logger
	orderUseCase *service.OrderUseCase
}

func NewHandler(log *logger.Logger, orderUseCase *service.OrderUseCase) handlers.Handler {
	return &handler{
		log:          log,
		orderUseCase: orderUseCase,
	}
}

// Регистрация эндпоинтов для работы с заказами
func (h *handler) Register(router *chi.Mux) {
	router.Post(orderURL, h.CreateOrder)
	router.Get(orderURL, h.GetList)
	router.Put(orderURL, h.UpdateOrder)
	router.Delete(orderURL, h.DeleteOrder)
}
