package models

import "time"

type GraboTweet struct {
	UsuarioID string    `bson:"usuario_id" json:"usuario_id,omitempty"`
	Mensaje   string    `bson:"mensaje" json:"mensaje,omitempty"`
	Fecha     time.Time `bson:"fecha" json:"fecha,omitempty"`
}
