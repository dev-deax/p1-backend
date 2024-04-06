package routes

import (
	"net/http"
	"p1-backend/api/pkg/api"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func InitPuntoControlRoutes(router *mux.Router, db *gorm.DB, authorizeRequest func(next http.Handler, tokened bool) http.Handler) {
	PuntoControlApi := api.InitializePuntoControlApi(db)
	PuntoControlRouter := router.PathPrefix("/PuntoControl").Subrouter()
	PuntoControlRouter.Handle("/create", authorizeRequest(PuntoControlApi.CreatePuntoControl(), true)).Methods("POST")
	PuntoControlRouter.Handle("/update", authorizeRequest(PuntoControlApi.UpdatePuntoControl(), true)).Methods("POST")
	PuntoControlRouter.Handle("/costo", authorizeRequest(PuntoControlApi.GetCostoByPuntoControl(), true)).Methods("POST")
	PuntoControlRouter.Handle("/change_state", authorizeRequest(PuntoControlApi.ChangeStateRuta(), true)).Methods("POST")
	PuntoControlRouter.Handle("/paquete_salida", authorizeRequest(PuntoControlApi.ProcesarPaquete(true), true)).Methods("POST")
	PuntoControlRouter.Handle("/paquete_entrada", authorizeRequest(PuntoControlApi.ProcesarPaquete(false), true)).Methods("POST")
	PuntoControlRouter.Handle("/all", authorizeRequest(PuntoControlApi.GetAll(), true)).Methods("GET")
	PuntoControlRouter.Handle("/id", authorizeRequest(PuntoControlApi.GetByIdPuntoControl(), true)).Methods("GET")
	PuntoControlRouter.Handle("/paquete_usuario", authorizeRequest(PuntoControlApi.GetPaquetesPuntosControlByUsuario(), true)).Methods("GET")
}
