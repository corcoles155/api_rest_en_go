package bd

import (
	"api_rest_en_go/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* InsertarTweet guarda un tweet en BBDD*/
func InsertarTweet(tweet models.GraboTweet) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second) //Al contexto le ponemos un timeout de 15 seg
	defer cancel()                                                           //defer se ejecuta como última instrucción. Va a dar de baja el timeout

	db := MongoCN.Database("twittor") //Para la BBDD twittor
	col := db.Collection("tweet")     //Voy a usar la colección tweet

	registro := bson.M{
		"usuarioId": tweet.UsuarioID,
		"mensaje":   tweet.Mensaje,
		"fecha":     tweet.Fecha,
	}

	result, err := col.InsertOne(ctx, registro)

	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID) //Obtener tweet.ID
	return objID.String(), true, nil                   //objID.String() formatea objID a string

}
