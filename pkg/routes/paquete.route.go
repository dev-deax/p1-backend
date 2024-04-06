package routes

import (
	"net/http"
	"p1-backend/api/pkg/api"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func InitPaqueteRoutes(router *mux.Router, db *gorm.DB, authorizeRequest func(next http.Handler, tokened bool) http.Handler) {
	PaqueteApi := api.InitializePaqueteApi(db)
	PaqueteRouter := router.PathPrefix("/Paquete").Subrouter()
	PaqueteRouter.Handle("/create_factura", authorizeRequest(PaqueteApi.CrearFactura(), true)).Methods("POST")
	PaqueteRouter.Handle("/id", authorizeRequest(PaqueteApi.GetPaqueteByID(), true)).Methods("GET")
	PaqueteRouter.Handle("/all_facturas", authorizeRequest(PaqueteApi.GetFacturaAll(), true)).Methods("GET")
}
