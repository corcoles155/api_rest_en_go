package routers

import (
	"api_rest_en_go/bd"
	"api_rest_en_go/models"
	"net/http"
)

/*EliminarRelacion */
func EliminarRelacion(wr http.ResponseWriter, r *http.Request) {
	usuarioAlQueQuieroDejarDeSeguirId := r.URL.Query().Get("id")
	if len(usuarioAlQueQuieroDejarDeSeguirId) < 1 {
		http.Error(wr, "Debe enviar el parámetro id", http.StatusBadRequest)
		return
	}

	var relacion models.Relacion
	relacion.UsuarioId = IDUsuario
	relacion.UsuarioRelacionId = usuarioAlQueQuieroDejarDeSeguirId

	status, err := bd.BorrarRelacion(relacion)
	if err != nil {
		http.Error(wr, "Ocurrió un error al eliminar la relación "+err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(wr, "No se ha logrado eliminar la relación", http.StatusBadRequest)
		return
	}

	wr.Header().Set("context-type", "application/json") //Indicamos que el resultado será un JSON
	wr.WriteHeader(http.StatusCreated)
}
