package routers

import (
	"api_rest_en_go/bd"
	"encoding/json"
	"net/http"
	"strconv"
)

/*LeerTweets permite leer los tweets de un usuario*/
func LeerTweets(wr http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id") //Extraemos de la URL el parámetro ID
	if len(id) < 1 {              //Si no encontró en la URL el parámetro id
		http.Error(wr, "Debe enviar el parámetro id", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("pagina")) < 1 { //Extraemos de la URL el parámetro pagina. Si no encontró en la URL el parámetro pagina devuelve un error
		http.Error(wr, "Debe enviar el parámetro pagina", http.StatusBadRequest)
		return
	}

	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina")) //Convertir un alfabético a un entero
	if err != nil {
		http.Error(wr, "El formato del parámetro página no es correcto. Debe ser un número "+err.Error(), http.StatusBadRequest)
		return
	}

	pag := int64(pagina)

	respuesta, correcto := bd.LeerTweets(id, pag)
	if !correcto {
		http.Error(wr, "Error al leer los tweets", http.StatusBadRequest)
		return
	}

	wr.Header().Set("context-type", "application/json") //Indicamos que el resultado será un JSON
	wr.WriteHeader(http.StatusCreated)                  //Devolvemos un status 201
	json.NewEncoder(wr).Encode(respuesta)
}
