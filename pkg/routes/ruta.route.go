package routes

import (
	"net/http"
	"p1-backend/api/pkg/api"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func InitRutaRoutes(router *mux.Router, db *gorm.DB, authorizeRequest func(next http.Handler, tokened bool) http.Handler) {
	rutaApi := api.InitializeRutaApi(db)
	rutaRouter := router.PathPrefix("/ruta").Subrouter()
	rutaRouter.Handle("/create", authorizeRequest(rutaApi.CreateRuta(), true)).Methods("POST")
	rutaRouter.Handle("/change_state", authorizeRequest(rutaApi.ChangeStateRuta(), true)).Methods("POST")
}
