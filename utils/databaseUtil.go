package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// Init se conecta a la BD y crea la tabla "contacts"
func Init() {
	username := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASS")
	dbName := os.Getenv("MYSQL_DB")
	dbHost := os.Getenv("MYSQL_HOST")
	dbURI := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, dbHost, dbName) // Formatear un string con parámetros
	log.Printf(dbURI)                                                                                               // Imprimir la URI de conexión para debug
	config := &gorm.Config{}

	conn, err := gorm.Open(mysql.Open(dbURI), config)

	if err != nil {
		log.Fatal(err)
	}
	db = conn
	// db.Debug().AutoMigrate(&Usuario{})  
}

// DB regresa el objeto de base de datos
func DB() *gorm.DB {
	return db
}
