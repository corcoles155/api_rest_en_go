package handlers

import (
	"log"
	"net/http"
	"os"

	"api_rest_en_go/middlew"
	"api_rest_en_go/routers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/* Manejadores seteo mi puerto, el handler y pongo a escuchar al servidor*/
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ComprobarBD(routers.Registro)).Methods("POST") //Cuando llegue el endpoint POST /registro va a llamar a middlew.ChequeoBD

	PORT := os.Getenv("PORT") //Recuperamos la variable de entorno PORT
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)        //Permisos a todos
	log.Fatal(http.ListenAndServe(":"+PORT, handler)) //El servidor escucha en el puerto

}