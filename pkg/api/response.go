package api

import (
	"encoding/json"
	"net/http"
)

type ErrorResponseDTO struct {
	Code    int      `json:"code"`
	Status  string   `json:"status"`
	Message string   `json:"message"`
	Errors  []string `json:"errors"`
}

func ResponseWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, ErrorResponseDTO{Code: code, Status: "Error", Message: message})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
