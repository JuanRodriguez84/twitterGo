package models

import "github.com/aws/aws-lambda-go/events"

type RespuestaAPI struct {
	Status int   // devuelve un status 200, 400, 500
	Message string  // mensaje
	CustomResponse *events.APIGatewayProxyResponse // esto es lo que devuelve la lambda , es lo que devuelve la función EjecutoLambda de main.go
}  