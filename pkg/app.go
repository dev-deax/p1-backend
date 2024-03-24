package app

import (
	"fmt"
	"log"
	"net/http"
	"p1-backend/api/pkg/config"
	"p1-backend/api/pkg/middleware"
	"p1-backend/api/pkg/models"
	"p1-backend/api/pkg/routes"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type App struct {
	Router      *mux.Router
	DB          *gorm.DB
	RateLimiter *middleware.RateLimiterStore
	Config      *config.Config
}

func ConfigAndRunApp(config *config.Config) {
	app := new(App)
	app.Initialize(config)
	app.Run(config.ServerHost)
}

func (app *App) Initialize(config *config.Config) {
	app.Config = config

	db, errConn := gorm.Open(mysql.Open(config.DatabaseURI()), &gorm.Config{})

	if errConn != nil {
		log.Fatal(errConn)
	}
	fmt.Println("Conexion con la base de datos exitosa!")
	app.DB = db

	errMigrate := app.DB.AutoMigrate(&models.Usuario{},
		&models.Cliente{},
		&models.Paquete{},
		&models.Ruta{},
		&models.Tarifa{},
		&models.Factura{},
		&models.ItemFactura{},
		&models.PuntoControl{},
		&models.PaquetesPuntosControl{})
	if errMigrate != nil {
		log.Fatal(errMigrate)
	}
	fmt.Println("Migracion de la base de datos exitosa!")
	// app.RateLimiter = middleware.NewRateLimiterStore(config.RateMinute)
	app.Router = mux.NewRouter()

	// app.UseMiddleware(handler.JSONContentTypeMiddleware)
	app.setRouters()
}

func (app *App) Run(host string) {
	fmt.Printf("Server started at http://localhost%s\n", host)
	log.Fatal(http.ListenAndServe(host, app.Router))
}

func (app *App) setRouters() {
	routes.InitUsuarioRoutes(app.Router, app.DB, app.authorizeRequest)

	routes := app.Router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pathTemplate, err := route.GetPathTemplate()
		if err == nil {
			fmt.Println("Ruta:", pathTemplate)
		}
		return nil
	})
	if routes != nil {
		fmt.Println("Error al obtener las rutas:", routes)
	}
}

func (app *App) authorizeRequest(next http.Handler, tokened bool) http.Handler {
	if tokened {
		return middleware.AppKeyAuthorization(middleware.AuthMiddleware(next), app.Config)
	} else {
		return middleware.AppKeyAuthorization(next, app.Config)
	}
}
