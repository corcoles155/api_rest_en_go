package routers

import (
	"api_rest_en_go/bd"
	"api_rest_en_go/models"
	"encoding/json"
	"net/http"
)

/* Registro función para dar de alta en BBDD el usuario*/
func Registro(w http.ResponseWriter, r *http.Request) {
	var usuario models.Usuario
	err := json.NewDecoder(r.Body).Decode(&usuario) //El objeto Body es un Stream, sólo se puede leer una vez

	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return
	}

	if len(usuario.Email) == 0 {
		http.Error(w, "El campo email es obligatorio", 400)
		return
	}
	if len(usuario.Password) < 6 {
		http.Error(w, "La contraseña debe de tener al menos 6 caracteres", 400)
		return
	}

	_, encontrado, _ := bd.ExisteUsuario(usuario.Email)
	if encontrado == true {
		http.Error(w, "Ya existe un usuario registrado con ese email", 400)
		return
	}

	_, status, err := bd.InsertarRegistro(usuario)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar registrar el usuario"+err.Error(), 500)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado registrar el usuario", 500)
		return
	}

	w.WriteHeader(http.StatusCreated) //Devolvemos un 200
}
