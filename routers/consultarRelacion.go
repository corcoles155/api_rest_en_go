package routers

import (
	"api_rest_en_go/bd"
	"api_rest_en_go/models"
	"encoding/json"
	"net/http"
)

/*ConsultarRelacion permite consultar la relación entre dos usuarios*/
func ConsultarRelacion(wr http.ResponseWriter, r *http.Request) {
	usuarioAlQueQuieroSeguirId := r.URL.Query().Get("id") //Extraemos de la URL el parámetro ID
	if len(usuarioAlQueQuieroSeguirId) < 1 {              //Si no encontró en la URL el parámetro id
		http.Error(wr, "Debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}

	var relacion models.Relacion
	relacion.UsuarioId = IDUsuario
	relacion.UsuarioRelacionId = usuarioAlQueQuieroSeguirId

	var respuestaConsultaRelacion models.RespuestaConsultaRelacion

	status, err := bd.ConsultarRelacion(relacion)
	if err != nil || !status {
		respuestaConsultaRelacion.Status = false
	} else {
		respuestaConsultaRelacion.Status = true
	}

	wr.Header().Set("context-type", "application/json")
	wr.WriteHeader(http.StatusCreated)
	json.NewEncoder(wr).Encode(respuestaConsultaRelacion)
}
