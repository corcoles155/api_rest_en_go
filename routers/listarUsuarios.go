package routers

import (
	"api_rest_en_go/bd"
	"encoding/json"
	"net/http"
	"strconv"
)

/*ListarUsuarios permite listar los usuarios*/
func ListarUsuarios(wr http.ResponseWriter, r *http.Request) {

	tipoUsuario := r.URL.Query().Get("tipo") //Extraemos de la URL el parámetro tipoUsuario
	search := r.URL.Query().Get("search")    //Extraemos de la URL el parámetro search
	pag := r.URL.Query().Get("pagina")       //Extraemos de la URL el parámetro pagina

	if len(tipoUsuario) < 1 { //Si no encontró en la URL el parámetro tipoUsuario
		http.Error(wr, "Debe enviar el parámetro id", http.StatusBadRequest)
		return
	}

	pagTemp, err := strconv.Atoi(pag) //Convertir un alfabético a un entero
	if err != nil {
		http.Error(wr, "Debe enviar el parámetro pagina como entero mayor a 0 "+err.Error(), http.StatusBadRequest)
		return
	}

	pagina := int64(pagTemp) //Convertimos pagTemp a int64

	resultados, status := bd.ListarUsuarios(IDUsuario, pagina, search, tipoUsuario)
	if !status {
		http.Error(wr, "Error al listar usuarios", http.StatusBadRequest)
		return
	}

	wr.Header().Set("context-type", "application/json") //Indicamos que el resultado será un JSON
	wr.WriteHeader(http.StatusCreated)                  //Devolvemos un status 201
	json.NewEncoder(wr).Encode(resultados)
}
