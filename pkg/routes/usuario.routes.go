package routes

import (
	"net/http"
	"p1-backend/api/pkg/api"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)
 
func InitUsuarioRoutes(router *mux.Router,db *gorm.DB, authorizeRequest func(next http.Handler, tokened bool) http.Handler) {
	userApi := api.InitializeUsuarioApi(db)
	
	router.Handle("/login", authorizeRequest(userApi.Login(), false)).Methods("POST")
	router.Handle("/register", authorizeRequest(userApi.RegisterUser(), false)).Methods("POST")

}
