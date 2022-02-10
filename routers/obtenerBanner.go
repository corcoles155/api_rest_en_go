package routers

import (
	"api_rest_en_go/bd"
	"io"
	"net/http"
	"os"
)

/*ObtenerBanner envia el Avatar al HTTP*/
func ObtenerBanner(wr http.ResponseWriter, r *http.Request) {

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

	openFile, err := os.Open("uploads/banners" + perfil.Banner)
	if err != nil {
		http.Error(wr, "Banner no encontrado", http.StatusNotFound)
		return
	}

	_, err = io.Copy(wr, openFile) //Enviar la imagen a ResponseWriter
	if err != nil {
		http.Error(wr, "Error al copiar la imagen", http.StatusInternalServerError)
	}

}
