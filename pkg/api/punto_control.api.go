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

type PuntoControlApi struct {
	service *service.PuntoControlService
}

func InitializePuntoControlApi(db *gorm.DB) *PuntoControlApi {
	PuntoControlService := service.InitializePuntoControlService(db)
	PuntoControlApi := PuntoControlApi{service: PuntoControlService}
	PuntoControlApi.service.Migrate()
	return &PuntoControlApi
}
func (api *PuntoControlApi) CreatePuntoControl() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var modelRegister models.PuntoControl

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
func (api *PuntoControlApi) UpdatePuntoControl() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var PuntoControl models.PuntoControl
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&PuntoControl); err != nil {
			ResponseWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		defer r.Body.Close()

		response := api.service.Update(&PuntoControl)
		status := http.StatusOK

		if !response.IsSuccessfull {
			status = http.StatusBadGateway
		}

		RespondWithJSON(w, status, response)
	}
}
func (api *PuntoControlApi) GetByIdPuntoControl() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var PuntoControl models.PuntoControl
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&PuntoControl); err != nil {
			ResponseWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		defer r.Body.Close()

		response := api.service.GetById(PuntoControl.ID)
		status := http.StatusOK

		if !response.IsSuccessfull {
			status = http.StatusBadGateway
		}

		RespondWithJSON(w, status, response)
	}
}
func (api *PuntoControlApi) GetAll() http.HandlerFunc {
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
func (api *PuntoControlApi) ChangeStateRuta() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var model models.PuntoControl
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&model); err != nil {
			ResponseWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		defer r.Body.Close()

		response := api.service.ChangeStatePuntoControl(model.ID, model.Activo)
		status := http.StatusOK

		if !response.IsSuccessfull {
			status = http.StatusBadGateway
		}
		RespondWithJSON(w, status, response)
	}
}
func (api *PuntoControlApi) ProcesarPaquete(salida bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var model models.PaquetesPuntosControl
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&model); err != nil {
			ResponseWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		defer r.Body.Close()

		response := api.service.ProcesarPaquete(model.PaqueteID, model.PuntoControlID, salida)
		status := http.StatusOK

		if !response.IsSuccessfull {
			status = http.StatusBadGateway
		}
		RespondWithJSON(w, status, response)
	}
}
func (api *PuntoControlApi) GetPaquetesPuntosControlByUsuario() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var model models.Usuario
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&model); err != nil {
			ResponseWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		defer r.Body.Close()

		response := api.service.GetPaquetesPuntosControlByUsuario(model.ID)
		status := http.StatusOK

		if !response.IsSuccessfull {
			status = http.StatusBadGateway
		}
		RespondWithJSON(w, status, response)
	}
}
func (api *PuntoControlApi) GetCostoByPuntoControl() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var model models.PuntoControl
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&model); err != nil {
			ResponseWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		defer r.Body.Close()

		response := api.service.GetCostoPaquetesPuntosControlByPuntoControl(model.ID)
		status := http.StatusOK

		if !response.IsSuccessfull {
			status = http.StatusBadGateway
		}
		RespondWithJSON(w, status, response)
	}
}
