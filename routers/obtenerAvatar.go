package routers

import (
	"api_rest_en_go/bd"
	"io"
	"net/http"
	"os"
)

/*ObtenerAvatar envia el Avatar al HTTP*/
func ObtenerAvatar(wr http.ResponseWriter, r *http.Request) {

	usuarioId := r.URL.Query().Get("id")
	if len(usuarioId) < 1 {
		http.Error(wr, "El parámetro id es obligatorio", http.StatusBadRequest)
		return
	}

	perfil, err := bd.BuscarPerfil(usuarioId)
	if err != nil {
		http.Error(wr, "Usuario no encontrado", http.StatusNotFound)
		return
	}

	openFile, err := os.Open("uploads/avatars" + perfil.Avatar)
	if err != nil {
		http.Error(wr, "Avatar no encontrado", http.StatusNotFound)
		return
	}

	_, err = io.Copy(wr, openFile) //Enviar la imagen a ResponseWriter
	if err != nil {
		http.Error(wr, "Error al copiar la imagen", http.StatusInternalServerError)
	}

}
