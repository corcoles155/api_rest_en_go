package bd

import (
	"api_rest_en_go/models"

	"golang.org/x/crypto/bcrypt"
)

/* ComprobarLogin comprueba el login a la BBDD*/
func ComprobarLogin(email string, password string) (models.Usuario, bool) {
	usuario, encontrado, _ := ExisteUsuario(email)

	if !encontrado {
		return usuario, false
	}

	passwordBytes := []byte(password)      //Convierto password a un slide de bytes
	passwordBD := []byte(usuario.Password) //Convierto password recuperada de BBDD a un slide de bytes

	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)

	if err != nil {
		return usuario, false
	}

	return usuario, true
}
