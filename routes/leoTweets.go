package routers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"twitter/bd"
)

func LeoTweets(w http.ResponseWriter, r *http.Request)  {
	ID:=r.URL.Query().Get("id")
	pagina, err:=strconv.Atoi(r.URL.Query().Get("pagina"))
	if err != nil {
		http.Error(w, "EL parámetro pagina debe ser matyor a cero"+err.Error(), http.StatusBadRequest)
		return
	}
	limite, err:=strconv.Atoi(r.URL.Query().Get("limite"))
	if err != nil {
		log.Fatal(err.Error())
	}
	if len(ID) <1{
		http.Error(w, "EL parámetro id es obligatorio", http.StatusBadRequest)
		return
	}
	if pagina <1{
		http.Error(w, "EL parámetro pagina es obligatorio", http.StatusBadRequest)
		return
	}
	if limite <1{
		limite=20
	}

	respuesta,correcto,err:=bd.LeoTweets(ID, int64(pagina), int64(limite))
	if !correcto{
		http.Error(w, "Eror al leer los tweets", http.StatusBadRequest)
		return
	}
	
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(respuesta)
}