package middlew

import (
	"net/http"
	routers "twitter/routes"
)

func ValidoJWT(next http.HandlerFunc)http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		_,_,_,err:=routers.ProcesoToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "No autorizado"+err.Error(),401)
			return
		}
		next.ServeHTTP(w,r)
	}
}