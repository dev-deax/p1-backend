package service

import (
	"fmt"
	"p1-backend/api/pkg/models"
	"strconv"

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
	return paquete.Peso*paquete.PecioLb + paquete.CuotaDestino
}

func (service *PaqueteService) CrearFactura(factura models.Factura) *models.ResponseMessage {
	tx := service.db.Begin()
	var total float64
	for _, paquete := range factura.Paquetes {
		total += CalcularPrecio(paquete)
	}
	fmt.Printf("=====" + strconv.FormatFloat(total, 'f', 2, 64))
	factura.Total = total

	err := tx.Create(&factura).Error
	if err != nil {
		tx.Rollback()
		return &models.ResponseMessage{IsSuccessfull: false, Message: err.Error()}
	}
	tx.Commit()
	return &models.ResponseMessage{IsSuccessfull: true, Message: "Factura creada exitosamente", Data: factura}
}

func (service *PaqueteService) GetFacturaAll(page, limit int) *models.ResponseMessage {
	var facturas []models.Factura
	err := service.db.Model(&models.Factura{}).Offset((page - 1) * limit).Limit(limit).Find(&facturas).Error
	if err != nil {
		return &models.ResponseMessage{IsSuccessfull: false, Message: err.Error()}
	}
	return &models.ResponseMessage{IsSuccessfull: true, Message: "Facturas encontradas exitosamente", Data: facturas}
}
func (service *PaqueteService) GetFacturaByID(facturaID int) *models.ResponseMessage {
	var factura models.Factura
	err := service.db.Model(&models.Factura{}).Where("id = ?", facturaID).Preload("Paquetes").Preload("Cliente").First(&factura).Error
	if err != nil {
		return &models.ResponseMessage{IsSuccessfull: false, Message: err.Error()}
	}
	return &models.ResponseMessage{IsSuccessfull: true, Message: "Factura encontrada exitosamente", Data: factura}
}

func (service *PaqueteService) AsignarPaqueteARuta(paquete models.Paquete, ruta models.Ruta) *models.ResponseMessage {
	tx := service.db.Begin()

	err := tx.Model(&models.Paquete{}).Where("id = ?", paquete.ID).Update("estado_id", EN_RUTA).Error
	if err != nil {
		tx.Rollback()
		return &models.ResponseMessage{IsSuccessfull: false, Message: err.Error()}
	}

	ruta.Paquetes = append(ruta.Paquetes, paquete)

	err = tx.Save(&ruta).Error
	if err != nil {
		tx.Rollback()
		return &models.ResponseMessage{IsSuccessfull: false, Message: err.Error()}
	}

	tx.Commit()
	fmt.Println("todo correcto")
	return &models.ResponseMessage{IsSuccessfull: true, Message: "Paquete asignado a ruta exitosamente"}
}

func (service *PaqueteService) GetPaquetesBodega() *[]models.Paquete {
	var paquetes []models.Paquete
	service.db.Model(&models.Paquete{}).Where("estado_id = ?", EN_BODEGA).Preload("Destino").Find(&paquetes)
	return &paquetes
}
