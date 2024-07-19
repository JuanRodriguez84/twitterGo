package bd

import (
	"context"

	"github.com/JuanRodriguez84/twitterGo/models"
	"go.mongodb.org/mongo-driver/bson" // se le debe hacer un go get
)

func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {
	// se crea un context vacio, por que hay una función em mongo que pide un context, pero aca no interesa el context que tenemmos de amazon
	ctx := context.TODO() // TODO devuelve un context vacio

	// traer la BD a trabajar, para indicar con que BD voy a trabajar
	db := MongoConnection.Database(DataBaseName) // MongoConnection y DataBaseName son variables publicas de bd

	//configurar la colección con la que voy a trabajar
	collection := db.Collection("usuarios") // la coleccion en bd se llama usuarios

	// validar si existe el usuario
	condition := bson.M{"email": email} // "M" es una interface especifica de bson y tiene un formato de clave-valo, para este caso {"email": email} que es que "email" va a ser filtrado por el campo que llega email, es como un Where en DB SQL

	var resultado models.Usuario // completa los datos cuando los lee de mongo

	err := collection.FindOne(ctx, condition).Decode(&resultado) // FindOne se utiliza en este caso para que triga el primero que encuentre. la funcion Decode es para decodificar el resultado y lo convierte en una estructura de GO

	ID := resultado.ID.Hex() // tomma el valor hexadecimal y lo convierte a string
	if err != nil {
		return resultado, false, ID
	}

	return resultado, true, ID // true significa que el usuario ya existe
}
