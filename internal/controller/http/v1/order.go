package v1

import (
	"context"
	"net/http"
	apperror "orders-service/internal/controller/http/apperror"
	"orders-service/internal/domain/order"
	service "orders-service/internal/services/order"
	"orders-service/pkg/logger"

	"github.com/gin-gonic/gin"
)

const (
	orderURL = "/order"
)

type ordersHandler struct {
	log          *logger.Logger
	orderUseCase *service.OrderUseCase
}

func NewOrdersHandler(log *logger.Logger, orderUseCase *service.OrderUseCase) *ordersHandler {
	return &ordersHandler{
		log:          log,
		orderUseCase: orderUseCase,
	}
}

// Регистрация эндпоинтов для работы с заказами
func (h *ordersHandler) Register(router *gin.Engine, uh *userHandler) {
	v1 := router.Group("/v1", uh.userIdentity)
	{
		v1.POST(orderURL, h.createOrder)
		v1.GET(orderURL, h.getList)
		v1.PUT(orderURL, h.updateOrder)
		v1.DELETE(orderURL, h.deleteOrder)
	}
}

// Получение заказа по id
// @Summary      Get order
// @Tags order
// @Description  Get order by ID
// @Produce      application/json
// @Param        id query string true "Order ID"
// @Success      200  {object}  order.Order "Success get order"
// @Failure		 400  {object}  errorResponse "Invalid parameters"
// @Failure		 404  {object}  errorResponse "Order not found"
// @Failure		 500  {object}  errorResponse "Server error"
// @Router       /v1/order [get]
func (h *ordersHandler) getList(c *gin.Context) {
	// Получение id из параметра запроса
	id := c.Query("id")
	if id == "" {
		newErrorResponse(c, http.StatusBadRequest, "Id not found")
		return
	}

	ctx := context.Background()

	// Поиск заказа usecase
	order, err := h.orderUseCase.FindOne(ctx, id)
	if err != nil {
		if err == apperror.ErrNotFound {
			newErrorResponse(c, http.StatusNotFound, "not found order")
		} else {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	h.log.Info("GetList")

	// Успешный ответ
	c.JSON(http.StatusOK, order)
}

// Создание заказа
// @Summary      Create order
// @Tags order
// @Description  Create order
// @Accept		 application/json
// @Produce      application/json
// @Param        order body order.Order true "Order object"
// @Success      200  {object}  IdResponse "Success create"
// @Failure		 400  {object}  errorResponse "Invalid body"
// @Failure		 500  {object}  errorResponse "Server error"
// @Router       /v1/order [post]
func (h *ordersHandler) createOrder(c *gin.Context) {
	var order order.Order

	// Валидация body
	if err := c.BindJSON(&order); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	ctx := context.Background()

	// Создание заказа usecase
	id, err := h.orderUseCase.CreateItem(ctx, order)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	h.log.Info("Create order")

	// Успешный ответ
	c.JSON(http.StatusOK, &IdResponse{
		Id: id,
	})
}

// Обновление заказа
// @Summary      Update order
// @Tags order
// @Description  Update order
// @Accept		 application/json
// @Produce      application/json
// @Param        order body order.Order true "Order object"
// @Success      200
// @Failure		 400  {object}  errorResponse "Invalid body"
// @Failure		 404  {object}  errorResponse "Order not found"
// @Failure		 500  {object}  errorResponse "Server error"
// @Router       /v1/order [put]
func (h *ordersHandler) updateOrder(c *gin.Context) {
	var order order.Order

	// Валидация body
	if err := c.BindJSON(&order); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	ctx := context.Background()

	// Обновление заказа usecase
	err := h.orderUseCase.Update(ctx, order)
	if err != nil {
		if err == apperror.ErrNotFound {
			newErrorResponse(c, http.StatusNotFound, err.Error())
		} else {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	h.log.Info("Update order")

	// Успешный ответ
	c.Writer.WriteHeader(200)
}

// Удаление заказа по id
// @Summary      Delete order
// @Tags order
// @Description  Delete order by ID
// @Param        id query string true "Order ID"
// @Success      200
// @Failure		 400  {object}  errorResponse "Invalid parameters"
// @Failure		 404  {object}  errorResponse "Order not found"
// @Failure		 500  {object}  errorResponse "Server error"
// @Router       /v1/order [delete]
func (h *ordersHandler) deleteOrder(c *gin.Context) {
	// Получение id из параметра запроса
	id := c.Query("id")
	if id == "" {
		newErrorResponse(c, http.StatusBadRequest, "Id not found")
		return
	}

	ctx := context.Background()

	// Удаление заказа usecase
	err := h.orderUseCase.Delete(ctx, id)
	if err != nil {
		if err == apperror.ErrNotFound {
			newErrorResponse(c, http.StatusNotFound, err.Error())
		} else {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	h.log.Info("Delete order")

	c.Writer.WriteHeader(200)
}
