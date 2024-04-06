package service

import (
	"errors"
	"p1-backend/api/pkg/models"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	// Iniciar transacción
	tx := service.db.Begin()
	register.Password = HashPassword(register.Password)
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	err := tx.Omit(clause.Associations).Create(&register).Error
	if err != nil {
		var mysqlError *mysql.MySQLError
		if ok := errors.As(err, &mysqlError); ok {
			if mysqlError.Number == 1062 {
				tx.Rollback()
				return &models.ResponseMessage{IsSuccessfull: false, Message: "El correo ya esta registrado, intente con otro"}
			}
		}
		tx.Rollback()
		return &models.ResponseMessage{IsSuccessfull: false, Message: err.Error()}
	}
	tx.Commit()
	return &models.ResponseMessage{IsSuccessfull: true, Message: "Usuario registrado exitosamente"}
}

func (service *UsuarioService) GetUserByEmail(email string) (*models.Usuario, error) {
	user := new(models.Usuario)
	err := service.db.Where(`email = ?`, email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, err
}

func (service *UsuarioService) GetUserById(id int) (*models.Usuario, error) {
	user := new(models.Usuario)
	err := service.db.Where(`id = ?`, id).First(&user).Error
	return user, err
}

func (service *UsuarioService) GetAllUsers(page int, limit int) ([]models.Usuario, error) {
	var users []models.Usuario
	err := service.db.Limit(limit).Offset((page - 1) * limit).Preload("PuntoControl").Preload("Rol").Find(&users).Error
	return users, err
}
func (service *UsuarioService) GetAllUsersByRol(page, limit int, rol string) ([]models.Usuario, error) {
	var users []models.Usuario
	err := service.db.Limit(limit).Offset((page-1)*limit).Where(`rol = ?`, rol).Find(&users).Error
	return users, err
}

func (service *UsuarioService) ChangeStateUser(id int, activate bool) *models.ResponseMessage {
	user, err := service.GetUserById(id)
	if err != nil {
		return &models.ResponseMessage{IsSuccessfull: false, Message: err.Error()}
	}
	user.Activo = activate
	err = service.db.Save(&user).Error
	if err != nil {
		return &models.ResponseMessage{IsSuccessfull: false, Message: err.Error()}
	}
	return &models.ResponseMessage{IsSuccessfull: true, Message: "Cambio de estado exitoso"}
}

func (service *UsuarioService) UpdateUser(update *models.Usuario) *models.ResponseMessage {
	update.Password = HashPassword(update.Password)
	user := models.Usuario{
		Nombre:   update.Nombre,
		Apellido: update.Apellido,
		Email:    update.Email,
		Password: update.Password,
		Rol:      update.Rol,
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
