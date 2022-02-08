package bd

import (
	"api_rest_en_go/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertarRegistro(usuario models.Usuario) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second) //Al contexto le ponemos un timeout de 15 seg
	defer cancel()                                                           //defer se ejecuta como última instrucción. Va a dar de baja el timeout

	db := MongoCN.Database("twittor") //Para la BBDD twittor
	col := db.Collection("usuarios")  //Voy a usar la colección usuario

	usuario.Password, _ = EncriptarPassword(usuario.Password)

	result, err := col.InsertOne(ctx, usuario)

	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID) //Obtener usuario.ID
	return objID.String(), true, nil                   //convertirlo a string
}
