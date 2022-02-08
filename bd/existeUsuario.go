package bd

import (
	"api_rest_en_go/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

/* ExisteUsuario recibe un email y comprueba si ya está en BBDD */
func ExisteUsuario(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second) //Al contexto le ponemos un timeout de 15 se
	defer cancel()                                                           //defer se ejecuta como última instrucción. Va a dar de baja el timeout

	db := MongoCN.Database("twittor") //Para la BBDD twittor
	col := db.Collection("usuario")   //Voy a usar la colección usuario

	condicion := bson.M{"email": email} //bson.M corresponde a mapa de string

	var result models.Usuario

	err := col.FindOne(ctx, condicion).Decode(&result) //Si no encuentra el registro graba un error en la variable err y sino guarda el resultado en la variable result
	ID := result.ID.Hex()                              //Convierto el ID a un hexadecimal en formato string
	if err != nil {
		return result, false, ID
	}

	return result, true, ID

}
