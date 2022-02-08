package routers

import (
	"api_rest_en_go/bd"
	"api_rest_en_go/models"
	"encoding/json"
	"net/http"
)

/* Registro función para dar de alta en BBDD el usuario*/
func Registro(wr http.ResponseWriter, r *http.Request) {
	var usuario models.Usuario
	err := json.NewDecoder(r.Body).Decode(&usuario) //El objeto Body es un Stream, sólo se puede leer una vez

	if err != nil {
		http.Error(wr, "Error en los datos recibidos"+err.Error(), http.StatusBadRequest)
		return
	}

	if len(usuario.Email) == 0 {
		http.Error(wr, "El campo email es obligatorio", http.StatusBadRequest)
		return
	}
	if len(usuario.Password) < 6 {
		http.Error(wr, "La contraseña debe de tener al menos 6 caracteres", http.StatusBadRequest)
		return
	}

	_, encontrado, _ := bd.ExisteUsuario(usuario.Email)
	if encontrado {
		http.Error(wr, "Ya existe un usuario registrado con ese email", http.StatusBadRequest)
		return
	}

	_, status, err := bd.InsertarRegistro(usuario)
	if err != nil {
		http.Error(wr, "Ocurrió un error al intentar registrar el usuario"+err.Error(), http.StatusInternalServerError)
		return
	}

	if !status {
		http.Error(wr, "No se ha logrado registrar el usuario", http.StatusInternalServerError)
		return
	}

	wr.WriteHeader(http.StatusCreated) //Devolvemos un 200
}
