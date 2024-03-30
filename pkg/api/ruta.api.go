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

type RutaApi struct {
	service *service.RutaService
}

func InitializeRutaApi(db *gorm.DB) *RutaApi {
	rutaService := service.InitializeRutaService(db)
	rutaApi := RutaApi{service: rutaService}
	rutaApi.service.Migrate()
	return &rutaApi
}

func (api *RutaApi) CreateRuta() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var modelRegister models.Ruta

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
func (api *RutaApi) UpdateRuta() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var ruta models.Ruta
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&ruta); err != nil {
			ResponseWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		defer r.Body.Close()

		response := api.service.Update(&ruta)
		status := http.StatusOK

		if !response.IsSuccessfull {
			status = http.StatusBadGateway
		}

		RespondWithJSON(w, status, response)
	}
}
func (api *RutaApi) ChangeStateRuta() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var ruta models.Ruta
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&ruta); err != nil {
			ResponseWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		defer r.Body.Close()

		response := api.service.ChangeStateRuta(ruta.ID, ruta.Activo)
		status := http.StatusOK

		if !response.IsSuccessfull {
			status = http.StatusBadGateway
		}

		RespondWithJSON(w, status, response)
	}
}
func (api *RutaApi) GetAll() http.HandlerFunc {
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
func (api *RutaApi) GetByDestino() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		destino, errorDestino := strconv.Atoi(r.URL.Query().Get("destino"))
		if errorDestino != nil {
			ResponseWithError(w, http.StatusInternalServerError, errorDestino.Error())
			return
		}
		response := api.service.GetByDestinoId(destino)
		if !response.IsSuccessfull {
			ResponseWithError(w, http.StatusInternalServerError, response.Message)
			return
		}
		RespondWithJSON(w, http.StatusOK, response)
	}
}
func (api *RutaApi) GetById() http.HandlerFunc {
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
