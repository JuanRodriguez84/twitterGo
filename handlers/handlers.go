package handlers

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/JuanRodriguez84/twitterGo/models"
	"github.com/JuanRodriguez84/twitterGo/jwt"
)

// primer parametro Context
// Segundo parametro request por que aca vamos a procesar y preguntar por los metodos y rutas, van a venir query strings en Gets, imagenes
// devuelve un modelo de respuesta de api, para esto se crea este modelo
func Manejadores(ctx context.Context, request events.APIGatewayProxyRequest) models.RespuestaAPI {

	// esto envia al cloudwatch un mensaje que dice Voy a procesar el path y el metodo
	// ejemplo dice "Voy a procesar login POST"
	fmt.Println("Voy a procesar" + ctx.Value(models.Key("path")).(string) + " > " + ctx.Value(models.Key("method")).(string))
	var respuesta models.RespuestaAPI

	respuesta.Status = 400 // valor por default

	// en el headr viene un campo llamado validoAuthorization y va a venir con el token, aunque no todos los endpoints van a requerir autenticación, por ejemplo cuando un usuario se registra todavia no hay ningun token
	isOk, statusCode, msg, claim := validoAuthorization(ctx, request)

	if !isOk {
		respuesta.Status = statusCode
		respuesta.Message = msg
		return respuesta
	}

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

	respuesta.Message = "Method invalid" // cada switch tiene un return, y si llega aca es por que llega un path distinto al configurado
	return respuesta
}

func validoAuthorization(ctx context.Context, request events.APIGatewayProxyRequest) (bool, int, string, models.Claim) {
	// existiran 4 path no que nona requerir de autorizacion

	path := ctx.Value(models.Key("path")).(string)

	if path == "registro" || path == "login" || path == "obteneravatar" || path == "obtenerbanner" {
		return true, 200, "", models.Claim{} // models.Claim{} indica que se envia un models.Claim vacio
	}

	token := request.Headers["Authorization"] // en el header viene un campo llamado Authorization y el valor se envia a la variable token
	if len(token) == 0 {                      // cuando no envian el token
		return false, 400, "Token requerido", models.Claim{}
	}

	claim, todoOk, msg, err := jwt.ProcesoToken(token, ctx.Value(models.Key("jwtSign")).(string))

	if !todoOk {
		// cuando el token no es correcto
		if err != nil {
			fmt.Println("Error en el token" + err.Error())
			return false, 401, err.Error(), models.Claim{}
		} else {
			fmt.Println("Error en el token" + msg)
			return false, 401, msg, models.Claim{}
		}
	}

	// cuando esta todo ok
	fmt.Println("Token Ok")
	return true, 200, msg, *claim
}
