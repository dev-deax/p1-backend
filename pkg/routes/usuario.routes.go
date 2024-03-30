package routes

import (
	"net/http"
	"p1-backend/api/pkg/api"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func InitUsuarioRoutes(router *mux.Router, db *gorm.DB, authorizeRequest func(next http.Handler, tokened bool) http.Handler) {
	userApi := api.InitializeUsuarioApi(db)
	router.Handle("/login", authorizeRequest(userApi.Login(), false)).Methods("POST")
	usersRouter := router.PathPrefix("/usuario").Subrouter()
	usersRouter.Handle("/register", authorizeRequest(userApi.RegisterUser(), true)).Methods("POST")
	usersRouter.Handle("/change_state", authorizeRequest(userApi.ChangeStateUser(), true)).Methods("POST")
	usersRouter.Handle("/list", authorizeRequest(userApi.GetAllUsers(), true)).Methods("GET")

}
