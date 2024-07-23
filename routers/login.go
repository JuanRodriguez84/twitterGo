package routers

import (
	"context"
	"encoding/json"

	"net/http"
	"time"

	"github.com/JuanRodriguez84/twitterGo/bd"
	"github.com/JuanRodriguez84/twitterGo/jwt"
	"github.com/JuanRodriguez84/twitterGo/models"
	"github.com/aws/aws-lambda-go/events"
)

func Login(context context.Context) models.RespuestaAPI {
	var usuario models.Usuario
	var respuesta models.RespuestaAPI

	respuesta.Status = 400

	// capturar el body y convertirlo en strings opr que el body viene en el context pero esta en formato models:key
	body := context.Value(models.Key("body")).(string)

	// el body lo vamos a transformar en el modelo de usuario
	err := json.Unmarshal([]byte(body), &usuario) // que []byte(body) se redirija al puntero de tipo usuario

	if err != nil {
		respuesta.Message = "Usuario y/o Contraseña Inválidos " + err.Error()
		return respuesta
	}

	if len(usuario.Email) == 0 {
		respuesta.Message = "El Email del usuario es requerido"
		return respuesta
	}

	// Aca se pueden hacer más validaciones como validar que venga el @  o el . en el Email

	userData, existe := bd.IntentoLogin(usuario.Email, usuario.Password)
	if !existe {
		respuesta.Message = "Usuario y/o Contraseña Inválidos "
		return respuesta
	}

	// funcion que genera un json web token y devuelve el token
	jwtKey, err := jwt.GeneroJWT(context, userData)

	if err != nil {
		respuesta.Message = "Ocurrió un error al intentar generar el token > " + err.Error()
		return respuesta
	}

	resp := models.RespuestaLogin{ // respuesta
		Token: jwtKey,
	}

	token, err2 := json.Marshal(resp) // El Marshal en este caso por que agarra la estructura de GO y la convierte en un json, por que el front espera un formato json

	if err2 != nil {
		respuesta.Message = "Ocurrió un error al intentar formatear el token a Json> " + err2.Error()
		return respuesta
	}

	// ademas de devolver un token se devuelve en el header como respuesta una cookie para que el usuario tenga el token grabado en una cookie de sus sistema  y cuando vuelva a logearase en nuestro twitter no tenga que buscar un token
	// lo autodetecte, no necesite logearse y que ese token permita ahorrar tiempo para eso es el paquete NET/http
	cookie := &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: time.Now().Add(time.Hour * 24),
	}

	// Para este proyecto existen algunos EndPoint que devuelven una respuesta de API gateway proxy diferente y por eso se construyo el main de la manera que se construyo
	// la siguiente es una de esas
	cookieStinng := cookie.String()

	res := &events.APIGatewayProxyResponse{ //respuesta
		StatusCode: 200,
		Body:       string(token), // convertir a string token que es slice de bytes
		Headers: map[string]string{ // mapa de tipo string y valores de tipo string
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*", // se coloca * por que es necesario para enviar la cookie
			"Set-Cookie":                  cookieStinng,
		},
	}

	// Así se responde con una cookie cuando se requiera

	respuesta.Status = 200
	respuesta.Message = string(token)
	respuesta.CustomResponse = res
	return respuesta
}
