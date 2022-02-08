package middlew

import (
	"api_rest_en_go/bd"
	"net/http"
)

/* ComprobarBD me permite conocer el estado de la BBDD */
func ComprobarBD(next http.HandlerFunc) http.HandlerFunc { //Los middleware reciben una función y devuelven una función
	return func(rw http.ResponseWriter, r *http.Request) { //devolvemos una función anónima (funciones sin nombre)
		if bd.ComprobarConexion() == 0 {
			http.Error(rw, "Conexión perdida con la Base de datos", 500) //Devolvemos un error 500
			return
		}
		next.ServeHTTP(rw, r)
	}
}
