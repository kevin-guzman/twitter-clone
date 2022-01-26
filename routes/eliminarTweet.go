package routers

import (
	"net/http"
	"twitter/bd"
)

func EliminarTweet(w http.ResponseWriter, r *http.Request)  {
	ID:=r.URL.Query().Get("id")
	if len(ID) <1{
		http.Error(w, "EL parámetro id es obligatorio", http.StatusBadRequest)
		return
	}
	
	err:=bd.BorroTweet(ID,IDUsuario)
	if err != nil {
		http.Error(w, "Error al borrar el tweet"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}