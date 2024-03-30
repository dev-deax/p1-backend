package service

import (
	"p1-backend/api/pkg/models"

	"gorm.io/gorm"
)

type TarifasService struct {
	db *gorm.DB
}

func InitializeTarifasService(db *gorm.DB) *TarifasService {
	return &TarifasService{db: db}
}
func (service *TarifasService) Create(create *models.Tarifas) *models.ResponseMessage {
	err := service.db.Save(&create).Error
	if err != nil {
		return &models.ResponseMessage{IsSuccessfull: false, Message: err.Error()}
	}
	return &models.ResponseMessage{IsSuccessfull: true, Message: "Registrado exitosamente"}
}
func (service *TarifasService) Update(update *models.Tarifas) *models.ResponseMessage {
	err := service.db.Save(&update).Error
	if err != nil {
		return &models.ResponseMessage{IsSuccessfull: false, Message: err.Error()}
	}
	return &models.ResponseMessage{IsSuccessfull: true, Message: "Actulizado exitosamente"}
}
func (service *TarifasService) GetById(id int) *models.ResponseMessage {
	var data models.Tarifas
	err := service.db.First(&data, id).Error
	if err != nil {
		return &models.ResponseMessage{IsSuccessfull: false, Message: err.Error()}
	}
	return &models.ResponseMessage{IsSuccessfull: true, Message: "Tarifas obtenida exitosamente", Data: data}
}
func (service *TarifasService) GetAll(page int, limit int) *models.ResponseMessage {
	var data []models.Tarifas
	offset := (page - 1) * limit
	err := service.db.Limit(limit).Offset(offset).Find(&data).Error
	if err != nil {
		return &models.ResponseMessage{IsSuccessfull: false, Message: err.Error()}
	}
	return &models.ResponseMessage{IsSuccessfull: true, Message: "Tarifas obtenidas exitosamente", Data: data}
}
