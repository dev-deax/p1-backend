package service

import (
	"p1-backend/api/pkg/models"

	"gorm.io/gorm"
)

type PaqueteService struct {
	db *gorm.DB
}

type EstadoPaquete int

const (
	EN_BODEGA    EstadoPaquete = 1
	EN_RUTA      EstadoPaquete = 2
	EN_ENTREGADO EstadoPaquete = 3
)

func InitializePaqueteService(db *gorm.DB) *PaqueteService {
	return &PaqueteService{db: db}
}

func (service *PaqueteService) Migrate() error {
	return service.db.AutoMigrate(&models.Paquete{})
}

func (service *PaqueteService) GetPaqueteByID(paqueteID int) *models.ResponseMessage {
	var paquete models.Paquete
	err := service.db.Model(&models.Paquete{}).Where("id = ?", paqueteID).First(&paquete).Error
	if err != nil {
		return &models.ResponseMessage{IsSuccessfull: false, Message: err.Error()}
	}
	return &models.ResponseMessage{IsSuccessfull: true, Message: "Paquete encontrado exitosamente", Data: paquete}
}

func CalcularPrecio(paquete models.Paquete) float64 {
	return paquete.Peso + paquete.CuotaDestino
}

func (service *PaqueteService) CrearFactura(factura models.Factura) *models.ResponseMessage {
	tx := service.db.Begin()
	var total float64
	for _, paquete := range factura.Paquetes {
		total += CalcularPrecio(paquete)
	}
	factura.Total = total
	err := tx.Model(&models.Factura{}).Create(&factura).Error
	if err != nil {
		tx.Rollback()
		return &models.ResponseMessage{IsSuccessfull: false, Message: err.Error()}
	}
	tx.Commit()
	return &models.ResponseMessage{IsSuccessfull: true, Message: "Factura creada exitosamente", Data: factura}
}

func (service *PaqueteService) AsignarPaqueteARuta(paqueteID int, rutaID int) *models.ResponseMessage {
	tx := service.db.Begin()

	err := tx.Model(&models.Paquete{}).Where("id = ?", paqueteID).Update("estado_id", EN_RUTA).Error
	if err != nil {
		tx.Rollback()
		return &models.ResponseMessage{IsSuccessfull: false, Message: err.Error()}
	}

	err = tx.Model(&models.Ruta{}).Where("id = ?", rutaID).Update("capacidad", gorm.Expr("capacidad - ?", 1)).Error
	if err != nil {
		tx.Rollback()
		return &models.ResponseMessage{IsSuccessfull: false, Message: err.Error()}
	}

	paqueteRuta := models.PaqueteRuta{
		PaqueteID: paqueteID,
		RutaID:    rutaID,
	}
	err = tx.Model(&models.PaqueteRuta{}).Create(&paqueteRuta).Error
	if err != nil {
		tx.Rollback()
		return &models.ResponseMessage{IsSuccessfull: false, Message: err.Error()}
	}

	tx.Commit()

	return &models.ResponseMessage{IsSuccessfull: true, Message: "Paquete asignado a ruta exitosamente"}
}
