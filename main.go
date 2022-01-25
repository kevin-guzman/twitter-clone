package main

import (
	"log"
	"twitter/bd"
	"twitter/handlers"
)

func main(){
	if !bd.CheckConnection(){
		log.Fatal("No se pudo conectar a la db")
	}
	handlers.Manejadores()
}