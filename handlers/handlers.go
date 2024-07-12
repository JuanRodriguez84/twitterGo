package handlers

import (
	"context"
	"fmt"
	"github.com/JuanRodriguez84/twitterGo/models"
	"github.com/aws/aws-lambda-go/events"
)

// primer parametro Context
// Segundo parametro request por que aca vamos a procesar y preguntar por los metodos y rutas, van a venir query strings en Gets, imagenes
// devuelve un modelo de respuesta de api, para esto se crea este modelo
func Manejadores(ctx context.Context, request events.APIGatewayProxyRequest) models.RespuestaAPI{
	
	// esto envia al cloudwatch un mensaje que dice Voy a procesar el path y el metodo
	// ejemplo dice "Voy a procesar login POST"
	fmt.Println("Voy a procesar" + ctx.Value(models.Key("path")).(string) + " > " + ctx.Value(models.Key("method")).(string))
	var respuesta models.RespuestaAPI

	respuesta.Status = 400 // valor por default

	switch ctx.Value(models.Key("method")).(string) {
	case "POST":
		switch ctx.Value(models.Key("path")).(string) {
			
		}

	case "GET":
		switch ctx.Value(models.Key("path")).(string) {
			
		}

	case "PUT":
		switch ctx.Value(models.Key("path")).(string) {
			
		}

	case "DELETE":
		switch ctx.Value(models.Key("path")).(string) {
			
		}

	}
  
	respuesta.Message = "Method invalid"  // cada switch tiene un return, y si llega aca es por que llega un path distinto al configurado
	return respuesta
}