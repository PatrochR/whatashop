package router

import (
	"fmt"
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

func (r *Router) Run() error {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Route("/user", func(subRouter chi.Router) {
		subRouter.Get("/", r.userHandler.GetAll)
		subRouter.Post("/", r.userHandler.Add)
	})
	log.Info(fmt.Sprintln("server start on port ", r.Address))
	return http.ListenAndServe(r.Address, router)
}
