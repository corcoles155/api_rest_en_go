package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*DevuelvoTweets es la estructura con la que devolveremos los Tweets*/
type DevuelvoTweets struct {
	ID        primitive.ObjectID `bson:"_id" json:"id,omitempty"` //Indicamos los datos de entrada a la BBDD y los datos de salida de la BBDD. omitempty, si el campo viene vacío lo omite. Se va a guardar en Mongo como _id
	UsuarioID string             `bson:"usuario_id" json:"usuario_id,omitempty"`
	Mensaje   string             `bson:"mensaje" json:"mensaje,omitempty"`
	Fecha     time.Time          `bson:"fecha" json:"fecha,omitempty"`
}
