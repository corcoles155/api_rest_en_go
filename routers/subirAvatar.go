package routers

import (
	"api_rest_en_go/bd"
	"api_rest_en_go/models"
	"io"
	"net/http"
	"os"
	"strings"
)

/*SubirAvatar guardar avatar en BBDD*/
func SubirAvatar(wr http.ResponseWriter, r *http.Request) {
	fichero, handler, _ := r.FormFile("avatar")
	var extension = strings.Split(handler.Filename, ".")[1]
	var nombreFichero string = "uploads/avatars/" + IDUsuario + "." + extension

	f, err := os.OpenFile(nombreFichero, os.O_WRONLY|os.O_CREATE, 0666) //Permisos de lectura, escritura y ejecución
	if err != nil {
		http.Error(wr, "Error al subir la imagen "+err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = io.Copy(f, fichero) //Copia el fichero en uploads/avatars
	if err != nil {
		http.Error(wr, "Error al copiar la imagen "+err.Error(), http.StatusInternalServerError)
		return
	}

	var usuario models.Usuario

	usuario.Avatar = IDUsuario + "." + extension
	status, err := bd.ModificarRegistro(usuario, IDUsuario)

	if err != nil || !status {
		http.Error(wr, "Error al guardar el avatar en BBDD "+err.Error(), http.StatusInternalServerError)
		return
	}

	wr.Header().Set("context-type", "application/json") //Indicamos que el resultado será un JSON
	wr.WriteHeader(http.StatusCreated)                  //Devolvemos un status 201
}
