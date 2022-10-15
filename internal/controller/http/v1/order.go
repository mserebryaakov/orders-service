package v1

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	handlers "orders-service/internal/controller/http"
	apperror "orders-service/internal/controller/http/apperror"
	"orders-service/internal/domain/order"
	service "orders-service/internal/services/order"
	"orders-service/pkg/logger"

	"github.com/gin-gonic/gin"
)

const (
	orderURL = "/v1/order"
)

type ordersHandler struct {
	log          *logger.Logger
	orderUseCase *service.OrderUseCase
}

func NewOrdersHandler(log *logger.Logger, orderUseCase *service.OrderUseCase) handlers.Handler {
	return &ordersHandler{
		log:          log,
		orderUseCase: orderUseCase,
	}
}

// Регистрация эндпоинтов для работы с заказами
func (h *ordersHandler) Register(router *gin.Engine) {
	router.POST(orderURL, h.CreateOrder)
	router.GET(orderURL, h.GetList)
	router.PUT(orderURL, h.UpdateOrder)
	router.DELETE(orderURL, h.DeleteOrder)
}

// Получение заказа по id
// @Summary      Get order
// @Tags order
// @Description  Get order by ID
// @Produce      application/json
// @Param        id query string true "Order ID"
// @Success      200  {object}  order.Order "Success get order"
// @Failure		 400  {object}  UploadResponse "Invalid parameters"
// @Failure		 404  {object}  UploadResponse "Order not found"
// @Failure		 500  {object}  UploadResponse "Server error"
// @Router       /v1/order [get]
func (h *ordersHandler) GetList(c *gin.Context) {
	// Получение id из параметра запроса
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, &UploadResponse{
			Msg: "Id not found",
		})
		return
	}

	ctx := context.Background()

	// Поиск заказа usecase
	order, err := h.orderUseCase.FindOne(ctx, id)
	if err != nil {
		if err == apperror.ErrNotFound {
			c.Writer.WriteHeader(404)
		} else {
			c.JSON(http.StatusInternalServerError, &UploadResponse{
				Msg: err.Error(),
			})
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
// @Failure		 400  {object}  UploadResponse "Invalid body"
// @Failure		 500  {object}  UploadResponse "Server error"
// @Router       /v1/order [post]
func (h *ordersHandler) CreateOrder(c *gin.Context) {
	// Чтение body
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, &UploadResponse{
			Msg: err.Error(),
		})
		return
	}

	var order order.Order

	// Десериализация из json
	err = json.Unmarshal(body, &order)
	if err != nil {
		c.JSON(http.StatusBadRequest, &UploadResponse{
			Msg: err.Error(),
		})
		return
	}

	ctx := context.Background()

	// Создание заказа usecase
	id, err := h.orderUseCase.CreateItem(ctx, order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &UploadResponse{
			Msg: err.Error(),
		})
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
// @Failure		 400  {object}  UploadResponse "Invalid body"
// @Failure		 404  {object}  UploadResponse "Order not found"
// @Failure		 500  {object}  UploadResponse "Server error"
// @Router       /v1/order [put]
func (h *ordersHandler) UpdateOrder(c *gin.Context) {
	// Чтение body
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, &UploadResponse{
			Msg: err.Error(),
		})
		return
	}

	var order order.Order

	// Десериализация из json
	err = json.Unmarshal(body, &order)
	if err != nil {
		c.JSON(http.StatusBadRequest, &UploadResponse{
			Msg: err.Error(),
		})
		return
	}

	ctx := context.Background()

	// Обновление заказа usecase
	err = h.orderUseCase.Update(ctx, order)
	if err != nil {
		if err == apperror.ErrNotFound {
			c.Writer.WriteHeader(404)
		} else {
			c.JSON(http.StatusInternalServerError, &UploadResponse{
				Msg: err.Error(),
			})
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
// @Failure		 400  {object}  UploadResponse "Invalid parameters"
// @Failure		 404  {object}  UploadResponse "Order not found"
// @Failure		 500  {object}  UploadResponse "Server error"
// @Router       /v1/order [delete]
func (h *ordersHandler) DeleteOrder(c *gin.Context) {
	// Получение id из параметра запроса
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, &UploadResponse{
			Msg: "Id not found",
		})
		return
	}

	ctx := context.Background()

	// Удаление заказа usecase
	err := h.orderUseCase.Delete(ctx, id)
	if err != nil {
		if err == apperror.ErrNotFound {
			c.Writer.WriteHeader(404)
		} else {
			c.JSON(http.StatusInternalServerError, &UploadResponse{
				Msg: err.Error(),
			})
		}
		return
	}

	h.log.Info("Delete order")

	c.Writer.WriteHeader(200)
}
