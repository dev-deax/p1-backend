package routes

import (
	"net/http"
	"p1-backend/api/pkg/api"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func InitDestinoRoutes(router *mux.Router, db *gorm.DB, authorizeRequest func(next http.Handler, tokened bool) http.Handler) {
	DestinoApi := api.InitializeDestinoApi(db)
	DestinoRouter := router.PathPrefix("/destino").Subrouter()
	DestinoRouter.Handle("/create", authorizeRequest(DestinoApi.CreateDestino(), true)).Methods("POST")
	DestinoRouter.Handle("/list", authorizeRequest(DestinoApi.GetAll(), true)).Methods("GET")
}
