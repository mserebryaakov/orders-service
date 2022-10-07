package http

import "github.com/go-chi/chi"

type Handler interface {
	Register(router *chi.Mux)
}
