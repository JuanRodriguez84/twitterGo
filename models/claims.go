package models

// el claims se refiere al payload

import (
	jwt "github.com/golang-jwt/jwt/v5"           // esta en la pagina de jwt para go , se debe hacer el respectivo go get en la terminal
	"go.mongodb.org/mongo-driver/bson/primitive" // se usa para el id que uno tiene en la BD como clave principal de mongo, es un tipo de dato especial que tiene este formato
	// esta en formato bson, que es una varientae de json pero encriptado
	// se importa la libreria con go get "go get go.mongodb.org/mongo-driver/bson/primitive"
)

type Claim struct { // se puede tenet muchos campos en el payload, pero vamos a tranajar con estos 3
	Email                string             `json:"email"`
	ID                   primitive.ObjectID `bson:"_id" json:"_id,omitempty"` // en mongo el id lo coloca con guion bajo, cuando se convierta el bson se coloca ya en formato json, omitempty es para que omita el tag en caso de que el dato este vacio y  por ende no lo incluye en el json
	jwt.RegisteredClaims                    // esto indica que el resto del json lo complete con el RegisteredClaims el cual contiene la expiracion de token y otros campos adicionales
}
