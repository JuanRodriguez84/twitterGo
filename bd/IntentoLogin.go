package bd

import (
	"github.com/JuanRodriguez84/twitterGo/models"
	"golang.org/x/crypto/bcrypt"
)

func IntentoLogin(email string, password string) (models.Usuario, bool){
	usuario, encontrado, _ := ChequeoYaExisteUsuario(email) // aca se conecta a la BD de mongo y trae los datos

	if !encontrado{
		return usuario, false
	}

	passwordBytes := []byte(password) // va a ser un slice de byte   que es la que digita el usuario en el formulario
	passwordDB := []byte(usuario.Password) // es la que esta en la BD

	// validar si la password encriptada de BD coincide con la digitada en el formulario , compara el hash
	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
	if err != nil{
		// contraseña invalida
		return usuario, false
	}

	return usuario, true
}