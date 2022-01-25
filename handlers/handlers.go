package handlers

import (
	"log"
	"net/http"
	"os"
	"twitter/middlew"
	routers "twitter/routes"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Manejadores() {
	router := mux.NewRouter()
	router.HandleFunc("/registro", middlew.ChaqueoDB(routers.Registro) ).Methods("POST")
	PORT := os.Getenv("POTR")
	if PORT == "" {
		PORT = "8080"
	}
	handler:=cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
