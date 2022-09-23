package v1

import (
	"net/http"
	handlers "orders-service/internal/controller/http"
	service "orders-service/internal/services/order"
	"orders-service/pkg/logger"

	"github.com/julienschmidt/httprouter"
)

const (
	orderURL = "/v1/order:uuid"
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

func (h *handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodPost, orderURL, h.CreateOrder)
	router.HandlerFunc(http.MethodGet, orderURL, h.GetList)
	router.HandlerFunc(http.MethodPut, orderURL, h.UpdateOrder)
	router.HandlerFunc(http.MethodDelete, orderURL, h.DeleteOrder)
}
