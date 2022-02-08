package bd

import "golang.org/x/crypto/bcrypt"

/* EncriptarPassword función para encriptar valores*/
func EncriptarPassword(pass string) (string, error) {
	costo := 8                                                     //cantidad de encriptaciones que va a hacer la func GenerateFromPassword para encriptar la contraseña. Cuantas más encriptaciones mayor seguridad
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo) //Convertimos pass a un slice de bytes []byte(pass)
	return string(bytes), err                                      //Convierto el slice de bytes a string
}
