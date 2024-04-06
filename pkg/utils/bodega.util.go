package utils

import (
	"fmt"
	service "p1-backend/api/pkg/services"
	"time"

	"gorm.io/gorm"
)

type BodegaUtil struct {
	paqueteService      *service.PaqueteService
	puntoControlService *service.PuntoControlService
	rutaService         *service.RutaService
}

func InitializeBodegaUtil(db *gorm.DB) *BodegaUtil {
	paqueteService := service.InitializePaqueteService(db)
	puntoControlService := service.InitializePuntoControlService(db)
	rutaService := service.InitializeRutaService(db)
	bodegaUtil := BodegaUtil{rutaService: rutaService, paqueteService: paqueteService, puntoControlService: puntoControlService}
	return &bodegaUtil
}

func (bodegaUtil *BodegaUtil) IniciarProcesoAutoBodega() {
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()
	for range ticker.C {
		bodegaUtil.assignPaquetes()
	}
}

func (bodegaUtil *BodegaUtil) assignPaquetes() {
	fmt.Println("ejecutando ")
	paquetes := bodegaUtil.paqueteService.GetPaquetesBodega()
	if len(*paquetes) == 0 {
		fmt.Println("Sin paquetes que asignar")
	}
	for _, paquete := range *paquetes {
		var destino = paquete.Destino
		var rutasDestino = bodegaUtil.rutaService.GetRutasByDestinoId(destino.ID)
		for _, ruta := range rutasDestino {
			cantPaquetesRuta, _ := bodegaUtil.rutaService.GetCantidadPaquetesRuta(ruta.ID)
			fmt.Println("=====================")
			fmt.Println("rutasDestino")
			fmt.Println(cantPaquetesRuta)
			fmt.Println(ruta.Capacidad)
			fmt.Println("=====================")
			if cantPaquetesRuta < int64(ruta.Capacidad) {
				fmt.Println("=====================")
				fmt.Println("Asiganado paquete")
				fmt.Println(paquete.ID)
				fmt.Println(ruta.Nombre)
				fmt.Println("=====================")
				bodegaUtil.paqueteService.AsignarPaqueteARuta(paquete, ruta)
			}
		}
		fmt.Println(paquete.ID)
		fmt.Println(destino.ID)

	}
}
