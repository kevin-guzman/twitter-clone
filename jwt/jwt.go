package jwt

import (
	"time"
	"twitter/models"

	"github.com/dgrijalva/jwt-go"
)

func GeneroJWT(t models.Usuario)(string, error) {
	miClave := []byte("Clave$$dgh872ry")
	payload := jwt.MapClaims{
		"email": t.Email,
		"nombre":t.Nombre,
		"apellidos":t.Apellidos,
		"biografia":t.Biografia,
		"fecha_nacimiento":t.FechaNacimiento,
		"ubicacion":t.Ubicacion,
		"sitioweb":t.SitioWeb,
		"_id":t.ID.Hex(),
		"exp":time.Now().Add(time.Hour*24).Unix(),
	}
	token:= jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)
	if err!=nil{
		return tokenStr, err
	}
	return tokenStr, nil
}
