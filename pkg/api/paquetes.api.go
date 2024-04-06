package api

import (
	"encoding/json"
	"net/http"
	"p1-backend/api/pkg/models"
	service "p1-backend/api/pkg/services"
	"strconv"

	"gorm.io/gorm"
)

type PaqueteApi struct {
	service *service.PaqueteService
}

func InitializePaqueteApi(db *gorm.DB) *PaqueteApi {
	PaqueteService := service.InitializePaqueteService(db)
	PaqueteApi := PaqueteApi{service: PaqueteService}
	PaqueteApi.service.Migrate()
	return &PaqueteApi
}

func (api *PaqueteApi) GetPaqueteByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id, error := strconv.Atoi(r.URL.Query().Get("id"))
		if error != nil {
			ResponseWithError(w, http.StatusBadRequest, error.Error())
			return
		}

		response := api.service.GetPaqueteByID(id)
		status := http.StatusOK

		if !response.IsSuccessfull {
			status = http.StatusBadGateway
		}

		RespondWithJSON(w, status, response)
	}
}

func (api *PaqueteApi) CrearFactura() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var factura models.Factura
		if err := json.NewDecoder(r.Body).Decode(&factura); err != nil {
			ResponseWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		defer r.Body.Close()

		response := api.service.CrearFactura(factura)
		status := http.StatusOK

		if !response.IsSuccessfull {
			status = http.StatusBadGateway
		}

		RespondWithJSON(w, status, response)
	}
}

func (api *PaqueteApi) GetFacturaAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		page, errorPage := strconv.Atoi(r.URL.Query().Get("page"))
		limit, errorLimit := strconv.Atoi(r.URL.Query().Get("limit"))

		if errorPage != nil || errorLimit != nil {
			page = 1
			limit = 10
		}
		response := api.service.GetFacturaAll(page, limit)
		status := http.StatusOK

		if !response.IsSuccessfull {
			status = http.StatusBadGateway
		}

		RespondWithJSON(w, status, response)
	}
}
