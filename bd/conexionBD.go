package bd

import (
	"context"
	"fmt"

	"github.com/JuanRodriguez84/twitterGo/models"
	"go.mongodb.org/mongo-driver/mongo"         // paquete de mongo que se descarga con go get desde la terminal "go get go.mongodb.org/mongo-driver/mango"
	"go.mongodb.org/mongo-driver/mongo/options" // paquete de mongo que se descarga con go get desde la terminal
)

// aca se construye la conexión
// variables que se van a exportar a otros paquetes
var MongoConnection *mongo.Client // esta va a ser la conexión abierta hacia mongo
var DataBaseName string

func ConectarBD(ctx context.Context) error { // puede devolver un error real o nil

	user := ctx.Value(models.Key("user")).(string) // para transformar el valor ctx.Value(models.Key("User"))  en string es asi ->  .(string)
	password := ctx.Value(models.Key("password")).(string)
	host := ctx.Value(models.Key("host")).(string)

	// armar el string de conexion asi:
	// el primer %s va a ser el usuario, }
	// el segundo %s el password
	// el tercer es @%s va a venir el host,
	// luego sigue ? que es un parametro a la conexion retryWrites=true
	// luego el otro parametro &w=majority
	connectionString := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority", user, password, host) // fmt no solo sirve para enviar al log, tambien sirve para formatear texto, cuando empieza por S mayuscula como Sprintf significa que lo va a convertir en string  formateado

	// Ahora se arma la conexion en si misma

	// opciones que vamos a enviar de tipo options
	var clientOptions = options.Client().ApplyURI(connectionString)

	// ya podemos conectarnos
	client, error := mongo.Connect(ctx, clientOptions)

	if error != nil {
		fmt.Println("Error Connect " + error.Error())
		return error
	}

	// ahora se debe hacer el pung para ver si la conexión quedó abierta o quedo con un error que no se pudo capturar

	error = client.Ping(ctx, nil) // es como el ping que se hace a una url , el segundo argumento es preferenia de lectura el cual se envia nil

	if error != nil {
		fmt.Println("Error Ping " + error.Error())
		return error
	}

	fmt.Println("Conexion exitosa")

	MongoConnection = client

	DataBaseName = ctx.Value(models.Key("database")).(string) // aca se configura la colecion de mongo

	return nil
}

func BaseConectada() bool {

	error := MongoConnection.Ping(context.TODO(), nil) // context.TODO()  por que no se recibe de parametro el context

	return error == nil // retorna si error es igual a nil , si es igual a nil devuelve true si no false
}
