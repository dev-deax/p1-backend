package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"p1-backend/api/pkg/models"
	service "p1-backend/api/pkg/services"

	"gorm.io/gorm"
)

type UsuarioApi struct {
	service *service.UsuarioService
}

func InitializeUsuarioApi(db *gorm.DB) *UsuarioApi {
	userService := service.InitializeUsuarioService(db)
	userApi := UsuarioApi{service: userService}
	userApi.service.Migrate()
	return &userApi
}

func (api *UsuarioApi) RegisterUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var usuarioRegister models.Usuario

		decoder := json.NewDecoder(r.Body)
		fmt.Println(decoder)
		if err := decoder.Decode(&usuarioRegister); err != nil {
			ResponseWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		defer r.Body.Close()

		response := api.service.RegisterUser(&usuarioRegister)

		RespondWithJSON(w, http.StatusOK, response)
	}
}

func (api *UsuarioApi) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var usuarioLogin models.Usuario

		decoder := json.NewDecoder(r.Body)

		if err := decoder.Decode(&usuarioLogin); err != nil {
			ResponseWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		defer r.Body.Close()

		user, err := api.service.GetUserByEmail(usuarioLogin.Email)
		if err != nil {
			ResponseWithError(w, http.StatusNotFound, "User not found")
			return
		}
		if !service.ValidatePassword(usuarioLogin.Password, user.Password) {
			ResponseWithError(w, http.StatusBadRequest, "Password is wrong")
			return
		}
		tokenDto, errorToken := service.GenerateToken(usuarioLogin)
		if errorToken != nil {
			ResponseWithError(w, http.StatusInternalServerError, "Could not create token")
			return
		}

		RespondWithJSON(w, http.StatusOK, tokenDto)
	}
}
