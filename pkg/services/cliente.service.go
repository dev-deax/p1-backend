package service

import (
	"errors"
	"p1-backend/api/pkg/models"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type ClienteService struct {
	db *gorm.DB
}

func InitializeClienteService(db *gorm.DB) *ClienteService {
	return &ClienteService{db: db}
}

func (service *ClienteService) Migrate() error {
	return service.db.AutoMigrate(&models.Cliente{})
}

func (service *ClienteService) Register(register *models.Cliente) *models.ResponseMessage {
	cliente := models.Cliente{
		NIT:       register.NIT,
		Nombre:    register.Nombre,
		Apellido:  register.Apellido,
		Dirección: register.Dirección,
		Teléfono:  register.Teléfono,
	}
	err := service.db.Save(&cliente).Error
	if err != nil {
		var mysqlError *mysql.MySQLError
		if ok := errors.As(err, &mysqlError); ok {
			if mysqlError.Number == 1062 {
				return &models.ResponseMessage{IsSuccessfull: false, Message: "El NIT ya esta registrado, intente con otro"}
			}
		}
		return &models.ResponseMessage{IsSuccessfull: false, Message: err.Error()}
	}

	return &models.ResponseMessage{IsSuccessfull: true, Message: "Usuario registrado exitosamente"}
}

func (service *ClienteService) GetAll(page int, limit int) *models.ResponseMessage {
	var clientes []models.Cliente
	offset := (page - 1) * limit
	err := service.db.Limit(limit).Offset(offset).Find(&clientes).Error
	if err != nil {
		return &models.ResponseMessage{IsSuccessfull: false, Message: err.Error()}
	}
	return &models.ResponseMessage{IsSuccessfull: true, Message: "Clientes obtenidas exitosamente", Data: clientes}
}

func (service *ClienteService) GetById(id int) *models.ResponseMessage {
	var cliente models.Cliente
	err := service.db.First(&cliente, id).Error
	if err != nil {
		return &models.ResponseMessage{IsSuccessfull: false, Message: err.Error()}
	}
	return &models.ResponseMessage{IsSuccessfull: true, Message: "Cliente obtenido exitosamente", Data: cliente}
}
