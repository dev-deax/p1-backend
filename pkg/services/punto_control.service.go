package service

import (
	"p1-backend/api/pkg/models"
	"time"

	"gorm.io/gorm"
)

type PuntoControlService struct {
	db *gorm.DB
}

type RoleUsuario int

const (
	ADMIN         RoleUsuario = 1
	OPERARIO      RoleUsuario = 2
	RECEPCIONISTA RoleUsuario = 3
)

func InitializePuntoControlService(db *gorm.DB) *PuntoControlService {
	return &PuntoControlService{db: db}
}
func (service *PuntoControlService) Migrate() error {
	return service.db.AutoMigrate(&models.Paquete{})
}
func (service *PuntoControlService) Create(create *models.PuntoControl) *models.ResponseMessage {
	var usuario models.Usuario
	errorUsuario := service.db.Where(`id = ?`, create.UsuarioID).First(&usuario).Error
	if errorUsuario != nil {
		return &models.ResponseMessage{IsSuccessfull: false, Message: errorUsuario.Error()}
	}
	if usuario.RolID != int(OPERARIO) {
		return &models.ResponseMessage{IsSuccessfull: false, Message: "No puede asignar a este usuario a este punto de control, este usuario no es operario"}
	}
	err := service.db.Save(&create).Error
	if err != nil {
		return &models.ResponseMessage{IsSuccessfull: false, Message: err.Error()}
	}
	return &models.ResponseMessage{IsSuccessfull: true, Message: "Registrado exitosamente"}
}
func (service *PuntoControlService) Update(update *models.PuntoControl) *models.ResponseMessage {
	err := service.db.Save(&update).Error
	if err != nil {
		return &models.ResponseMessage{IsSuccessfull: false, Message: err.Error()}
	}
	return &models.ResponseMessage{IsSuccessfull: true, Message: "Actualizado exitosamente"}
}
func (service *PuntoControlService) GetById(id int) *models.ResponseMessage {
	var data models.PuntoControl
	err := service.db.Preload("Paquete").Preload("Ruta").First(&data, id).Error
	if err != nil {
		return &models.ResponseMessage{IsSuccessfull: false, Message: err.Error()}
	}
	return &models.ResponseMessage{IsSuccessfull: true, Message: "Punto de control obtenido exitosamente", Data: data}
}
func (service *PuntoControlService) GetAll(page int, limit int) *models.ResponseMessage {
	var data []models.PuntoControl
	offset := (page - 1) * limit
	err := service.db.Limit(limit).Offset(offset).Preload("Paquetes").Preload("Ruta").Preload("Ruta.Destino").Find(&data).Error
	if err != nil {
		return &models.ResponseMessage{IsSuccessfull: false, Message: err.Error()}
	}
	return &models.ResponseMessage{IsSuccessfull: true, Message: "Puntos de control obtenidos exitosamente", Data: data}
}
func (service *PuntoControlService) GetPuntoControlById(id int) (*models.PuntoControl, error) {
	model := new(models.PuntoControl)
	err := service.db.Where(`id = ?`, id).First(&model).Error
	return model, err
}
func (service *PuntoControlService) ChangeStatePuntoControl(id int, activate bool) *models.ResponseMessage {
	tienePaquetesCola, _, err := service.TienePaquetesCola(id)
	if err != nil {
		return &models.ResponseMessage{IsSuccessfull: false, Message: err.Error()}
	}
	if tienePaquetesCola {
		return &models.ResponseMessage{IsSuccessfull: false, Message: "No se puede desactivar porque tiene paquetes en cola"}
	}
	model, err := service.GetPuntoControlById(id)
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
func (service *PuntoControlService) TienePaquetesCola(puntoControlID int) (bool, int64, error) {
	var count int64
	err := service.db.Model(&models.PaquetesPuntosControl{}).Where("punto_control_id = ? AND (tiempo_punto_control IS NULL OR tiempo_punto_control = 0)", puntoControlID).Count(&count).Error
	if err != nil {
		return false, 0, err
	}
	return count > 0, count, nil
}
func (service *PuntoControlService) ProcesarPaquete(paqueteID int, puntoControlID int, salida bool) *models.ResponseMessage {
	fecha := time.Now()
	columnaFecha := "fecha_llegada"
	if salida {
		columnaFecha = "fecha_salida"
	}
	err := service.db.Model(&models.PaquetesPuntosControl{}).Where("paquete_id = ? AND ", paqueteID).Where("punto_control_id = ? AND ", puntoControlID).Update(columnaFecha, fecha).Error
	if err != nil {
		return &models.ResponseMessage{IsSuccessfull: false, Message: err.Error()}
	}
	return &models.ResponseMessage{IsSuccessfull: true, Message: "Cambio de estado exitoso"}
}
func (service *PuntoControlService) GetPuntoControlByUsuarioID(UsuarioID int) (*models.PuntoControl, error) {
	var puntoControl models.PuntoControl
	err := service.db.Where("usuario = ?", UsuarioID).First(&puntoControl).Error
	return &puntoControl, err
}
func (service *PuntoControlService) GetPaquetesPuntosControlByUsuario(UsuarioID int) *models.ResponseMessage {
	puntoControl, err := service.GetPuntoControlByUsuarioID(UsuarioID)
	if err != nil {
		return &models.ResponseMessage{IsSuccessfull: false, Message: err.Error()}
	}
	var paquetesPuntosControl []models.PaquetesPuntosControl
	err = service.db.Where("punto_control_id = ?", puntoControl.ID).Find(&paquetesPuntosControl).Error
	if err != nil {
		return &models.ResponseMessage{IsSuccessfull: false, Message: err.Error()}
	}
	return &models.ResponseMessage{IsSuccessfull: true, Message: "PaquetesPuntosControl obtenidos exitosamente", Data: paquetesPuntosControl}
}
func (service *PuntoControlService) GetPuntoControlsByUsuarioID(UsuarioID int) (*[]models.PuntoControl, error) {
	var puntoControls []models.PuntoControl
	err := service.db.Where("usuario_id = ?", UsuarioID).Find(&puntoControls).Error
	return &puntoControls, err
}
func (service *PuntoControlService) GetPaquetesProcesarPuntosControl(UsuarioID int, puntoControlID int) *models.ResponseMessage {

	var paquetes []models.Paquete
	query := service.db.
		Joins("INNER JOIN paquetes_rutas ON paquetes.id = paquetes_rutas.paquete_id").
		Joins("INNER JOIN punto_controls ON paquetes_rutas.ruta_id = punto_controls.id").
		Joins("INNER JOIN usuarios ON punto_controls.usuario_id = usuarios.id").
		Where("paquetes.estado_id = ?", 2).
		Where("usuarios.id = ?", UsuarioID).
		Where("paquetes.deleted_at IS NULL")

	if puntoControlID != 0 {
		query = query.Where("punto_controls.id = ?", puntoControlID)
	}
	err := query.
		Find(&paquetes).
		Error

	if err != nil {
		return &models.ResponseMessage{IsSuccessfull: false, Message: err.Error()}
	}

	// var paquetesPuntosControl []models.PaquetesPuntosControl
	// err = service.db.
	// 	Where("punto_control_id = ?", puntoControl.ID).
	// 	Find(&paquetesPuntosControl).Error
	// if err != nil {
	// 	return &models.ResponseMessage{IsSuccessfull: false, Message: err.Error()}
	// }
	return &models.ResponseMessage{IsSuccessfull: true, Message: "PaquetesPuntosControl obtenidos exitosamente", Data: paquetes}
}
func (service *PuntoControlService) GetCostoPaquetesPuntosControlByPuntoControl(puntoControlID int) *models.ResponseMessage {
	var PuntoControl models.PuntoControl
	err := service.db.Model(&models.PuntoControl{}).Where("id = ?", puntoControlID).Preload("Paquete").First(&PuntoControl).Error
	if err != nil {
		return &models.ResponseMessage{IsSuccessfull: false, Message: err.Error()}
	}
	timpoTotal := 0.0

	paquetes := PuntoControl.Paquetes

	for _, paquete := range paquetes {

		tiempo, errPuntoControl := service.GetTiempoTotalEnRuta(paquete)
		if errPuntoControl != nil {
			timpoTotal += 0.0
		}
		timpoTotal += tiempo
	}

	var Tarifas models.Tarifas
	errTarifas := service.db.Where("tipo like('%?%')", "Tarifa operacion").First(&Tarifas).Error
	if errTarifas != nil {
		return &models.ResponseMessage{IsSuccessfull: false, Message: errTarifas.Error()}
	}
	costo := float64(timpoTotal) * float64(Tarifas.Valor)

	if PuntoControl.TarifaOperacion > 0 {
		costo = float64(timpoTotal) * float64(PuntoControl.TarifaOperacion)
	}
	return &models.ResponseMessage{IsSuccessfull: true, Message: "Costo obtenido exitosamente", Data: costo}
}
func (service *PuntoControlService) GetCostoPaquetesPuntosControlByPaquete(paqueteID int) *models.ResponseMessage {
	var paquete models.Paquete
	err := service.db.Model(&models.Paquete{}).Where("id = ?", paqueteID).First(&paquete).Error
	if err != nil {
		return &models.ResponseMessage{IsSuccessfull: false, Message: err.Error()}
	}
	timpoTotal, errPuntoControl := service.GetTiempoTotalEnRuta(paquete)

	if errPuntoControl != nil {
		return &models.ResponseMessage{IsSuccessfull: false, Message: errPuntoControl.Error()}
	}

	var Tarifas models.Tarifas
	errTarifas := service.db.Where("tipo like('%?%')", "Tarifa operacion").First(&Tarifas).Error
	if errTarifas != nil {
		return &models.ResponseMessage{IsSuccessfull: false, Message: errTarifas.Error()}
	}
	costo := float64(timpoTotal) * float64(Tarifas.Valor)

	if paquete.TarifaOperacion > 0 {
		costo = float64(timpoTotal) * float64(paquete.TarifaOperacion)
	}
	return &models.ResponseMessage{IsSuccessfull: true, Message: "Costo obtenido exitosamente", Data: costo}
}
func (service *PuntoControlService) GetTiempoTotalEnRuta(paquete models.Paquete) (float64, error) {

	var paquetePuntosControls []models.PaquetesPuntosControl
	ruta, errors := service.GetRutaActualPaquete(int(paquete.ID))
	if errors != nil {
		return 0, errors
	}
	err := service.db.Model(&models.PaquetesPuntosControl{}).Where("ruta_id = ?", ruta.ID).Find(&paquetePuntosControls).Error
	if err != nil {
		return 0, err
	}

	var timpoTotal float64
	for _, puntoControl := range paquetePuntosControls {
		timpoTotal += CalcularTiempoEnHoras(puntoControl.FechaLlegada, puntoControl.FechaSalida)
	}

	return timpoTotal, nil
}
func (service *PuntoControlService) GetRutaActualPaquete(paqueteID int) (*models.Ruta, error) {
	var paqueteRuta models.PaquetesRutas
	err := service.db.Model(&models.PaquetesRutas{}).Where("paquete_id = ?", paqueteID).Order("created_at desc").First(&paqueteRuta).Error
	if err != nil {
		return nil, err
	}

	var ruta *models.Ruta
	err = service.db.Model(&models.Ruta{}).Where("id = ?", paqueteRuta.RutaID).First(&ruta).Error
	if err != nil {
		return nil, err
	}

	return ruta, nil
}
func CalcularTiempoEnHoras(fechaLlegada, fechaSalida *time.Time) float64 {
	if fechaLlegada == nil || fechaSalida == nil {
		return 0
	}

	diferencia := fechaSalida.Sub(*fechaLlegada)
	return diferencia.Hours()
}
