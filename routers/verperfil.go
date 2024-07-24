package routers

import (
	"encoding/json"
	"fmt"

	"github.com/JuanRodriguez84/twitterGo/bd"
	"github.com/JuanRodriguez84/twitterGo/models"
	"github.com/aws/aws-lambda-go/events"
)

// Todos las funciones de routers devuelven RespuestaAPI
func VerPerfil(request events.APIGatewayProxyRequest) models.RespuestaAPI {
	var respuesta models.RespuestaAPI
	respuesta.Status = 400

	fmt.Println("Entré en VerPerfil")

	// extraer de las query strings que vienen en la Url el id del usuario al que se le va a ver el perfil
	// cuando en la url viene el signo ?   el nombre de una variable   =   valor
	ID := request.QueryStringParameters["Id"]

	if len(ID) < 1 {
		respuesta.Message = "El párametro ID es obligatorio"
		return respuesta
	}

	// ejecutar una funcionde bd para buscar
	perfil, err := bd.BuscoPerfil(ID)
	if err != nil {
		respuesta.Message = "Ocurrió un error al intentar buscar el registro " + err.Error()
		return respuesta
	}

	// Se hace conversiona Json
	respJson, err := json.Marshal(perfil) // Marshal recibe un slice de byte y devulve un json

	if err != nil {
		respuesta.Status = 500 // cuando el error es interno del servidor
		respuesta.Message = "Error al formatear los datos de los usuarios como Json " + err.Error()
		return respuesta
	}

	respuesta.Status = 200
	respuesta.Message = string(respJson)
	return respuesta
}
