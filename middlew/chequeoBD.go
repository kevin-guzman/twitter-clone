package middlew

import (
	"net/http"
	"twitter/bd"
)

func ChequeoDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !bd.CheckConnection() {
			http.Error(w, "Conexión perdida con al db", 500)
			return
		}
		next.ServeHTTP(w,r)
	}
}
