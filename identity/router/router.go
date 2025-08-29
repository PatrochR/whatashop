package router

import (
	"net/http"

	"github.com/PatrochR/whatashop/handler"
	"github.com/charmbracelet/log"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Router struct {
	Address     string
	userHandler *handler.UserHandler
}

func NewRouter(Address string, userHandler *handler.UserHandler) *Router {
	return &Router{
		Address:     Address,
		userHandler: userHandler,
	}
}

func(r *Router) Run() error {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/admin/user/" , r.userHandler.GetAll)
	log.Info("server start on port " , r.Address)
	return http.ListenAndServe(r.Address, router)
}
