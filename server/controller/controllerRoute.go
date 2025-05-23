package controller

import (
	"fmt"
	"net/http"
	"text/template"
)

var temp, _ = template.ParseGlob("site/templates/*.html")

// =============================== Login & Register ==============================

func Connect(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		temp.ExecuteTemplate(w, "connect", nil)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		fmt.Println(Red, "Erreur d'input", Reset)
		return
	}
}

func Register(w http.ResponseWriter, r *http.Request) {

}

// =============================== Threads ==============================

func ThreadsHandler(w http.ResponseWriter, r *http.Request) {

}

func ThreadByIDHandler(w http.ResponseWriter, r *http.Request) {

}

// =============================== Messages ==============================

func MessageHandler(w http.ResponseWriter, r *http.Request) {

}

func MessageByIDHandler(w http.ResponseWriter, r *http.Request) {

}

// =============================== Likes / Dislikes ==============================

func LikeDislikeHandler(w http.ResponseWriter, r *http.Request) {

}

// =============================== Search & Tags ==============================

func SearchThreads(w http.ResponseWriter, r *http.Request) {

}

func ThreadsByTag(w http.ResponseWriter, r *http.Request) {

}

// =============================== Admin ==============================

func AdminThreadHandler(w http.ResponseWriter, r *http.Request) {

}

func AdminDeleteMessage(w http.ResponseWriter, r *http.Request) {

}

func AdminBanUser(w http.ResponseWriter, r *http.Request) {

}

// =============================== Profils ==============================

func ProfileHandler(w http.ResponseWriter, r *http.Request) {

}

// =============================== Amis ==============================

func FriendHandler(w http.ResponseWriter, r *http.Request) {

}

func GetPrivateThreads(w http.ResponseWriter, r *http.Request) {

}
