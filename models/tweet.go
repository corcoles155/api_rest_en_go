package models

import "time"

type Tweet struct {
	UsuarioId string    `bson:"usuario_id" json:"usuario_id"`
	Mensaje   string    `bson:"mensaje" json:"mensaje,omitempty"`
	Fecha     time.Time `bson:"fecha" json:"fecha,omitempty"`
}
