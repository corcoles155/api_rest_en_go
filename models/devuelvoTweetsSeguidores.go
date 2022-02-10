package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*DevuelvoTweetsSeguidores*/
type DevuelvoTweetsSeguidores struct {
	ID                primitive.ObjectID `bson:"_id" json:"id,omitempty"` //Indicamos los datos de entrada a la BBDD y los datos de salida de la BBDD. omitempty, si el campo viene vacío lo omite. Se va a guardar en Mongo como _id
	UsuarioID         string             `bson:"usuario_id" json:"usuario_id,omitempty"`
	UsuarioRelacionId string             `bson:"usuario_relacion_id" json:"usuario_relacion_id,omitempty"`
	Tweet             struct {
		Mensaje string    `bson:"mensaje" json:"mensaje,omitempty"`
		Fecha   time.Time `bson:"fecha" json:"fecha,omitempty"`
		ID      string    `bson:"_id" json:"usuario_id,omitempty"`
	}
}
