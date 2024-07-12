package jwt

import (
	"errors"
	"strings"

	// esta en la pagina de jwt para go , se debe hacer el respectivo go get en la terminal
	"github.com/JuanRodriguez84/twitterGo/models"
	jwt "github.com/golang-jwt/jwt/v5" // para verificar si el usuario que viene el el token existe
)

// aca se va a hacer todo lo que tiene que ver con procesar el token
// lo que se va a hacer es que viene un token ,  el front envia en uno de los endpoint el token, nosotros lo decodificarlo, verificar de que
// el token no este expirado y que sea correcto
// el token se va a decodificar con la palabra clave

var Email string
var IDUsuario string

// aca se va a porcesar el token, pero tambien se va a crear el token
// La dinamica es, el usuario se logea contra la BD, obtenemos los datos, generamos un token y se lo enviamos al front

// recibe en el primer parametro el token
// recibe en el segundo  parametro la clave para el ejemplo es "MastersdelDesarrollo_grupodeFacebook"  y esta configurado en el secretmanager de Amason

// Devuelve:
// un puntero models.Claim,
// un bool para indicar si el token es valido
// un string para un mensaje
// un error
func ProcesoToken(tk string, JWTSign string) (*models.Claim, bool, string, error) {

	miClave := []byte(JWTSign) // vamos a convertir el JWTSign que es la clave en un slice de bytes  "MastersdelDesarrollo_grupodeFacebook"
	var claims models.Claim    // para trabajar sobre ella y procesar la información

	// se hace un split al token "tk" para quitarle la palabra Bearer
	splitToken := strings.Split(tk, "Bearer")
	// splitToken debe tener un arreglo de 2 posiciones, en la cero esta "Bearer:" y en el uno el token

	if len(splitToken) != 2 {
		return &claims, false, "", errors.New("formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1]) // TrimSpace quita espacios que pueda tener al inicio y al final de uan cadena de caracteres
	// tk toma el valor del token, que esta en la posicon 1 ya que en el 0 esta el valor Bearer

	// ahora vamos a procesar el token

	tkn, err := jwt.ParseWithClaims(
		tk,
		&claims, // puntero a estructura que implementa la interfaz jwt.Claims, donde se almacenarán los claims (datos) extraídos del token JWT despues de parseado.
		func(token *jwt.Token) (interface{}, error) { // func(token *jwt.Token) (interface{}, error) {  es una función anonima  que se utiliza para validar y decodificar la firma del token. ->   *jwt.Token es donde viene el token
			// interface{}  es porque puede contener cualquier tipo de valor, es generico
			return miClave, nil // aca la clave se decodifica
		})

	if err == nil {
		// Aca se va a tener la rutina que chequea contra la BD
	}

	if !tkn.Valid {
		return &claims, false, string(""), errors.New("Token invalido") // esto es lo que va a legar a postman, por ejemplo si el token esta expirado

	}

	return &claims, false, string(""), nil

}
