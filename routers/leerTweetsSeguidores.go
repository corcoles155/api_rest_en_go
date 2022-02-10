package routers

import (
	"api_rest_en_go/bd"
	"encoding/json"
	"net/http"
	"strconv"
)

/*LeerTweetsRelacion permite leer los tweets de un usuario*/
func LeerTweetsSeguidores(wr http.ResponseWriter, r *http.Request) {

	if len(r.URL.Query().Get("pagina")) < 1 { //Extraemos de la URL el parámetro pagina. Si no encontró en la URL el parámetro pagina devuelve un error
		http.Error(wr, "Debe enviar el parámetro pagina", http.StatusBadRequest)
		return
	}

	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina")) //Convertir un alfabético a un entero
	if err != nil {
		http.Error(wr, "El formato del parámetro página no es correcto. Debe ser un número "+err.Error(), http.StatusBadRequest)
		return
	}

	respuesta, correcto := bd.LeerTweetsSeguidores(IDUsuario, pagina)
	if !correcto {
		http.Error(wr, "Error al leer los tweets de los seguidores", http.StatusBadRequest)
		return
	}

	wr.Header().Set("context-type", "application/json") //Indicamos que el resultado será un JSON
	wr.WriteHeader(http.StatusCreated)                  //Devolvemos un status 201
	json.NewEncoder(wr).Encode(respuesta)
}
