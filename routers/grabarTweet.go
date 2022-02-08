package routers

import (
	"api_rest_en_go/bd"
	"api_rest_en_go/models"
	"encoding/json"
	"net/http"
	"time"
)

/*GrabarTweet permite grabar el tweet en BBDD*/
func GrabarTweet(wr http.ResponseWriter, r *http.Request) {
	var tweet models.Tweet

	err := json.NewDecoder(r.Body).Decode(&tweet) //El objeto Body es un Stream, sólo se puede leer una vez

	if err != nil {
		http.Error(wr, "Error en los datos recibidos "+err.Error(), http.StatusBadRequest)
		return
	}

	graboTweet := models.GraboTweet{
		UsuarioID: IDUsuario, //IDUsuario se corresponde a la variable global que declaramos en procesoToken
		Mensaje:   tweet.Mensaje,
		Fecha:     time.Now(),
	}

	_, status, err := bd.InsertarTweet(graboTweet)
	if err != nil {
		http.Error(wr, "Ocurrió un error al intentar insertar el registro. Reintente de nuevo "+err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(wr, "No se ha logrado insertar el tweet", http.StatusBadRequest)
		return
	}

	wr.WriteHeader(http.StatusCreated) //Devolvemos un 201
}
