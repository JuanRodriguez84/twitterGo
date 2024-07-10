package secretmanager

import (
	"encoding/json" // libreria que vienen dentro de GO estandar que permite trabajar con toda la codificación desde y hacia json
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager" // para llamar al servicio secretmanager , dentro de la carpeta github.com/aws/aws-sdk-go-v2/service/ estan todos los servicios de Amazon

	// para utilizar este paquete se debe ingresar por terminal go get  -->   go get github.com/aws/aws-sdk-go-v2/service/secretsmanager
	"github.com/JuanRodriguez84/twitterGo/awsgo"
	"github.com/JuanRodriguez84/twitterGo/models" // el que tiene la estructura donde va a decodificar nuestro secreto
)

// devuelve models.Secret que es donde definimos la estructura
// lo otro que devuelve es un error en caso de halla
func GetSecret(secretname string) (models.Secret, error) { // secretname se recibirá por variable de entorno
	var datosSecret models.Secret
	fmt.Println("> Pido Secreto " + secretname)

	svc := secretsmanager.NewFromConfig(awsgo.Cfg) // el config grabado del inicio de sesión de Amazon
	clave, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretname), // Amazon tiene algunas funciones que los stricgs no los maneja como los string de GO, tienen un formato diferente por eso se debe convertir con aws.String
		// esto ya me traeria el secreto
	})

	if err != nil {
		fmt.Println("Error GetSecret" + err.Error())
		return datosSecret, err
	}

	// convertir el secret decodificado en la estructura json que tenemos creada

	// en clave.SecretString está todo el conjunto de datos, los 5 campos del API decodificados y eso lo voy a convertir en una estructura
	// en el segundo parametro se indica donde se debe guardar  en este caso en &datosSecret; como puntero para indicarle, que guarde el resultado en la direccion de memoria donde esta la variable datosSecret, de esta manera es más rapido ya que no hay que
	// indicarle a GO la referencia-directa, es decir sin el puntero (datosSecret) ya que el motor de Go tiene que ir a buscar datosSecret, buscar cual es la dirección de memoria hacer esta conversión es más demorado
	// siempre es mas rapido programar con punteros
	json.Unmarshal([]byte(*clave.SecretString), &datosSecret) // Unmarshal función que recibe un slice de bits no un string, por eso se debe convertir en slice de bits
	fmt.Println("> Lectura de secret OK", secretname)
	return datosSecret, nil
}
