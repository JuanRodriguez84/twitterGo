package jwt

import(
	"context"
	"time"

	"github.com/JuanRodriguez84/twitterGo/models"
	jwt "github.com/golang-jwt/jwt/v5" // para verificar si el usuario que viene el el token existe
)

func GeneroJWT(contexto context.Context, usuario models.Usuario) (string, error){

	jwtSign := contexto.Value(models.Key("jwtSign")).(string)

	miClave:=[]byte(jwtSign)  // convertir jwtSign a slice de byte

	//jwt : compuesto de 3 partes  
		// 1 Es el header donde va la información de que tipo de encriptación tiene, etc
		// 2 payload donde estan todos los datos
		// 3 firma

	// Despues de leer en Bd cuando busca el login trae todo en la estructura usuario con todos los datos

	payload:= jwt.MapClaims{  // MapClaims : Mapa de todo lo requerido 
		"email": usuario.Email,
		"nombre": usuario.Nombre,
		"apellidos": usuario.Apellidos,
		"fecha_nacimiento": usuario.FechaNacimiento,
		"biografia": usuario.Biografia,
		"ubicacion": usuario.Ubicacion,
		"sitioweb": usuario.SitioWeb,
		"_id" : usuario.ID.Hex(), // como ID es de tipo primitivo en la estructura Usuario se convierte a string a travez de la funcion Hex
		"exp": time.Now().Add(time.Hour*24).Unix(), // función para la expiracion y va a durar un dia, por esto se multiplica por 24, se transforma aunique por que la expiración de JWT se grana en formato unix (SO)
	}

	// aca se crea el token
	// SigningMethodHS256 valor utilizado para decodificar nuestro token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload) // como el payload se define a mano va como NewWithClaims
	
	// crear el token definitivo
	tokenStr, err := token.SignedString(miClave) // SignedString para indicarle la palabra clave, por que si no no se va a saber como decodificarlo

	if err != nil{
		return tokenStr, err
	}

	return tokenStr, nil
}