package awsgo

// Este paquete se crea por que permite conectarnos con AWS, apsesar de estar en una lambda y apesar de que la lambda está dentro de Amazon
// tenemos que iniciar sesión, ya que la lambda está corriendo dentro del entorno de Amzon pero tiene que reconocer al entono de Amazon, esto
// no es automatico

import (
	"context" // paquete que ya viene con go, es nativo

	"github.com/aws/aws-sdk-go-v2/aws"// paquete de Amazon , version 2 para el sdk de go se ejecuta en la terminal  --> go get github.com/aws/aws-sdk-go-v2/aws
	"github.com/aws/aws-sdk-go-v2/config"// paquete de Amazon se ejecuta en la terminal  --> go get github.com/aws/aws-sdk-go-v2/config
)

var Ctx context.Context
var Cfg aws.Config    // paquete de amazon
var err error

func InicializoAWS(){
	// se configuran las variables  Ctx y Cfg  para arrancar la configuración con Amazon

	Ctx = context.TODO()  // context.TODO() crea un context vacio
	Cfg, err = config.LoadDefaultConfig(Ctx, config.WithDefaultRegion("us-east-1"))  // esto permite conectarnos a Amazon
	if err != nil {
		panic("Error al cargar la configuración ./aws/config" + err.Error()) // aborta la lambda y graba en cloudwatch el mensaje
	}
}