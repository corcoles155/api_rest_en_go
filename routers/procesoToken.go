package routers

import (
	"api_rest_en_go/bd"
	"api_rest_en_go/models"
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go" //Creamos el alias jwt
)

/*Email es el email devuelto del modelo, que se usará en todos los endpointss*/
var Email string

/*IDUsuario es el ID devuelto del modelo, que se usará en todos los endpoints*/
var IDUsuario string

/*ProcesoToken para extraer los valores del token*/
func ProcesoToken(token string) (*models.Claim, bool, string, error) {
	miClave := []byte("EstaEsMiClave")
	claims := &models.Claim{}

	splitToken := strings.Split(token, "Bearer") //Queremos eliminar la palabra Bearer, esto nos va a generar un array de dos posiciones, en la primera vendrá Bearer y en la segunda el token
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	token = strings.TrimSpace(splitToken[1]) //Elimino los espacios que envuelven el valor del token

	//Validamos que el token sea válido
	tkn, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return miClave, nil
	})

	if err == nil {
		_, encontrado, _ := bd.ExisteUsuario(claims.Email)
		if encontrado {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, encontrado, IDUsuario, nil
	}

	if !tkn.Valid { //Si el token no fue válido
		return claims, false, string(""), errors.New("token invalido")
	}

	return claims, false, string(""), err
}
