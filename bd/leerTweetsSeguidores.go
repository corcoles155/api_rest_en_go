package bd

import (
	"api_rest_en_go/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func LeerTweetsSeguidores(id string, pagina int) ([]models.DevuelvoTweetsSeguidores, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second) //Al contexto le ponemos un timeout de 15 se
	defer cancel()

	db := MongoCN.Database("twittor") //Para la BBDD twittor
	col := db.Collection("relacion")  //Voy a usar la colección relacion

	skip := (pagina - 1) * 20 //para paginar los tweets

	condiciones := make([]bson.M, 0)
	condiciones = append(condiciones, bson.M{"$match": bson.M{"usuario_id": id}})
	condiciones = append(condiciones, bson.M{"$loolup": bson.M{"from": "tweet", "localField": "usuario_relacion_id", "foreignField": "usuario_id", "as": "tweet"}}) //Unir tablas en MongoDB con framework aggregate, en el parámetro from ponemos con que tabla queremos unir nuestra tabla, en el parámetro localfield ponemos cual es el campo de la tabla origen (en nuestro caso relacion) con el que vamos a unir las tablas y en el parámetro foreignField cual es el campo de la tabla destino (en nuestro caso tweet). El parámetro as corresponde a un alias para nuestra tabla destino
	condiciones = append(condiciones, bson.M{"$unwind": "tweet"})                                                                                                   //todos los documentos vendrán con el mismo formato. Vendrá la info repetida en lugar de un maestro con subdocumentos
	condiciones = append(condiciones, bson.M{"$sort": bson.M{"fecha": -1}})                                                                                         //Ordenamos de fecha de mayor a menor (1) o de menor a mayor (-1)
	condiciones = append(condiciones, bson.M{"$skip": skip})                                                                                                        //Cuantos registros tengo que saltar para paginar
	condiciones = append(condiciones, bson.M{"$limit": 20})                                                                                                         //Cuantos registros voy a mostrar al paginar

	cursor, _ := col.Aggregate(ctx, condiciones)
	var resultados []models.DevuelvoTweetsSeguidores
	err := cursor.All(ctx, &resultados) //Procesamos el cursor y decodifica el valor en resultados
	if err != nil {
		return resultados, false
	}

	return resultados, true
}
