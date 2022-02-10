package bd

import (
	"api_rest_en_go/models"
	"context"
	"time"
)

/* BorrarRelacion borra una relacion entre dos usuarios*/
func BorrarRelacion(relacion models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second) //Al contexto le ponemos un timeout de 15 seg
	defer cancel()                                                           //defer se ejecuta como última instrucción. Va a dar de baja el timeout

	db := MongoCN.Database("twittor") //Para la BBDD twittor
	col := db.Collection("relacion")  //Voy a usar la colección relacion

	_, err := col.DeleteOne(ctx, relacion)
	if err != nil {
		return false, err
	}

	return true, nil
}
