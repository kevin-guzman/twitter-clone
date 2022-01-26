package routers

import (
	"encoding/json"
	"net/http"
	"time"
	"twitter/bd"
	"twitter/models"
)

func GraboTweet(w http.ResponseWriter, r *http.Request)  {
	var mensaje models.Tweet
	err:= json.NewDecoder(r.Body).Decode(&mensaje)
	if err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	registro:=models.GraboTweet{
		UserID: IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha: time.Now(),
	}
	_,status,err:=bd.InsertoTweet(registro)
	if err != nil {
		http.Error(w, "Error al crear el tweet"+err.Error(), http.StatusInternalServerError)
		return
	}
	if !status {
		http.Error(w, "Error al crear el tweet, no ha sido posible crearlo", http.StatusInternalServerError)
		return
	}
	
	w.WriteHeader(http.StatusCreated)
}