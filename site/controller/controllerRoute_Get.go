package controller

import (
	"fmt"
	"html/template"
	"net/http"
	models "server/models"
	DB "server/repository"
	repository "server/repository"
	"strconv"
)

var CheckError models.Error
var temp *template.Template

func Init() {
	var err error
	temp, err = template.ParseGlob("templates/*.html")
	if err != nil {
		fmt.Println("Erreur lors du chargement des templates :", err)
	}
}

// ===============================     Accueil      ==============================

func AccueilHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/forum/home", http.StatusSeeOther)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		if temp == nil {
			http.Error(w, "Templates non chargés", http.StatusInternalServerError)
			return
		}
		cookie, err := r.Cookie("user_id")
		if err != nil {
			http.Redirect(w, r, "/forum/connect", http.StatusSeeOther)
			return
		}

		userID, err := strconv.Atoi(cookie.Value)
		if err != nil {
			http.Redirect(w, r, "/forum/connect", http.StatusSeeOther)
			return
		}

		UserConnect, err = repository.GetUserByID(userID)
		if err != nil {
			http.Redirect(w, r, "/forum/connect", http.StatusSeeOther)
			return
		}
		UserConnect.IsConnect = true
		structHome = *ReloadHome()
		temp.ExecuteTemplate(w, "home", structHome)
	}
}

// =============================== Login & Register ==============================

func Connect(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		temp.ExecuteTemplate(w, "connect", CheckError)
	}
}

//=============================================================================================

// ========================================== TOPIC ===========================================

func AddTopic_Get(w http.ResponseWriter, r *http.Request) {
	structReturn := ReloadAddTopic()
	structReturn.Profil = UserConnect
	if r.Method == http.MethodGet {
		temp.ExecuteTemplate(w, "addTopic", structReturn)
	}
}

// ------------------ Handler ------------------

func TopicHandler(w http.ResponseWriter, r *http.Request) {
	topicIDStr := r.URL.Query().Get("id")
	topicID, err := strconv.Atoi(topicIDStr)
	if err != nil {
		http.Error(w, "Invalid topic ID", http.StatusBadRequest)
		return
	}

	// Récupération de l'utilisateur
	userID := UserConnect.Id // On garde le UserConnect déjà chargé
	topic, err := DB.RecupTopicByID(topicID)
	if err != nil {
		http.Error(w, "Topic not found", http.StatusNotFound)
		return
	}
	NewTopic = topic

	threads := DB.RecupThreadsByTopicID(topicID, userID)
	listTopic := DB.RecupTopics(userID)
	TopicID = topic.Id
	topicPage := models.TopicPage{
		Topic:      topic,
		ListTopic:  listTopic,
		Profil:     UserConnect,
		ListThread: threads,
	}
	topicPage.Profil = UserConnect

	err = temp.ExecuteTemplate(w, "topic", topicPage)
	if err != nil {
		fmt.Println(Red, err, Reset)
		http.Error(w, "Erreur serveur", http.StatusInternalServerError)
	}
}

func TopicHandlerByName(w http.ResponseWriter, r *http.Request) {

}

// ========================================== Tag ==========================================

func AddTag_get(w http.ResponseWriter, r *http.Request) {

}

func AddTag_Post(w http.ResponseWriter, r *http.Request) {

}

// ========================================== Thread ==========================================

func AddThread_Get(w http.ResponseWriter, r *http.Request) {
	topicIDStr := r.URL.Query().Get("id")
	topicID, err := strconv.Atoi(topicIDStr)
	if err != nil {
		http.Error(w, "Invalid topic ID", http.StatusBadRequest)
		return
	}

	TopicThread := models.AddTopic_Thread{
		TopicId: topicID,
		Profil:  UserConnect,
	}

	if r.Method == http.MethodGet {
		temp.ExecuteTemplate(w, "addThread", TopicThread)
	}
}

// ========================================== Profil ==========================================

func ProfilHandler(w http.ResponseWriter, r *http.Request) {
	structProfil := models.Profil_Page{
		Profil: UserConnect,
	}
	temp.ExecuteTemplate(w, "Profil", structProfil)
}

func ConsultProfil(w http.ResponseWriter, r *http.Request) {
	ProfilName := r.URL.Query().Get("name")

	OtherUser, erreur := DB.GetUserByUsername(ProfilName)
	if erreur != nil {
		fmt.Println(Red, "Error recup user ID, error : ", erreur, Reset)
		return
	}
	structHome := models.Profil_Page{
		Profil: OtherUser,
	}

	temp.ExecuteTemplate(w, "profil", structHome)
}

// ========================================== Follow ==========================================

func FollowTopic(w http.ResponseWriter, r *http.Request) {
	TopicID := r.URL.Query().Get("id")
	topicID, err := strconv.Atoi(TopicID)
	if err != nil {
		fmt.Println(Red, "Error Invalid Topic ID, error : ", err, Reset)
		return
	}
	follow, Err := DB.IsFollowing(UserConnect.Id, topicID)
	if Err != nil {
		fmt.Println(Red, "Error verif is following, error : ", Err, Reset)
		return
	}

	if follow {
		http.Redirect(w, r, "/forum/topic?id="+TopicID, http.StatusSeeOther)
		return
	} else {
		erreur := DB.FollowTopic(UserConnect.Id, topicID)
		if erreur != nil {
			fmt.Println(Red, "error with following topic, error : ", erreur, Reset)
			return
		}
		http.Redirect(w, r, "/forum/topic?id="+TopicID, http.StatusSeeOther)
		return
	}
}

// ========================================== UnFollow ==========================================

func UnFollowTopic(w http.ResponseWriter, r *http.Request) {
	TopicID := r.URL.Query().Get("id")
	topicID, err := strconv.Atoi(TopicID)
	if err != nil {
		fmt.Println(Red, "Error Invalid Topic ID, error : ", err, Reset)
		return
	}
	erreur := DB.UnfollowTopic(UserConnect.Id, topicID)
	if erreur != nil {
		fmt.Println(Red, "error with unfollowing topic, error : ", erreur, Reset)
		return
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "user_id",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})
	http.Redirect(w, r, "/forum/connect", http.StatusSeeOther)
}
