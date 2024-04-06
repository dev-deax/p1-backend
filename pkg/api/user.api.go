package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"p1-backend/api/pkg/models"
	service "p1-backend/api/pkg/services"
	"strconv"

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
		status := http.StatusOK

		if !response.IsSuccessfull {
			status = http.StatusBadGateway
		}

		RespondWithJSON(w, status, response)
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
			ResponseWithError(w, http.StatusNotFound, "Usuario no encontrado.")
			return
		}
		if !service.ValidatePassword(usuarioLogin.Password, user.Password) {
			ResponseWithError(w, http.StatusBadRequest, "La contraseña es incorrecta.")
			return
		}
		if !user.Activo {
			ResponseWithError(w, http.StatusUnauthorized, "El usuario no está activo.")
			return
		}
		tokenDto, errorToken := service.GenerateToken(user)
		if errorToken != nil {
			ResponseWithError(w, http.StatusInternalServerError, "No se pudo crear el token.")
			return
		}

		RespondWithJSON(w, http.StatusOK, tokenDto)
	}
}

func (api *UsuarioApi) GetAllUsers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		page, errorPage := strconv.Atoi(r.URL.Query().Get("page"))
		limit, errorLimit := strconv.Atoi(r.URL.Query().Get("limit"))

		if errorPage != nil || errorLimit != nil {
			// ResponseWithError(w, http.StatusBadRequest, "Invalid page or limit")
			// return
			page = 1
			limit = 10
		}
		response, err := api.service.GetAllUsers(page, limit)
		if err != nil {
			ResponseWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		RespondWithJSON(w, http.StatusOK, response)
	}
}

func (api *UsuarioApi) ChangeStateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var usuario models.Usuario
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&usuario); err != nil {
			ResponseWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		defer r.Body.Close()

		response := api.service.ChangeStateUser(int(usuario.ID), usuario.Activo)
		status := http.StatusOK

		if !response.IsSuccessfull {
			status = http.StatusBadGateway
		}

		RespondWithJSON(w, status, response)
	}
}
