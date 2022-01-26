package routers

import (
	"net/http"
	"twitter/bd"
	"twitter/models"
)

func AltaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "EL parámetro id es obligatorio", http.StatusBadRequest)
		return
	}

	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := bd.InsertoRelacion(t)
	if err != nil || !status {
		http.Error(w, "Error al crear la relacion", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
