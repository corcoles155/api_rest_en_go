package routers

import (
	"api_rest_en_go/bd"
	"encoding/json"
	"net/http"
)

/*VerPerfil permite extraer del body los valores del perfil*/
func VerPerfil(wr http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id") //Extraemos de la URL el parámetro ID
	if len(ID) < 1 {              //Si no encontró en la URL el parámetro id
		http.Error(wr, "Debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}

	perfil, err := bd.BuscarPerfil(ID)
	if err != nil {
		http.Error(wr, "Ocurrió un error al intentar buscar el registro "+err.Error(), http.StatusBadRequest)
		return
	}

	wr.Header().Set("context-type", "application/json")
	wr.WriteHeader(http.StatusCreated)
	json.NewEncoder(wr).Encode(perfil)
}
