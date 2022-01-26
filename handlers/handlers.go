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
	router.HandleFunc("/registro", middlew.ChequeoDB(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.ChequeoDB(routers.Login)).Methods("POST")

	router.HandleFunc("/verperfil", middlew.ChequeoDB(middlew.ValidoJWT(routers.VerPerfil))).Methods("GET")
	router.HandleFunc("/modificarPerfil", middlew.ChequeoDB(middlew.ValidoJWT(routers.ModificarPerfil))).Methods("PUT")
	router.HandleFunc("/listaUsuarios", middlew.ChequeoDB(middlew.ValidoJWT(routers.ListaUsuarios))).Methods("GET")

	router.HandleFunc("/tweet", middlew.ChequeoDB(middlew.ValidoJWT(routers.GraboTweet))).Methods("POST")
	router.HandleFunc("/leoTweets", middlew.ChequeoDB(middlew.ValidoJWT(routers.LeoTweets))).Methods("GET")
	router.HandleFunc("/eliminarTweet", middlew.ChequeoDB(middlew.ValidoJWT(routers.EliminarTweet))).Methods("DELETE")
	router.HandleFunc("/leoTweetsSeguidores", middlew.ChequeoDB(middlew.ValidoJWT(routers.LeoTweetsSeguidores))).Methods("GET")

	router.HandleFunc("/subirAvatar", middlew.ChequeoDB(middlew.ValidoJWT(routers.SubirAvatar))).Methods("POST")
	router.HandleFunc("/obtenerAvatar", middlew.ChequeoDB(routers.ObtenerAvatar)).Methods("GET")

	router.HandleFunc("/subirBanner", middlew.ChequeoDB(middlew.ValidoJWT(routers.SubirBanner))).Methods("POST")
	router.HandleFunc("/obtenerBanner", middlew.ChequeoDB(routers.ObtenerBanner)).Methods("GET")

	router.HandleFunc("/altaRelacion", middlew.ChequeoDB(middlew.ValidoJWT(routers.AltaRelacion))).Methods("POST")
	router.HandleFunc("/bajaRelacion", middlew.ChequeoDB(middlew.ValidoJWT(routers.BajaRelacion))).Methods("DELETE")


	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
