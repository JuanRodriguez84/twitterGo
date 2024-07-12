package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Usuario struct { // Define un nuevo tipo de dato llamado usuario. En Go, struct se utiliza para definir una estructura que agrupa juntos diferentes tipos de datos relacionados.
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"` // bson que indica cómo se deben serializar y deserializar los datos cuando se trabaja con MongoDB o BSON (Binary JSON).
	// "_id" es el nombre del campo en la base de datos.
	// json:"id": Esta es una etiqueta json que indica cómo se deben serializar y deserializar los datos cuando se trabaja con JSON. "id" es el nombre del campo en la representación JSON.
	Nombre          string    `bson:"nombre" json:"nombre,omitempty"`
	Apellidos       string    `bson:"apellidos" json:"apellidos,omitempty"`
	FechaNacimiento time.Time `bson:"fechaNacimiento" json:"fechaNacimiento,omitempty"`
	Email           string    `bson:"email" json:"email"`
	Password        string    `bson:"password" json:"password,omitempty"`
	Avatar          string    `bson:"avatar" json:"avatar,omitempty"`
	Banner          string    `bson:"banner" json:"banner,omitempty"`
	Biografia       string    `bson:"biografia" json:"biografia,omitempty"` // Biografia información del perfil
	Ubicacion       string    `bson:"ubicacion" json:"ubicacion,omitempty"`
	SitioWeb        string    `bson:"sitioweb" json:"sitioweb,omitempty"`
}
