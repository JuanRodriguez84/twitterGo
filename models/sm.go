package models

type Secret struct { //  type   para crear una estructura

	Host string `json:"host"` // ``: se consigue con el Alt+96 
	Username string `json:"username"`
	Password string `json:"password"`
	JWTSign string `json:"jwtsign"`
	Database string `json:"database"`

	// si luego en el secreto se crean más campos, se van a tener que incluir dentro de esta estructura por que cuando se decodifique
	// se va a tomar todooo el secreto que viene encriptado y decodificado en 64 bits  y lo va a transformar en cada uno de estos campos
}