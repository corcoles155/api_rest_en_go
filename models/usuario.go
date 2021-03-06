package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* Usuario */
type Usuario struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"` //Indicamos los datos de entrada a la BBDD y los datos de salida de la BBDD. omitempty, si el campo viene vacío lo omite. Se va a guardar en Mongo como _id
	Nombre          string             `bson:"nombre" json:"nombre,omitempty"`
	Apellidos       string             `bson:"apellidos" json:"apellidos,omitempty"`
	FechaNacimiento time.Time          `bson:"fecha_nacimiento" json:"fecha_nacimiento,omitempty"`
	Email           string             `bson:"email" json:"email"`
	Password        string             `bson:"password" json:"password,omitempty"`
	Avatar          string             `bson:"avatar" json:"avatar,omitempty"`
	Banner          string             `bson:"banner" json:"banner,omitempty"`
	Biografia       string             `bson:"biografia" json:"biografia,omitempty"`
	Ubicacion       string             `bson:"ubicacion" json:"ubicacion,omitempty"`
	SitioWeb        string             `bson:"sitio_web" json:"sitio_web,omitempty"`
}
