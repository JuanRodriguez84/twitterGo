package main

import (
	"context"
	"fmt"
	"os" // de sistema operativo por que se va a incorporar el tema de las variables de entorno ya que alas lambdas tienen variables de entorno
		 // os permite leer las variables de entorno
	lambda "github.com/aws/aws-lambda-go/lambda" // paquete permite trabajar con lamdas para hacer un go get
	// github.com/aws/aws-lambda-go/lambda    es colocar en la terminal  go get github.com/aws/aws-lambda-go/lambda  para instalar el paquete  y adiciona al go.mod el git en la version v1.47.0
	"github.com/aws/aws-lambda-go/events"  // todo el manejo de eventos dentro de Amazon se maneja con este paquete  tambien se debe descargar desde la terminal con   go get github.com/aws/aws-lambda-go/events
	//events : maneja tambien los de ALB para load balancer
)

func main(){
	// Cuando se trabaja con lambdas lo que nos pide es que main llame a una función y le pase un parametro 
	// En AWS configuramos el controlador colocamos main, con eso le indicamos que cuando la lambda se ejecute y sea llamada por la API Gateway este main va a ser la puerta de entrada
	lambda.Start(EjecutoLambda) // inicializa el handler de la lambda, se pasa el nombre de la función "EjecutoLambda" como parametro para procesar 

}


// 1 parametro es un context
// 2 parametro que recibe es la lambda, la lambda va a enviar de parametro al codigo el API Gateway un objeto de API que da Amazon
// APIGatewayProxyRequest cuando es una API REST
// devueve un puntero a APIGatewayProxyResponse y error en caso de que falle
func EjecutoLambda(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {// entre parentesis colocar los parametros que va a recibir
// cuando llamemos a nuestra API Gateway desde la aplicación o desde postman la Api Gateway se va a conectar con mi lambda y le va a enviar un objeto de tipo API GatewayProxyRequesty 
// cuando es http se envia un objeto events.APIGatewayV2HTTPRequest

var respuesta *events.APIGatewayProxyResponse  // para interactuar con la variable respuesta

	if !ValidoParametros(){
		respuesta = &events.APIGatewayProxyResponse{  // & obtener la dirección de memoria
			StatusCode: 400,   // puede ser error 500
			Body: "Error en las variables de entorno, deben incluir SecretName, BucketName y UrlPrefix",
			Headers: map[string]string{
				"Content-Type": "application/json", // no siempre es application/json	
			},
		}
		return respuesta, nil
	}

}

// En Go (Golang), los símbolos * y & tienen significados diferentes y se utilizan para trabajar con punteros
// *events.APIGatewayProxyResponse es un tipo de dato que es un puntero,  * Una variable de tipo puntero es una variable que almacena la dirección de memoria de otra variable en lugar de almacenar directamente su valor. En lenguajes de programación como Go (Golang), C, C++, y otros, los punteros son útiles para trabajar con datos de manera eficiente, especialmente cuando se trata de pasar datos grandes o estructuras complejas a funciones, o cuando se necesita modificar el valor de una variable desde varias ubicaciones en el programa.
// mientras que &events.APIGatewayProxyResponse es el operador que obtiene un puntero a una instancia específica 
// '&' se utiliza para obtener la dirección de memoria de una variable o expresión y devolver un puntero al tipo de dato correspondiente.


// Es obligatorio que la lambda reciba 3 variables de entorno
func ValidoParametros() bool{

	_ , traeParametro := os.LookupEnv("SecretName")  // LookupEnv permite validar si viene la variable

	if !traeParametro{
		return traeParametro
	}

	_ , traeParametro = os.LookupEnv("BucketName")  // se trae por variable de entorno el nombre del bucket S3 y asi no tenerlo hard code

	if !traeParametro{
		return traeParametro
	}

	_ , traeParametro = os.LookupEnv("UrlPrefix")  // UrlPrefix para quitarle el prefijo a la URL y ese prefijo se configura en las variables de entorno

	if !traeParametro{
		return traeParametro
	}

	return traeParametro
}

// Ejemplos punteros
func Punteros() {
    var x int = 10      // Declara una variable int
    var ptr *int = &x   // Declara un puntero a int y asigna la dirección de memoria de x

    fmt.Println("Valor de x:", x)      // Imprime el valor de x
    fmt.Println("Dirección de x:", &x) // Imprime la dirección de memoria de x
    fmt.Println("Valor apuntado por ptr:", *ptr) // Imprime el valor apuntado por ptr (dereferenciación)

    *ptr = 20 // Modifica el valor de x a través de ptr (dereferenciación y asignación)

    fmt.Println("Nuevo valor de x:", x) // Imprime el nuevo valor de x
}