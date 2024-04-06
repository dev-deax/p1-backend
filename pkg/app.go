package app

import (
	"fmt"
	"log"
	"net/http"
	"p1-backend/api/pkg/config"
	"p1-backend/api/pkg/middleware"
	"p1-backend/api/pkg/models"
	"p1-backend/api/pkg/routes"
	"p1-backend/api/pkg/utils"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
	Config *config.Config
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

	errMigrate := app.DB.AutoMigrate(
		// &models.Usuario{},
		// &models.Cliente{},
		// &models.Paquete{},
		// &models.Destino{},
		// &models.Ruta{},
		// &models.Bodega{},
		// &models.PaquetesDestinos{},
		// &models.PuntoControl{},
		&models.Factura{},
		&models.Tarifas{},
		&models.PaquetesPuntosControl{})
	if errMigrate != nil {
		log.Fatal(errMigrate)
	}
	fmt.Println("Migracion de la base de datos exitosa!")
	// app.RateLimiter = middleware.NewRateLimiterStore(config.RateMinute)
	app.Router = mux.NewRouter()

	// app.UseMiddleware(handler.JSONContentTypeMiddleware)
	app.initRouters()
}

func (app *App) Run(host string) {
	// fmt.Printf("El servidor en: http://localhost%s\n", host)
	log.Fatal(http.ListenAndServe(host, app.Router))
}

func (app *App) initRouters() {
	routes.InitUsuarioRoutes(app.Router, app.DB, app.authorizeRequest)
	routes.InitDestinoRoutes(app.Router, app.DB, app.authorizeRequest)
	routes.InitRutaRoutes(app.Router, app.DB, app.authorizeRequest)
	routes.InitPuntoControlRoutes(app.Router, app.DB, app.authorizeRequest)
	routes.InitClienteRoutes(app.Router, app.DB, app.authorizeRequest)
	routes.InitPaqueteRoutes(app.Router, app.DB, app.authorizeRequest)
	bodegaUtil := utils.InitializeBodegaUtil(app.DB)
	go bodegaUtil.IniciarProcesoAutoBodega()
	errorRoutes := printRoutes(app.Router, app.Config.ServerHost)
	if errorRoutes != nil {
		fmt.Println("Error al obtener las rutas:", errorRoutes)
	}
}

func printRoutes(Router *mux.Router, serverURL string) error {
	errorRoutes := Router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pathTemplate, err := route.GetPathTemplate()
		if err == nil {
			fmt.Printf("http://localhost%s%s\n", serverURL, pathTemplate)
		}
		return nil
	})
	return errorRoutes
}

func (app *App) authorizeRequest(next http.Handler, tokened bool) http.Handler {
	if tokened {
		return middleware.AppKeyAuthorization(middleware.AuthMiddleware(next), app.Config)
	} else {
		return middleware.AppKeyAuthorization(next, app.Config)
	}
}
