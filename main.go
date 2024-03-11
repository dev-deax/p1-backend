package main

import (
	"log"
	"net/http"
	"os"
	database "p1-backend/api/utils"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// ... (se omite el código para las demás entidades)

// Función principal
// func main() {
// 	// Conectar a la base de datos
// 	// conectarDB()

// 	// Crear router Gorilla MUX
// 	router := mux.NewRouter().PathPrefix("/api").Subrouter()

// 	// Definir rutas para usuarios
// 	router.HandleFunc("/usuarios", obtenerUsuarios).Methods("GET")
// 	router.HandleFunc("/usuarios/{id}", obtenerUsuario).Methods("GET")
// 	router.HandleFunc("/usuarios", crearUsuario).Methods("POST")
// 	router.HandleFunc("/usuarios/{id}", actualizarUsuario).Methods("PUT")
// 	router.HandleFunc("/usuarios/{id}", eliminarUsuario).Methods("DELETE")

// 	// ... (se omite la definición de rutas para las demás entidades)

// 	// Iniciar servidor HTTP
// 	log.Println("Iniciando servidor en http://localhost:8080")
// 	log.Fatal(http.ListenAndServe(":8080", router))

//		// Cerrar conexión a la base de datos
//		// cerrarDB()
//	}
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err) // Imprimir en consola y terminar el programa.
	}
	database.Init()

	router := mux.NewRouter().PathPrefix("/api").Subrouter()
	log.Println("Iniciando servidor en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":"+os.Getenv("HTTP_PORT"), router))
}
