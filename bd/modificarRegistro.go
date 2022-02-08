package bd

import (
	"api_rest_en_go/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*ModificarRegistro permite modificar el perfil del usuario*/
func ModificarRegistro(usuario models.Usuario, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second) //Al contexto le ponemos un timeout de 15 seg
	defer cancel()                                                           //defer se ejecuta como última instrucción. Va a dar de baja el timeout

	db := MongoCN.Database("twittor") //Para la BBDD twittor
	col := db.Collection("usuarios")  //Voy a usar la colección usuario

	registro := make(map[string]interface{})
	if len(usuario.Nombre) > 0 {
		registro["nombre"] = usuario.Nombre
	}

	if len(usuario.Apellidos) > 0 {
		registro["apellidos"] = usuario.Apellidos
	}

	if len(usuario.Email) > 0 {
		registro["email"] = usuario.Email
	}

	if len(usuario.Password) > 0 {
		registro["password"] = usuario.Password
	}

	if len(usuario.Avatar) > 0 {
		registro["avatar"] = usuario.Avatar
	}

	if len(usuario.Banner) > 0 {
		registro["banner"] = usuario.Banner
	}

	if len(usuario.Biografia) > 0 {
		registro["biografia"] = usuario.Biografia
	}

	if len(usuario.Ubicacion) > 0 {
		registro["ubicacion"] = usuario.Ubicacion
	}

	if len(usuario.SitioWeb) > 0 {
		registro["sitio_web"] = usuario.SitioWeb
	}

	registro["fecha_nacimiento"] = usuario.FechaNacimiento

	updateString := bson.M{ //bson.M corresponde a mapa de string
		"$set": registro, //Con $set actualizamos el valor de la colección
	}

	objID, _ := primitive.ObjectIDFromHex(ID) //Convertimos el string ID a ObjectID

	filtro := bson.M{ //bson.M corresponde a mapa de string
		"_id": bson.M{"$eq": objID}, //Creamos un filtro para no actualizar todos los registros de la BBDD
	}

	_, err := col.UpdateOne(ctx, filtro, updateString)
	if err != nil {
		return false, err
	}

	return true, nil

}
