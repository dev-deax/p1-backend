package models

type ResponseMessage struct {
	IsSuccessfull bool        `json:"is_successfull"`
	Message       string      `json:"message"`
	Data          interface{} `json:"data,omitempty"`
}
