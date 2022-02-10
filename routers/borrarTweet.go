package routers

import (
	"api_rest_en_go/bd"
	"net/http"
)

/*EliminarTweet*/
func EliminarTweet(wr http.ResponseWriter, r *http.Request) {
	tweetId := r.URL.Query().Get("id")
	if len(tweetId) < 1 {
		http.Error(wr, "Debe enviar el parámetro id", http.StatusBadRequest)
		return
	}

	err := bd.BorrarTweet(tweetId, IDUsuario)
	if err != nil {
		http.Error(wr, "Ocurrió un error al borrar el tweet "+err.Error(), http.StatusBadRequest)
		return
	}

	wr.Header().Set("context-type", "application/json") //Indicamos que el resultado será un JSON
	wr.WriteHeader(http.StatusCreated)
}
