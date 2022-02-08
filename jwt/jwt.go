package jwt

import (
	"api_rest_en_go/models"
	"time"

	jwt "github.com/dgrijalva/jwt-go" //Creamos el alias jwt
)

/* GenerarJWT genera el encriptado con JWT*/
func GenerarJWT(usuario models.Usuario) (string, error) {
	miClave := []byte("EstaEsMiClave")

	payload := jwt.MapClaims{
		"email":            usuario.Email,
		"nombre":           usuario.Apellidos,
		"fecha_nacimiento": usuario.FechaNacimiento,
		"biografia":        usuario.Biografia,
		"ubicacion":        usuario.Ubicacion,
		"sitio_web":        usuario.SitioWeb,
		"_id":              usuario.ID.Hex(),                      //Convertimos el ID en un hexadecimal de string
		"exp":              time.Now().Add(time.Hour * 24).Unix(), //Transformamos la fecha de expiración a formato Unix
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload) //Le pasamos el algoritmo de encriptación y el payload
	tokenStr, err := token.SignedString(miClave)                //Firmamos el token con la clave que hemos generado
	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
