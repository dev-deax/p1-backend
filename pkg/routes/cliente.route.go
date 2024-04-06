package routes

import (
	"net/http"
	"p1-backend/api/pkg/api"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func InitClienteRoutes(router *mux.Router, db *gorm.DB, authorizeRequest func(next http.Handler, tokened bool) http.Handler) {
	clienteApi := api.InitializeClienteApi(db)
	usersRouter := router.PathPrefix("/cliente").Subrouter()
	usersRouter.Handle("/create", authorizeRequest(clienteApi.Register(), true)).Methods("POST")
	usersRouter.Handle("/id", authorizeRequest(clienteApi.GetById(), true)).Methods("GET")
	usersRouter.Handle("/list", authorizeRequest(clienteApi.GetAll(), true)).Methods("GET")
}
