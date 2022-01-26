package routers

import (
	"errors"
	"strings"
	"twitter/bd"
	"twitter/models"

	"github.com/dgrijalva/jwt-go"
)

var Email, IDUsuario string

func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("Clave$$dgh872ry")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer ")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}
	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(t *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if !tkn.Valid || claims.ID.IsZero() {
		return claims, false, string(""), errors.New("token invalido")
	}
	if err == nil {
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, encontrado, IDUsuario, nil
	}

	return claims, false, string(""), err
}
