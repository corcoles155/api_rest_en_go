package bd

import (
	"context"
	"log"
	"time"

	"api_rest_en_go/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*LeerTweets recupera todos los tweets por usuario de BBDD*/
func LeerTweets(UsuarioID string, pagina int64) ([]*models.DevuelvoTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second) //Al contexto le ponemos un timeout de 15 se
	defer cancel()

	db := MongoCN.Database("twittor") //Para la BBDD twittor
	col := db.Collection("tweet")     //Voy a usar la colección tweet

	var resultados []*models.DevuelvoTweets

	condicion := bson.M{ //bson.M corresponde a mapa de string
		"usuario_id": UsuarioID,
	}

	opciones := options.Find()
	opciones.SetLimit(20)                               //Indicamos que el máximo de documentos mostrados será 20
	opciones.SetSort(bson.D{{Key: "fecha", Value: -1}}) //Ordenamos de forma descendente por fecha
	opciones.SetSkip((pagina - 1) * 20)                 //Mi limite de paginas es 20. Si pagina vale 1 el resultado será 0, no voy a saltar ningún documento. Si pagina vale 2 el resultado será 20, saltará los 20 primeros documentos y así sucesivamente.

	cursor, err := col.Find(ctx, condicion, opciones)
	if err != nil {
		log.Fatal(err.Error())
		return resultados, false
	}

	for cursor.Next(context.TODO()) { //context.TODO() crea un contexto vacío
		var tweet models.DevuelvoTweets
		err := cursor.Decode(&tweet) //Si hubo un error lo graba en la variable err sino en la variable tweet
		if err != nil {
			return resultados, false
		}
		resultados = append(resultados, &tweet) //Por cada iteración añade a resultados un tweet
	}

	return resultados, true
}
