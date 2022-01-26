package routers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"twitter/bd"
)

func ListaUsuarios(w http.ResponseWriter, r *http.Request) {
	typeUser := r.URL.Query().Get("type")
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "El paŕametro page es obligatorio", http.StatusBadRequest)
		return
	}
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 20
	}
	search := r.URL.Query().Get("search")

	result, status, err := bd.LeoUsuariosTodos(IDUsuario, search, typeUser, int64(page), int64(limit))
	if !status || err != nil {
		http.Error(w, "Error al leer usuarios "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(result)
}
