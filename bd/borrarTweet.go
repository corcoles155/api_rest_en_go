package bd

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* BorrarTweet borra un tweet determinado*/
func BorrarTweet(TweetID string, UsuarioID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second) //Al contexto le ponemos un timeout de 15 seg
	defer cancel()                                                           //defer se ejecuta como última instrucción. Va a dar de baja el timeout

	db := MongoCN.Database("twittor") //Para la BBDD twittor
	col := db.Collection("tweet")     //Voy a usar la colección tweet

	objectId, _ := primitive.ObjectIDFromHex(TweetID)

	condicion := bson.M{
		"_id":        objectId,
		"usuario_id": UsuarioID,
	}

	_, err := col.DeleteOne(ctx, condicion)
	return err
}
