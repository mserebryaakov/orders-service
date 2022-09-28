package v1

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	apperror "orders-service/internal/controller/http"
	"orders-service/internal/domain/order"
)

// Получение заказа по id
// @Summary      Get order
// @Description  Get order by ID
// @Tags         Order
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Order ID"
// @Success      200  {object}  order.Order
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /v1/order{id} [get]
func (h *handler) GetList(w http.ResponseWriter, r *http.Request) {
	// Получение id из параметра запроса
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Id not found", http.StatusBadRequest)
		return
	}

	ctx := context.Background()

	// Поиск заказа usecase
	order, err := h.orderUseCase.FindOne(ctx, id)
	if err != nil {
		if err == apperror.ErrNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Преобразование Order в json
	output, err := json.Marshal(order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(output)

	h.log.Info("GetList")
}

// Создание заказа
func (h *handler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	// Чтение body
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var order order.Order

	// Десериализация из json
	err = json.Unmarshal(body, &order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := context.Background()

	// Создание заказа usecase
	id, err := h.orderUseCase.CreateItem(ctx, order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write([]byte(id))

	h.log.Info("Create order")
}

// Обновление заказа
func (h *handler) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	// Чтение body
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var order order.Order

	// Десериализация из json
	err = json.Unmarshal(body, &order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := context.Background()

	// Обновление заказа usecase
	err = h.orderUseCase.Update(ctx, order)
	if err != nil {
		if err == apperror.ErrNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(204)
	h.log.Info("Update order")
}

// Удаление заказа по id
func (h *handler) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	// Получение id из параметра запроса
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Id not found", http.StatusBadRequest)
		return
	}

	ctx := context.Background()

	// Удаление заказа usecase
	err := h.orderUseCase.Delete(ctx, id)
	if err != nil {
		if err == apperror.ErrNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(204)
	h.log.Info("Delete order")
}
