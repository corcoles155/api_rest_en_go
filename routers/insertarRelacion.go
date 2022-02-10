package routers

import (
	"api_rest_en_go/bd"
	"api_rest_en_go/models"
	"net/http"
)

/*InsertarRelacion permite grabar la relacion entre dos usuarios en BBDD*/
func InsertarRelacion(wr http.ResponseWriter, r *http.Request) {

	usuarioAlQueQuieroSeguirId := r.URL.Query().Get("id")
	if len(usuarioAlQueQuieroSeguirId) < 1 {
		http.Error(wr, "El parámetro id es obligatorio", http.StatusBadRequest)
		return
	}

	var relacion models.Relacion
	relacion.UsuarioId = IDUsuario
	relacion.UsuarioRelacionId = usuarioAlQueQuieroSeguirId

	status, err := bd.InsertarRelacion(relacion)
	if err != nil {
		http.Error(wr, "Ocurrió un error al intentar relación en BBDD. Reintente de nuevo "+err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(wr, "No se ha logrado insertar la relacion en BBDD", http.StatusBadRequest)
		return
	}

	wr.WriteHeader(http.StatusCreated) //Devolvemos un 201
}
