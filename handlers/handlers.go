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

	router.HandleFunc("/registro", middlew.ComprobarBD(routers.Registro)).Methods("POST")                                  //Cuando llegue el endpoint POST /registro va a llamar a middlew.ComprobarBD, ComprobarBD devolverá la función Registro
	router.HandleFunc("/login", middlew.ComprobarBD(routers.Login)).Methods("POST")                                        //Cuando llegue el endpoint POST /login va a llamar a middlew.ComprobarBD, ComprobarBD devolverá la función Login
	router.HandleFunc("/verPerfil", middlew.ComprobarBD(middlew.ValidarJWT(routers.VerPerfil))).Methods("GET")             //Cuando llegue el endpoint GET /verperfil va a llamar a middlew.ComprobarBD, ComprobarBD devolverá la función ValidarJWT que a su vez devolverá la función VerPerfil
	router.HandleFunc("/modificarPerfil", middlew.ComprobarBD(middlew.ValidarJWT(routers.ModificarPerfil))).Methods("PUT") //Cuando llegue el endpoint PUT /modificarPerfil va a llamar a middlew.ComprobarBD, ComprobarBD devolverá la función ValidarJWT que a su vez devolverá la función ModificarPerfil
	router.HandleFunc("/tweet", middlew.ComprobarBD(middlew.ValidarJWT(routers.GrabarTweet))).Methods("POST")              //Cuando llegue el endpoint POST /tweet va a llamar a middlew.ComprobarBD, ComprobarBD devolverá la función ValidarJWT que a su vez devolverá la función GrabarTweet
	router.HandleFunc("/leerTweets", middlew.ComprobarBD(middlew.ValidarJWT(routers.LeerTweets))).Methods("GET")           //Cuando llegue el endpoint GET /leerTweets va a llamar a middlew.ComprobarBD, ComprobarBD devolverá la función ValidarJWT que a su vez devolverá la función LeerTweets

	PORT := os.Getenv("PORT") //Recuperamos la variable de entorno PORT
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)        //Permisos a todos
	log.Fatal(http.ListenAndServe(":"+PORT, handler)) //El servidor escucha en el puerto

}
