package middlew

import (
	"api_rest_en_go/routers"
	"net/http"
)

/*ValidarJWT permite validar el JWT que nos viene en la petición*/
func ValidarJWT(next http.HandlerFunc) http.HandlerFunc { //Los middleware reciben una función y devuelven una función
	return func(rw http.ResponseWriter, r *http.Request) { //devolvemos una función anónima (funciones sin nombre)
		_, _, _, err := routers.ProcesoToken(r.Header.Get("Authorization")) //La variable Authorization es la variable que recibe el token
		if err != nil {
			http.Error(rw, "Error en el token "+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(rw, r)
	}
}
