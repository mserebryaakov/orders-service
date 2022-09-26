package v1

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"orders-service/internal/domain/order"
)

func (h *handler) GetList(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	ctx := context.Background()

	order, err := h.orderUseCase.FindOne(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	output, err := json.Marshal(order)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(output)

	h.log.Info("GetList")
}

func (h *handler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var order order.Order
	err = json.Unmarshal(body, &order)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	ctx := context.Background()

	h.orderUseCase.CreateItem(ctx, order)

	h.log.Info("Create order")
}

func (h *handler) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Update order"))
	h.log.Info("Update order")
}

func (h *handler) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(204)
	w.Write([]byte("Delete order"))
	h.log.Info("Delete order")
}
