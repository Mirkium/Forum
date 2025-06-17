package controller

import (
	"fmt"
	"net/http"
	"regexp"
	DB "server/repository"
	"strconv"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Println(Red, "Erreur d'input (Login)", Reset)
		CheckError.ValueError = true
		CheckError.IsLogin = true
		http.Redirect(w, r, "/forum/connect", http.StatusSeeOther)
		return
	}

	username := r.FormValue("Username")
	password := r.FormValue("password")

	userRepo := &DB.UserRepository{}
	user, err := userRepo.VerifUser(username, password)
	if err != nil {
		fmt.Println(Red, "Error with recup in login, error : ", err, Reset)
		CheckError.ValueError = true
		CheckError.IsLogin = true
		http.Redirect(w, r, "/forum/connect", http.StatusSeeOther)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "user_id",
		Value:    strconv.Itoa(user.Id),
		Path:     "/",
		HttpOnly: true,
		MaxAge:   3600,
	})
	UserConnect = *user
	UserConnect.IsConnect = true
	http.Redirect(w, r, "/forum/home", http.StatusSeeOther)
}

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("Username")
	password := r.FormValue("password")

	validUsername, _ := regexp.MatchString("^[a-zA-Z-]{1,64}$", username)
	if !validUsername {
		fmt.Println(Red, "Error with Username", Reset)
		CheckError.ValueError = true
		CheckError.IsLogin = false
		http.Redirect(w, r, "/forum/connect", http.StatusSeeOther)
		return
	}

	_, err := DB.AddUser(username, password)
	if err != nil {
		fmt.Println(Red, "Erreur register :", err, Reset)
		CheckError.ValueError = true
		CheckError.IsLogin = false
		http.Redirect(w, r, "/forum/connect", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/forum/connect", http.StatusSeeOther)
}

func AddTopic_Post(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		http.Redirect(w, r, "/forum/topic/get_add", http.StatusSeeOther)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Erreur lecture formulaire", http.StatusBadRequest)
		return
	}

	topic_name := r.FormValue("name")
	categoryIDStr := r.FormValue("category_id")

	categoryID, err := strconv.Atoi(categoryIDStr)
	if err != nil {
		http.Error(w, "ID de catégorie invalide", http.StatusBadRequest)
		return
	}

	_, err = DB.AddTopic(topic_name, categoryID)
	if err != nil {
		fmt.Println(Red, "Error with add topic, error : ", err, Reset)
		http.Redirect(w, r, "/forum/topic/get_add", http.StatusSeeOther)
		return
	}
	http.Redirect(w, r, "/forum/home", http.StatusSeeOther)
}

func AddThread_Post(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		fmt.Println(Red, "error method post, type method : ", r.Method, Reset)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Erreur lecture formulaire", http.StatusBadRequest)
		return
	}

	NewThread.Title = r.FormValue("title")
	NewThread.Content = r.FormValue("content")
	NewThread.Description = r.FormValue("description")

	threadId, err := DB.AddThread(NewTopic.Id, NewThread.Title, NewThread.Content, NewThread.Description, UserConnect.Id)
	if err != nil {
		http.Error(w, "Erreur lors de l'ajout du thread", http.StatusInternalServerError)
		fmt.Println(Red, "error : ", err, Reset)
		return
	}
	NewThread.Id = int(threadId)

	http.Redirect(w, r, "/forum/topic?id="+strconv.Itoa(NewTopic.Id), http.StatusSeeOther)

}



func FollowUser(w http.ResponseWriter, r *http.Request) {

}
