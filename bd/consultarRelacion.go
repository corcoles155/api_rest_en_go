package bd

import (
	"context"
	"fmt"
	"time"

	"api_rest_en_go/models"

	"go.mongodb.org/mongo-driver/bson"
)

/*ConsultarRelacion recupera la relacion entre dos usuarios en BBDD*/
func ConsultarRelacion(relacion models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second) //Al contexto le ponemos un timeout de 15 se
	defer cancel()

	db := MongoCN.Database("twittor") //Para la BBDD twittor
	col := db.Collection("relacion")  //Voy a usar la colección relacion

	condicion := bson.M{ //bson.M corresponde a mapa de string
		"usuario_id":          relacion.UsuarioId,
		"usuario_relacion_id": relacion.UsuarioRelacionId,
	}

	var resultado models.Relacion
	fmt.Println(resultado)

	err := col.FindOne(ctx, condicion).Decode(&resultado)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}

	return true, nil
}
