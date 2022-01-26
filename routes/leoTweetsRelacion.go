package routers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"twitter/bd"
)

func LeoTweetsSeguidores(w http.ResponseWriter, r *http.Request)  {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "El paŕametro page es obligatorio", http.StatusBadRequest)
		return
	}
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 0
	}

	respuesta,err:=bd.LeoTweetsSeguidores(IDUsuario, page, limit)
	if err != nil {
		http.Error(w, "Ocurrió un error al crear el usuario "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(respuesta)
}