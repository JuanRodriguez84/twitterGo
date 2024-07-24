package bd

import (
	"context"

	"github.com/JuanRodriguez84/twitterGo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func BuscoPerfil(ID string) (models.Usuario, error) {

	context := context.TODO() // creación de context vacio con el TODO
	db := MongoConnection.Database(DataBaseName)
	collection := db.Collection("usuarios")

	var perfil models.Usuario

	// convertir el id en primitivo, por que vamos a buscar dentro de mongo un Id determinado para ello se debe cambiar de string a primitive hex
	objectID, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{
		"_id": objectID, // la condicion es que _id sea igual a objectID
	}

	// esa condición la usamos en la función de MongoDB

	err := collection.FindOne(context, condicion).Decode(&perfil) // que decodifique en el modelo usuario llamado perfil

	// el password no interesa enviarlo al front
	perfil.Password = "" //de esta manera el omitEmpty detecta la cadena vacia y no incluye el campo en el Json

	if err != nil {
		return perfil, err
	}

	return perfil, nil
}
