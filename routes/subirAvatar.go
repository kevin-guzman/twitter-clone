package routers

import (
	"io"
	"net/http"
	"os"
	"strings"
	"twitter/bd"
	"twitter/models"
)

func SubirAvatar(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("avatar")
	if err != nil {
		http.Error(w, "Error no se ha encontrado el archivo de subida "+err.Error(), http.StatusInternalServerError)
		return
	}
	
	var extension = strings.Split(handler.Filename, ".")[1]
	var archivo = "uploads/avatars/" + IDUsuario + "." + extension

	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error al subir la imagen "+err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error al copiar la imagen "+err.Error(), http.StatusInternalServerError)
		return
	}

	var usuario models.Usuario
	var status bool
	usuario.Avatar = IDUsuario + "." + extension
	status, err = bd.ModificoRegistro(usuario, IDUsuario)
	if err != nil || !status {
		http.Error(w, "Error al guardar el avatar en la db "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
