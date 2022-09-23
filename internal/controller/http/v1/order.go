package v1

import "net/http"

func (h *handler) GetList(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Get order by id"))
	h.log.Info("GetList")
}

func (h *handler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Create order"))
	h.log.Info("GetList")
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
