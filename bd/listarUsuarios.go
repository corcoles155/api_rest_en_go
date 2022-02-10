package bd

import (
	"context"
	"fmt"
	"time"

	"api_rest_en_go/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*ListarUsuarios recupera todos los usuarios de BBDD*/
func ListarUsuarios(usuarioID string, pagina int64, search string, tipoDeUsuario string) ([]*models.Usuario, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second) //Al contexto le ponemos un timeout de 15 se
	defer cancel()

	db := MongoCN.Database("twittor") //Para la BBDD twittor
	col := db.Collection("usuarios")  //Voy a usar la colección usuarios

	var resultados []*models.Usuario

	query := bson.M{ //bson.M corresponde a mapa de string
		"nombre": bson.M{"$regex": `(?i)` + search}, //Nombres que contienen el literal que venga en la variable search
	}

	findOptions := options.Find()
	findOptions.SetLimit(20)               //Indicamos que el máximo de documentos mostrados será 20
	findOptions.SetSkip((pagina - 1) * 20) //Mi limite de paginas es 20. Si pagina vale 1 el resultado será 0, no voy a saltar ningún documento. Si pagina vale 2 el resultado será 20, saltará los 20 primeros documentos y así sucesivamente.

	cursor, err := col.Find(ctx, query, findOptions)
	if err != nil {
		fmt.Println(err.Error())
		return resultados, false
	}

	var encontrado, incluir bool
	for cursor.Next(context.TODO()) { //context.TODO() crea un contexto vacío. Next me permite avanzar al siguiente registro
		var usuario models.Usuario
		err := cursor.Decode(&usuario) //Si hubo un error lo graba en la variable err sino en la variable usuario
		if err != nil {
			fmt.Println(err.Error())
			return resultados, false
		}
		var relacion models.Relacion
		relacion.UsuarioId = usuarioID
		relacion.UsuarioRelacionId = usuario.ID.Hex()

		incluir = false

		encontrado, _ = ConsultarRelacion(relacion)
		if tipoDeUsuario == "new" && !encontrado { //Obtener los usuarios a los que NO sigo
			incluir = true
		}
		if tipoDeUsuario == "follow" && encontrado { //Obtener los usuarios a los que sigo
			incluir = true
		}
		if relacion.UsuarioRelacionId == usuarioID { //Si por error un error se sigue a si mismo no lo incluyo
			incluir = false
		}
		if incluir { //Elimino los usuarios que no quiero incluir en el listado, como tienen omitempty no los va a mostrar
			usuario.Password = ""
			usuario.Biografia = ""
			usuario.SitioWeb = ""
			usuario.Ubicacion = ""
			usuario.Banner = ""
			usuario.Email = ""

			resultados = append(resultados, &usuario) //Agrego usuario al slide de resultados
		}
	}

	err = cursor.Err()
	if err != nil {
		fmt.Println(err.Error())
		return resultados, false
	}
	cursor.Close(ctx)

	return resultados, true
}
