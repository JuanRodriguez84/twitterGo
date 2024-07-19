package bd

import "golang.org/x/crypto/bcrypt"// bcrypt libreria para encriptar texto y devuelve un slice de bytes, se debe descargar con el go get

func EncriptarPassword(password string)(string, error){

	costo := 8 // el costo es, lo que hace bcrypt de manera muy inteligente es encriptar la password 8 veces, entre mas alto el costo mas segura es la contraseña pero utiliza más recursos
	// con 6 es aceptable y con 8 es más seguro
	
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), costo)  // convertor la password que llega como parametro en un slice de bytes

	if err != nil {
		return err.Error(), err
	}

	return string(bytes), nil // convertir un slice de bytes en string 
}