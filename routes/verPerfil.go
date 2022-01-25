package routers

import (
	"encoding/json"
	"net/http"
	"twitter/bd"
)

func VerPerfil(w http.ResponseWriter, r *http.Request){
	ID:=r.URL.Query().Get("id")
	if len(ID)<1{
		http.Error(w, "Debe enviar un id", http.StatusBadRequest)
		return
	}
	perfil,err:= bd.BuscoPerfil(ID)
	if err != nil {
		http.Error(w, "Ocurrio un error al buscar el registro "+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(perfil) 
}