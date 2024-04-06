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

type DestinoApi struct {
	service *service.DestinoService
}

func InitializeDestinoApi(db *gorm.DB) *DestinoApi {
	DestinoService := service.InitializeDestinoService(db)
	DestinoApi := DestinoApi{service: DestinoService}
	DestinoApi.service.Migrate()
	return &DestinoApi
}

func (api *DestinoApi) CreateDestino() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var modelRegister models.Destino

		decoder := json.NewDecoder(r.Body)
		fmt.Println(decoder)
		if err := decoder.Decode(&modelRegister); err != nil {
			ResponseWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		defer r.Body.Close()

		response := api.service.Create(&modelRegister)
		status := http.StatusOK

		if !response.IsSuccessfull {
			status = http.StatusBadGateway
		}

		RespondWithJSON(w, status, response)
	}
}
func (api *DestinoApi) UpdateDestino() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var Destino models.Destino
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&Destino); err != nil {
			ResponseWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		defer r.Body.Close()

		response := api.service.Update(&Destino)
		status := http.StatusOK

		if !response.IsSuccessfull {
			status = http.StatusBadGateway
		}

		RespondWithJSON(w, status, response)
	}
}

func (api *DestinoApi) GetAll() http.HandlerFunc {
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
func (api *DestinoApi) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, errorId := strconv.Atoi(r.URL.Query().Get("id"))
		if errorId != nil {
			ResponseWithError(w, http.StatusInternalServerError, errorId.Error())
			return
		}
		response := api.service.GetById(id)
		if !response.IsSuccessfull {
			ResponseWithError(w, http.StatusInternalServerError, response.Message)
			return
		}
		RespondWithJSON(w, http.StatusOK, response)
	}
}
