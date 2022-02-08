package routers

import (
	"api_rest_en_go/bd"
	"api_rest_en_go/models"
	"encoding/json"
	"net/http"
)

/* ModificarPerfil modifica el perfil del usuario */
func ModificarPerfil(wr http.ResponseWriter, r *http.Request) {
	var usuario models.Usuario
	err := json.NewDecoder(r.Body).Decode(&usuario) //Leemos el stream body y lo guardamos en la variable usuario. El objeto Body es un Stream, sólo se puede leer una vez
	if err != nil {
		http.Error(wr, "Datos incorrectos "+err.Error(), http.StatusBadRequest)
		return
	}

	var status bool
	status, err = bd.ModificarRegistro(usuario, IDUsuario) //la variable global IDUsuario está creada en procesoToken.go
	if err != nil {
		http.Error(wr, "Ocurrió un error al intentar modificar el registro. Reintente de nuevo "+err.Error(), http.StatusBadRequest)
		return
	}
	if !status {
		http.Error(wr, "No se ha logrado modificar el resgitro del usuario "+err.Error(), http.StatusBadRequest)
		return
	}

	wr.WriteHeader(http.StatusCreated) //Devolvemos un 201
}
