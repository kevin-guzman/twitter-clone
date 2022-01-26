package routers

import (
	"net/http"
	"twitter/bd"
	"twitter/models"
)

func BajaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "EL parámetro id es obligatorio", http.StatusBadRequest)
		return
	}

	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := bd.Borrorelacion(t)
	if err != nil || !status {
		http.Error(w, "Error al borrar la relacion", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
