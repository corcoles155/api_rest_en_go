package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"api_rest_en_go/bd"
	"api_rest_en_go/jwt"
	"api_rest_en_go/models"
)

/* Login realiza el login*/
func Login(wr http.ResponseWriter, r *http.Request) {
	wr.Header().Add("content-type", "application/json") //Indicamos que vamos a devolver un JSON como respuesta de nuestro endpoint

	var usuario models.Usuario

	err := json.NewDecoder(r.Body).Decode(&usuario) //Leo el stream Body y escribo el resultado en la variable usuario, si hay error escribe en la variable err
	if err != nil {
		http.Error(wr, "Usuario y/o contraseña incorrecta"+err.Error(), http.StatusBadRequest)
		return
	}

	if len(usuario.Email) == 0 {
		http.Error(wr, "El campo email es obligatorio", http.StatusBadRequest)
		return
	}

	documento, existe := bd.ComprobarLogin(usuario.Email, usuario.Password)
	if !existe {
		http.Error(wr, "Usuario y/o contraseña incorrecta", http.StatusBadRequest)
		return
	}

	jwtKey, err := jwt.GenerarJWT(documento)
	if err != nil {
		http.Error(wr, "Ocurrió un error al intentar generar el token"+err.Error(), http.StatusInternalServerError)
		return
	}

	resp := models.RespuestaLogin{
		Token: jwtKey,
	}

	wr.Header().Set("Content-Type", "application/json")
	wr.WriteHeader(http.StatusCreated) //Devolvemos un status 200
	json.NewEncoder(wr).Encode(resp)

	expierationTime := time.Now().Add(24 * time.Hour) //Creamos una fecha de expiración, a la fecha actual le agrega 24h
	http.SetCookie(wr, &http.Cookie{                  //Seteamos el valor del token en una cookie
		Name:    "token",
		Value:   jwtKey,
		Expires: expierationTime,
	})
}
