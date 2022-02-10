package models

/*Relacion modelo para guardar la relación de un usuario con otro */
type Relacion struct {
	UsuarioId         string `bson:"usuario_id" json:"usuario_id"`
	UsuarioRelacionId string `bson:"usuario_relacion_id" json:"usuario_relacion_id"`
}
