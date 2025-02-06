package router

import (
	"GoProjects/pkg/controller"
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
)

type Router struct {
	Address         string
	Router          *mux.Router
	Controller      *controller.Controller
	QueryController *controller.QueryController
}

func New(addr string, controller *controller.Controller, queryController *controller.QueryController) *Router {
	slog.Info("Создаю роутер")
	router := &Router{
		Address:         addr,
		Router:          mux.NewRouter(),
		Controller:      controller,
		QueryController: queryController,
	}
	router.setupRoutes()
	slog.Info("Создал")
	return router
}

func (r *Router) setupRoutes() {
	slog.Info("Собираю все ручки для роутера")

	r.Router.HandleFunc("/book/{id}", r.QueryController.GetBookById).Methods(http.MethodGet)

	r.Router.HandleFunc("/books", r.QueryController.GetBooks).Methods(http.MethodGet)

	r.Router.HandleFunc("/book", r.Controller.CreateBook).Methods(http.MethodPost)

	r.Router.HandleFunc("/book/{id}", r.Controller.UpdateBook).Methods(http.MethodPut)

	r.Router.HandleFunc("/book/{id}", r.Controller.DeleteBook).Methods(http.MethodDelete)

	slog.Info("Собрал все ручки для роутера")
}

func (r *Router) ListenAndServe() error {
	slog.Info("Listening on " + r.Address)
	return http.ListenAndServe(r.Address, r.Router)
}
