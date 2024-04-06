package service

import (
	"p1-backend/api/pkg/models"

	"gorm.io/gorm"
)

type DestinoService struct {
	db *gorm.DB
}

func InitializeDestinoService(db *gorm.DB) *DestinoService {
	return &DestinoService{db: db}
}

func (service *DestinoService) Migrate() error {
	return service.db.AutoMigrate(&models.Destino{})
}

func (service *DestinoService) Create(create *models.Destino) *models.ResponseMessage {
	err := service.db.Save(&create).Error
	if err != nil {
		return &models.ResponseMessage{IsSuccessfull: false, Message: err.Error()}
	}
	return &models.ResponseMessage{IsSuccessfull: true, Message: "Destino registrado exitosamente"}
}
func (service *DestinoService) Update(update *models.Destino) *models.ResponseMessage {

	err := service.db.Save(&update).Error
	if err != nil {
		return &models.ResponseMessage{IsSuccessfull: false, Message: err.Error()}
	}
	return &models.ResponseMessage{IsSuccessfull: true, Message: "Destino registrado exitosamente"}
}
func (service *DestinoService) GetById(id int) *models.ResponseMessage {
	var Destino models.Destino
	err := service.db.First(&Destino, id).Error
	if err != nil {
		return &models.ResponseMessage{IsSuccessfull: false, Message: err.Error()}
	}
	return &models.ResponseMessage{IsSuccessfull: true, Message: "Destino obtenida exitosamente", Data: Destino}
}
func (service *DestinoService) GetAll(page int, limit int) *models.ResponseMessage {
	var Destinos []models.Destino
	offset := (page - 1) * limit
	err := service.db.Limit(limit).Offset(offset).Find(&Destinos).Error
	if err != nil {
		return &models.ResponseMessage{IsSuccessfull: false, Message: err.Error()}
	}
	return &models.ResponseMessage{IsSuccessfull: true, Message: "Destinos obtenidas exitosamente", Data: Destinos}
}
