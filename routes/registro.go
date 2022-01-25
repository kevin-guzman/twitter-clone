package routers

import (
	"encoding/json"
	"net/http"
	"twitter/bd"
	"twitter/models"
)

func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en datos de resgistro "+err.Error(), 400)
		return
	}
	if len(t.Email)==0{
		http.Error(w, "El email es requerido", 400)
		return
	}
	if len(t.Password)<6{
		http.Error(w, "La contraseña debe ser de al menos 6 caracteres", 400)
		return
	}
	_,encontrado,_:=bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado{
		http.Error(w, "Ya existe un suario registrado con este email", 400)
		return
	}
	_,status,err:=bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Error al registrar el usuario "+err.Error(), 400)
		return
	}
	if !status {
		http.Error(w, "Error al registrar el usuario ", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
