package bd

import (
	"context"
	"fmt"
	"time"

	"api_rest_en_go/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*BuscarPerfil busca un perfil en BBDD*/
func BuscarPerfil(ID string) (models.Usuario, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second) //Al contexto le ponemos un timeout de 15 se
	defer cancel()

	db := MongoCN.Database("twittor") //Para la BBDD twittor
	col := db.Collection("usuario")   //Voy a usar la colección usuario

	var perfil models.Usuario
	objID, _ := primitive.ObjectIDFromHex(ID) //Convertimos el string ID a ObjectID

	condicion := bson.M{ //bson.M corresponde a mapa de string
		"_id": objID,
	}

	err := col.FindOne(ctx, condicion).Decode(&perfil) //Si no encuentra el registro graba un error en la variable err y sino guarda el resultado en la variable perfil
	perfil.Password = ""                               //Eliminamos el valor password
	if err != nil {
		fmt.Println("Registro no encontrado " + err.Error())
		return perfil, err
	}

	return perfil, nil
}
