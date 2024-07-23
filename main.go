package main

import (
	"context"
	"fmt"
	"os" // de sistema operativo por que se va a incorporar el tema de las variables de entorno ya que alas lambdas tienen variables de entorno
	"strings"

	// os permite leer las variables de entorno
	lambda "github.com/aws/aws-lambda-go/lambda" // paquete permite trabajar con lamdas para hacer un go get
	// github.com/aws/aws-lambda-go/lambda    es colocar en la terminal  go get github.com/aws/aws-lambda-go/lambda  para instalar el paquete  y adiciona al go.mod el git en la version v1.47.0
	"github.com/aws/aws-lambda-go/events" // todo el manejo de eventos dentro de Amazon se maneja con este paquete  tambien se debe descargar desde la terminal con   go get github.com/aws/aws-lambda-go/events
	//events : maneja tambien los de ALB para load balancer
	"github.com/JuanRodriguez84/twitterGo/awsgo"
	"github.com/JuanRodriguez84/twitterGo/bd"
	"github.com/JuanRodriguez84/twitterGo/handlers"
	"github.com/JuanRodriguez84/twitterGo/models"
	"github.com/JuanRodriguez84/twitterGo/secretmanager"
)

func main() {
	// Cuando se trabaja con lambdas lo que nos pide es que main llame a una función y le pase un parametro
	// En AWS configuramos el controlador colocamos main, con eso le indicamos que cuando la lambda se ejecute y sea llamada por la API Gateway este main va a ser la puerta de entrada
	lambda.Start(EjecutoLambda) // inicializa el handler de la lambda, se pasa el nombre de la función "EjecutoLambda" como parametro para procesar

}

// 1 parametro es un context
// 2 parametro que recibe es la lambda, la lambda va a enviar de parametro al codigo el API Gateway un objeto de API que da Amazon
// APIGatewayProxyRequest cuando es una API REST
// devueve un puntero a APIGatewayProxyResponse y error en caso de que falle
func EjecutoLambda(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) { // entre parentesis colocar los parametros que va a recibir
	// cuando llamemos a nuestra API Gateway desde la aplicación o desde postman la Api Gateway se va a conectar con mi lambda y le va a enviar un objeto de tipo API GatewayProxyRequesty
	// cuando es http se envia un objeto events.APIGatewayV2HTTPRequest

	var respuesta *events.APIGatewayProxyResponse // para interactuar con la variable respuesta

	awsgo.InicializoAWS()

	if !ValidoParametros() {
		respuesta = &events.APIGatewayProxyResponse{ // & obtener la dirección de memoria
			StatusCode: 400, // puede ser error 500
			Body:       "Error en las variables de entorno, deben incluir SecretName, BucketName y UrlPrefix",
			Headers: map[string]string{
				"Content-Type": "application/json", // no siempre es application/json
			},
		}
		return respuesta, nil
	}

	// la fucnion GetSecret devuleve un tipo de dato de estrutura y un error
	SecretModel, err := secretmanager.GetSecret(os.Getenv("SecretName")) // os.Getenv  para obtener el valor de la variable de entorno

	if err != nil {
		respuesta = &events.APIGatewayProxyResponse{ // & obtener la dirección de memoria
			StatusCode: 400, // puede ser error 500
			Body:       "Error en la lectura de secret " + err.Error(),
			Headers: map[string]string{
				"Content-Type": "application/json", // no siempre es application/json
			},
		}
		return respuesta, nil
	}

	// aca se trabaja con el context

	// el path viene en un objeto del request llamado PathParameters, en PathParameters viene "twitterGo/login"  que le quite "twitterGo/""
	path := strings.Replace(request.PathParameters["twittergo"], os.Getenv("UrlPrefix"), "", -1) // remplazar lo que esta en la variable de entorno UrlPrefix por nada "" y que comience desde -1 para que busque desde el inicio de la cadena del string
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("path"), path)

	//empezamos creando variables en el contxeto con el paquete context.WithValue()

	// el primer parametro es el context}
	// segundo es la clave, que debe ser de un tipo de dato en particular, pero lo que pide GO es no utilizar explicitamente un tipo de dato string para una clave, tenemos que crear un tipo de dato propio
	// tercero es el valor
	//awsgo.Ctx = context.WithValue(awsgo.Ctx, "method", request.HTTPMethod)  // method :  si es un POST, PUT, GET, etc dejarlo en "method" genera inconvenientes por eso se debe crear un tipo de dato
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("method"), request.HTTPMethod)
	// el context esta por encima de toda la lambda y ademas si se tiene una serializacion de lambdas que una lambda tiene que ejecutar otra lambda, el context exporta esas variable
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("user"), SecretModel.Username) // user viene del secretmanager
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("password"), SecretModel.Password)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("host"), SecretModel.Host)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("database"), SecretModel.Database)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("jwtSign"), SecretModel.JWTSign)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("body"), request.Body)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("bucketName"), os.Getenv("BucketName")) // las imagenes bienen codificadas en base 64 pero vienen como texto

	// en el context se grabo todas las variables de ejecución que voy a necesitar

	// Chequeo conexión a la base de datos o conecto la base de datos
	err = bd.ConectarBD(awsgo.Ctx)

	if err != nil {
		respuesta = &events.APIGatewayProxyResponse{ // & obtener la dirección de memoria
			StatusCode: 500, // puede ser error 500
			Body:       "Error conectando en la Base de datos" + err.Error(),
			Headers: map[string]string{
				"Content-Type": "application/json", // no siempre es application/json
			},
		}
		return respuesta, nil
	}

	RespAPI := handlers.Manejadores(awsgo.Ctx, request)

	if RespAPI.CustomResponse == nil { // si no viene aramado , por que cuando procesa imagenes ya viene armado un CustomResponse personalizadp para essa API
		respuesta = &events.APIGatewayProxyResponse{ // & obtener la dirección de memoria
			StatusCode: RespAPI.Status,
			Body:       RespAPI.Message,
			Headers: map[string]string{
				"Content-Type": "application/json", // no siempre es application/json
			},
		}
		return respuesta, nil
	} else {
		return RespAPI.CustomResponse, nil
	}

}

// En Go (Golang), los símbolos * y & tienen significados diferentes y se utilizan para trabajar con punteros
// *events.APIGatewayProxyResponse es un tipo de dato que es un puntero,  * Una variable de tipo puntero es una variable que almacena la dirección de memoria de otra variable en lugar de almacenar directamente su valor. En lenguajes de programación como Go (Golang), C, C++, y otros, los punteros son útiles para trabajar con datos de manera eficiente, especialmente cuando se trata de pasar datos grandes o estructuras complejas a funciones, o cuando se necesita modificar el valor de una variable desde varias ubicaciones en el programa.
// mientras que &events.APIGatewayProxyResponse es el operador que obtiene un puntero a una instancia específica
// '&' se utiliza para obtener la dirección de memoria de una variable o expresión y devolver un puntero al tipo de dato correspondiente.

// Es obligatorio que la lambda reciba 3 variables de entorno
func ValidoParametros() bool {

	_, traeParametro := os.LookupEnv("SecretName") // LookupEnv permite validar si viene la variable

	if !traeParametro {
		return traeParametro
	}

	_, traeParametro = os.LookupEnv("BucketName") // se trae por variable de entorno el nombre del bucket S3 y asi no tenerlo hard code

	if !traeParametro {
		return traeParametro
	}

	_, traeParametro = os.LookupEnv("UrlPrefix") // UrlPrefix para quitarle el prefijo a la URL y ese prefijo se configura en las variables de entorno

	if !traeParametro {
		return traeParametro
	}

	return traeParametro

}

// Ejemplos punteros
func punteros() {
	var x int = 10    // Declara una variable int
	var ptr *int = &x // Declara un puntero a int y asigna la dirección de memoria de x , 

	// Declaración de puntero: Cuando se declara una variable como puntero, el asterisco (*) se 
	// coloca antes del tipo de dato para indicar que la variable es un puntero a ese tipo de dato. 
	// Por ejemplo:   var ptr *int
	// En este caso, ptr es un puntero a un valor de tipo int. Esto significa que ptr almacenará la dirección de memoria 
	// donde se encuentra almacenado un valor de tipo int.

	fmt.Println("Valor de x:", x)                // Imprime el valor de x    -> "10"
	fmt.Println("Dirección de x:", &x)           // Imprime la dirección de memoria de x       -> "0xc0000a0a8 u otra dirección"

	// &x es la dirección de memoria donde está almacenado el valor de x. Esto es útil cuando se desea pasar la dirección de una 
	// variable (un puntero) a funciones o cuando se desea asignar esa dirección a otra variable que sea de tipo puntero.

	fmt.Println("Valor apuntado por ptr:", *ptr) // Imprime el valor apuntado por ptr (dereferenciación)    -> "10"

	// Dereferenciación: Cuando se usa el asterisco (*) como operador de dereferenciación, se coloca antes de un puntero para obtener 
	// el valor almacenado en la dirección de memoria apuntada por ese puntero. Por ejemplo:
	// fmt.Println("Valor apuntado por ptr:", *ptr)
	// Aquí, *ptr obtiene el valor de tipo int almacenado en la dirección de memoria apuntada por ptr.
	
	*ptr = 20 // Modifica el valor de x a través de ptr (dereferenciación y asignación)

	fmt.Println("Nuevo valor de x:", x) // Imprime el nuevo valor de x     -> "20"

	// - ptr *int se coloca antes del tipo de dato para indicar que la variable es un puntero a ese tipo de dato. Esto significa  
	// 	          que ptr almacenará la dirección de memoria donde se encuentra almacenado un valor de tipo int.
	
	// - *ptr obtener el valor de un tipo (definido en la declaracion) almacenada en al dirección de memoria apuntada por ptr
	// - &x obtener la dirección de memoria donde está almacenado el valor de x Esto es útil cuando se desea pasar la dirección de 
	//      una variable (un puntero) a funciones o cuando se desea asignar esa dirección a otra variable que sea de tipo puntero.
}
