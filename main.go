package main

import (
	"log"
	app "p1-backend/api/pkg"
	"p1-backend/api/pkg/config"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err) // Imprimir en consola y terminar el proceso.
	}

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	config := config.NewConfig()
	app.ConfigAndRunApp(config)
}
