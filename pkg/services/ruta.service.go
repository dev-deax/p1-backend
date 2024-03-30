package service

import (
	"p1-backend/api/pkg/models"

	"gorm.io/gorm"
)

type RutaService struct {
	db *gorm.DB
}

func InitializeRutaService(db *gorm.DB) *RutaService {
	return &RutaService{db: db}
}

func (service *RutaService) Migrate() error {
	return service.db.AutoMigrate(&models.Ruta{})
}

func (service *RutaService) Create(create *models.Ruta) *models.ResponseMessage {
	ruta := models.Ruta{
		Nombre:    create.Nombre,
		Capacidad: create.Capacidad,
		DestinoID: create.DestinoID,
	}
	err := service.db.Save(&ruta).Error
	if err != nil {
		return &models.ResponseMessage{IsSuccessfull: false, Message: err.Error()}
	}
	return &models.ResponseMessage{IsSuccessfull: true, Message: "Ruta registrado exitosamente"}
}
func (service *RutaService) Update(update *models.Ruta) *models.ResponseMessage {
	ruta := models.Ruta{
		Nombre:    update.Nombre,
		Capacidad: update.Capacidad,
		DestinoID: update.DestinoID,
	}
	err := service.db.Save(&ruta).Error
	if err != nil {
		return &models.ResponseMessage{IsSuccessfull: false, Message: err.Error()}
	}
	return &models.ResponseMessage{IsSuccessfull: true, Message: "Ruta registrado exitosamente"}
}
func (service *RutaService) GetById(id int) *models.ResponseMessage {
	var ruta models.Ruta
	err := service.db.First(&ruta, id).Error
	if err != nil {
		return &models.ResponseMessage{IsSuccessfull: false, Message: err.Error()}
	}
	return &models.ResponseMessage{IsSuccessfull: true, Message: "Ruta obtenida exitosamente", Data: ruta}
}
func (service *RutaService) GetByDestinoId(destinoID int) *models.ResponseMessage {
	var rutas []models.Ruta
	err := service.db.Where("destino_id = ?", destinoID).Preload("Paquete").Preload("Destino").Find(&rutas).Error
	if err != nil {
		return &models.ResponseMessage{IsSuccessfull: false, Message: err.Error()}
	}
	return &models.ResponseMessage{IsSuccessfull: true, Message: "Rutas obtenidas exitosamente", Data: rutas}
}
func (service *RutaService) GetAll(page int, limit int) *models.ResponseMessage {
	var rutas []models.Ruta
	offset := (page - 1) * limit
	err := service.db.Limit(limit).Offset(offset).Preload("Paquete").Preload("Destino").Find(&rutas).Error
	if err != nil {
		return &models.ResponseMessage{IsSuccessfull: false, Message: err.Error()}
	}
	return &models.ResponseMessage{IsSuccessfull: true, Message: "Rutas obtenidas exitosamente", Data: rutas}
}
func (service *RutaService) GetRutaById(id int) (*models.Ruta, error) {
	model := new(models.Ruta)
	err := service.db.Where(`id = ?`, id).First(&model).Error
	return model, err
}
func (service *RutaService) ChangeStateRuta(id int, activate bool) *models.ResponseMessage {
	model, err := service.GetRutaById(id)
	if err != nil {
		return &models.ResponseMessage{IsSuccessfull: false, Message: err.Error()}
	}
	model.Activo = activate
	err = service.db.Save(&model).Error
	if err != nil {
		return &models.ResponseMessage{IsSuccessfull: false, Message: err.Error()}
	}
	return &models.ResponseMessage{IsSuccessfull: true, Message: "Cambio de estado exitoso"}
}
