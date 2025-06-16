package main

import (
	"log"
	"net/http"
	//"github.com/gorilla/sessions"
	controller "server/controller"
	repository "server/repository"
	routes "server/route"
)

//var Store = sessions.NewCookieStore([]byte("clé-secrète-super-secrète"))

func main() {

	repository.InitEnv()
	repository.InitDB()
	controller.Init()
	routes.InitRoutes()
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	log.Println("Serveur lancé sur http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Erreur serveur : %v", err)
	}
}
