package bd

import (
	"context"

	"github.com/JuanRodriguez84/twitterGo/models"
	"go.mongodb.org/mongo-driver/bson/primitive" // bson primitive a converir
)

func InsertoRegistro(usuario models.Usuario) (string, bool, error) {

	ctx := context.TODO() // TODO devuelve un context vacio

	// traer la BD a trabajar, para indicar con que BD voy a trabajar
	db := MongoConnection.Database(DataBaseName) // MongoConnection y DataBaseName son variables publicas de bd

	//configurar la colección con la que voy a trabajar
	collection := db.Collection("usuarios") // la coleccion en bd se llama usuarios

	// Encriptar la password

	usuario.Password, _ = EncriptarPassword(usuario.Password)

	result, err := collection.InsertOne(ctx, usuario) // InsertOne  para insertar un solo registro

	if err != nil {
		return "", false, err
	}

	// cuando se devuelva al postman la respuesta o al front se devuelve el id del usuario

	ObjectID, _ := result.InsertedID.(primitive.ObjectID) // esult.InsertedID  se debe convertir por que viene en un formato primitivo y se devuelve un string

	return ObjectID.String(), true, nil

}
