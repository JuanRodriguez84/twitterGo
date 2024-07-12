package routers

import (
	"context"
	"encoding/json" // por que se va a trabajar con formato json
	"fmt"

	"github.com/JuanRodriguez84/twitterGo/bd"
	"github.com/JuanRodriguez84/twitterGo/models"
)

// se recibe el context por que dentro de context viene el body
func Registro(ctx context.Context) models.RespuestaAPI {
	var usuario models.Usuario
	var respuesta models.RespuestaAPI

	respuesta.Status = 400

	fmt.Println("Entrando a Registro")

	//extraer el body del context

	body := ctx.Value(models.Key("body")).(string)

	// La función json.Unmarshal en Go se utiliza para convertir datos JSON (body en este caso) en un tipo de dato Go (usuario en este caso) específico.
	err := json.Unmarshal([]byte(body), &usuario) // leer el body, luedo desestructurarlo en formato json y todos los valores los convierte al modelo usuario

	if err != nil {
		respuesta.Message = err.Error()
		fmt.Println("Error Unmarshal de Registro" + err.Error())
		return respuesta
	}

	if len(usuario.Email) == 0 { // no llego email
		respuesta.Message = "Debe especificar el Email"
		fmt.Println(respuesta.Message)
		return respuesta
	}

	if len(usuario.Password) < 6 { // debe tener al menos 6 caracteres
		respuesta.Message = "Debe especificar una contraseña de al menos 6 caracteres"
		fmt.Println(respuesta.Message)
		return respuesta
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(usuario.Email) // cuando viene un mail y ya existe no se debe permitir

	if encontrado {
		respuesta.Message = "Ya existe un usuario registrado con este Email"
		fmt.Println(respuesta.Message)
		return respuesta
	}

	// si llega aca el email no existe en la coleccion de mongo
	_, status, err := bd.InsertoRegistro(usuario)

	if err != nil {
		respuesta.Message = "Ocurrrio un error al intentar grabar el usuario" + err.Error()
		fmt.Println(respuesta.Message)
		return respuesta
	}

	if !status {
		respuesta.Message = "No se ha logrado insertar el registro del usuario"
		fmt.Println(respuesta.Message)
		return respuesta
	}

	respuesta.Status = 200
	respuesta.Message = "Registro OK"
	fmt.Println(respuesta.Message)
	return respuesta

}
