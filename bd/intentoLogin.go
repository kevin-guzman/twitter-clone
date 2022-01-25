package bd

import (
	"twitter/models"

	"golang.org/x/crypto/bcrypt"
)

func IntentoLogin(email, pass string)(models.Usuario, bool){
	usu, encontrado,_ := ChequeoYaExisteUsuario(email)
	if !encontrado{
		return usu, false
	}
	passwordBytes:=[]byte(pass)
	passwordBD:= []byte(usu.Password)
	err:=bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
	if err!=nil{
		return usu, false
	}
	return usu, true
}