package routes

import (
	"github.com/CavnHan/wallet-rpc-service/services/rest/service"
	"github.com/go-chi/chi/v5"
)

type Routes struct {
	router *chi.Mux
	svc    service.Service
}

func NewRoutes(r *chi.Mux, svc service.Service) Routes {
	return Routes{
		router: r,
		svc:    svc,
	}
}
