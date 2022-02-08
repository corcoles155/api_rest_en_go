package main

import (
	"log"

	"api_rest_en_go/bd"
	"api_rest_en_go/handlers"
)

func main() {
	if bd.ComprobarConexion() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()
}
