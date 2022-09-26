package v1

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	apperror "orders-service/internal/controller/http"
	"orders-service/internal/domain/order"
)

func (h *handler) GetList(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Id not found", http.StatusBadRequest)
		return
	}

	ctx := context.Background()

	order, err := h.orderUseCase.FindOne(ctx, id)
	if err != nil {
		if err == apperror.ErrNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	output, err := json.Marshal(order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var order order.Order
	err = json.Unmarshal(body, &order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := context.Background()

	id, err := h.orderUseCase.CreateItem(ctx, order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write([]byte(id))

	h.log.Info("Create order")
}

func (h *handler) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var order order.Order
	err = json.Unmarshal(body, &order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := context.Background()

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

func (h *handler) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Id not found", http.StatusBadRequest)
		return
	}

	ctx := context.Background()

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
