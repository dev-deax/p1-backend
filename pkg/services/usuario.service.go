package service

import (
	"errors"
	"p1-backend/api/pkg/models"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type UsuarioService struct {
	db *gorm.DB
}

func InitializeUsuarioService(db *gorm.DB) *UsuarioService {
	return &UsuarioService{db: db}
}

func (service *UsuarioService) Migrate() error {
	return service.db.AutoMigrate(&models.Usuario{})
}

func (service *UsuarioService) RegisterUser(register *models.Usuario) *models.ResponseMessage {
	register.Password = HashPassword(register.Password)
	user := models.Usuario{
		Nombre:   register.Nombre,
		Apellido: register.Apellido,
		Email:    register.Email,
		Password: register.Password,
		Rol:      register.Rol,
	}
	err := service.db.Save(&user).Error
	if err != nil {
		var mysqlError *mysql.MySQLError
		if ok := errors.As(err, &mysqlError); ok {
			if mysqlError.Number == 1062 {
				return &models.ResponseMessage{IsSuccessfull: false, Message: "El correo ya esta registrado, intente con otro"}
			}
		}
		return &models.ResponseMessage{IsSuccessfull: false, Message: err.Error()}
	}

	return &models.ResponseMessage{IsSuccessfull: true, Message: "Usuario registrado exitosamente"}
}

func (service *UsuarioService) GetUserByEmail(email string) (*models.Usuario, error) {
	user := new(models.Usuario)
	err := service.db.Where(`email = ?`, email).First(&user).Error
	return user, err
}
