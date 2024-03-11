package controllers

import (
	"fmt"
	"net/http"
)

func obtenerUsuarios(w http.ResponseWriter, r *http.Request) {
	// response := ExampleResponse{Message: "Hola, este es un mensaje JSON"}

	// Convertir la estructura a JSON

	// if err != nil {
	// 	http.Error(w, "Error al serializar la respuesta JSON", http.StatusInternalServerError)
	// 	return
	// }

	// Configurar encabezados de respuesta
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("dsfdsfdsf")

	// Escribir la respuesta JSON
	// w.Write({'sss':"ssasasd"})
}

// Función para obtener un usuario por ID
func obtenerUsuario(w http.ResponseWriter, r *http.Request) {
	// ...

}

// Función para crear un nuevo usuario
func crearUsuario(w http.ResponseWriter, r *http.Request) {
	// ...

}

// Función para actualizar un usuario por ID
func actualizarUsuario(w http.ResponseWriter, r *http.Request) {
	// ...

}

// Función para eliminar un usuario por ID
func eliminarUsuario(w http.ResponseWriter, r *http.Request) {
	// ...

}
