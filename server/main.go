package main

import (
	"log"
	"net/http"
	routes "server/route"
)

func main() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("../site/assets"))))
	routes.InitRoutes()
	log.Println("Le serveur écoute sur le port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Erreur du serveur : ", err)
	}
}
