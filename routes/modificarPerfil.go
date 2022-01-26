package routers

import (
	"encoding/json"
	"net/http"
	"twitter/bd"
	"twitter/models"
)

func ModificarPerfil(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}
	
	status, err := bd.ModificoRegistro(t, IDUsuario)
	if err != nil {
		http.Error(w, "Error al modificar el registro, reintente nuevamente", http.StatusInternalServerError)
		return
	}
	if !status {
		http.Error(w, "No se ha podido modificar el registro del usuario", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
