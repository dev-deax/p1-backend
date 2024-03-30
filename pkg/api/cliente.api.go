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

type ClienteApi struct {
	service *service.ClienteService
}

func InitializeClienteApi(db *gorm.DB) *ClienteApi {
	clienteService := service.InitializeClienteService(db)
	clienteApi := ClienteApi{service: clienteService}
	clienteApi.service.Migrate()
	return &clienteApi
}

func (api *ClienteApi) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var clienteRegister models.Cliente

		decoder := json.NewDecoder(r.Body)
		fmt.Println(decoder)
		if err := decoder.Decode(&clienteRegister); err != nil {
			ResponseWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		defer r.Body.Close()

		response := api.service.Register(&clienteRegister)

		RespondWithJSON(w, http.StatusOK, response)
	}
}
func (api *ClienteApi) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		page, errorPage := strconv.Atoi(r.URL.Query().Get("page"))
		limit, errorLimit := strconv.Atoi(r.URL.Query().Get("limit"))

		if errorPage != nil || errorLimit != nil {
			page = 1
			limit = 10
		}
		response := api.service.GetAll(page, limit)
		if !response.IsSuccessfull {
			ResponseWithError(w, http.StatusInternalServerError, response.Message)
			return
		}
		RespondWithJSON(w, http.StatusOK, response)
	}
}
