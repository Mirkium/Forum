package controller

import (
	"fmt"
	"html/template"
	"net/http"
	models "server/models"
	DB "server/repository"
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
		structHome = *ReloadHome()
		err := temp.ExecuteTemplate(w, "home", structHome)
		if err != nil {
			http.Error(w, "Erreur exécution template : "+err.Error(), http.StatusInternalServerError)
			return
		}
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
	if r.Method == http.MethodGet {
		temp.ExecuteTemplate(w, "addTopic", structReturn)
	}
}



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

	threads := DB.RecupThreadsByTopicID(topicID, userID)
	listTopic := DB.RecupTopics(userID)
	TopicID = topic.Id
	topicPage := models.TopicPage{
		Topic:      topic,
		ListTopic:  listTopic,
		Profil:     UserConnect,
		ListThread: threads,
	}

	err = temp.ExecuteTemplate(w, "topic", topicPage)
	if err != nil {
		fmt.Println(Red, err, Reset)
		http.Error(w, "Erreur serveur", http.StatusInternalServerError)
	}
}

func LikeThreadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		threadIdStr := r.URL.Query().Get("threadId")
		threadId, err := strconv.Atoi(threadIdStr)
		if err != nil {
			http.Error(w, "Invalid thread ID", http.StatusBadRequest)
			return
		}

		userId := UserConnect.Id
		liked := DB.CheckLike(userId, threadId)

		if liked {
			DB.RemoveLike(userId, threadId)
		} else {
			DB.AddLike(userId, threadId)
		}

		http.Redirect(w, r, fmt.Sprintf("/forum/topic?id=%d", UserConnect.Id), http.StatusSeeOther)
	}
}

func AddTag_get(w http.ResponseWriter, r *http.Request) {

}

func AddTag_Post(w http.ResponseWriter, r *http.Request) {

}

func AddThread_Get(w http.ResponseWriter, r *http.Request) {
	topicIDStr := r.URL.Query().Get("id")
	topicID, err := strconv.Atoi(topicIDStr)
	if err != nil {
		http.Error(w, "Invalid topic ID", http.StatusBadRequest)
		return
	}
	TopicThread := models.AddTopic_Thread{
		TopicId: topicID,
	}
	if r.Method == http.MethodGet {
		temp.ExecuteTemplate(w, "addThread", TopicThread)
	}
}


