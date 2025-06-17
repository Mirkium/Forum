package controller

import (
	"fmt"
	"html/template"
	"log"
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

	user, err := repository.GetUserByID(userID)
	if err != nil {
		http.Redirect(w, r, "/forum/connect", http.StatusSeeOther)
		return
	}

	user.IsConnect = true
	UserConnect = user
	structHome.Profil = UserConnect
	structHome = *ReloadHome()

	err = temp.ExecuteTemplate(w, "home", structHome)
	if err != nil {
		http.Error(w, "Erreur exécution template : "+err.Error(), http.StatusInternalServerError)
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
	topicName := r.URL.Query().Get("name")
	if topicName == "" {
		http.Error(w, "Missing topic name", http.StatusBadRequest)
		return
	}

	// Récupération du topic par son nom
	topic, err := DB.RecupTopicByName(topicName)
	if err != nil {
		http.Error(w, "Topic not found", http.StatusNotFound)
		return
	}
	NewTopic = topic

	userID := UserConnect.Id // Utilisateur connecté déjà chargé
	threads := DB.RecupThreadsByTopicID(topic.Id, userID)
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

	user, err := repository.GetUserByID(userID)
	if err != nil {
		http.Redirect(w, r, "/forum/connect", http.StatusSeeOther)
		return
	}

	user.IsConnect = true
	UserConnect = user

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

	user, err := repository.GetUserByID(userID)
	if err != nil {
		http.Redirect(w, r, "/forum/connect", http.StatusSeeOther)
		return
	}

	user.IsConnect = true
	user.Post, err = DB.RecupThreadsByUserID(userID)
	user.ThreadLike, err = DB.RecupThreadsLikedByUser(userID)
	UserConnect = user
	ListTopic := DB.RecupTopics(userID)
	structProfil := models.Profil_Page{
		Profil:    UserConnect,
		ListTopic: ListTopic,
	}
	err = temp.ExecuteTemplate(w, "Profil", structProfil)
	if err != nil {
		http.Error(w, "Erreur exécution template : "+err.Error(), http.StatusInternalServerError)
		return
	}
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

func LikeThreadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		ThreadID := r.URL.Query().Get("id")
		threadID, err := strconv.Atoi(ThreadID)
		if err != nil {
			http.Error(w, "Invalid thread ID", http.StatusBadRequest)
			return
		}
		userId := UserConnect.Id
		thread, err := DB.RecupThreadByID(threadID)
		if err != nil {
			log.Println("Erreur :", err)
		} else {
			fmt.Println("Thread :", thread.Title)
		}
		exists, isLike, erreur := DB.CheckLike(userId, threadID)
		if erreur != nil {
			log.Println("Erreur CheckLike:", err)
		} else if exists && isLike {
			DB.RemoveLike(userId, threadID)
		} else {
			DB.AddLike(userId, threadID)
		}

		http.Redirect(w, r, "/forum/home", http.StatusSeeOther)
	}
}

func LikeThreadTopic(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		ThreadID := r.URL.Query().Get("id")
		threadID, err := strconv.Atoi(ThreadID)
		if err != nil {
			http.Error(w, "Invalid thread ID", http.StatusBadRequest)
			return
		}
		userId := UserConnect.Id
		thread, err := DB.RecupThreadByID(threadID)
		if err != nil {
			log.Println("Erreur :", err)
		} else {
			fmt.Println("Thread :", thread.Title)
		}
		exists, isLike, erreur := DB.CheckLike(userId, threadID)
		if erreur != nil {
			log.Println("Erreur CheckLike:", err)
		} else if exists && isLike {
			if thread.NbDisLike > 0 || thread.NbLike > 0 {
				DB.RemoveLike(userId, threadID)
			}
		} else {
			DB.AddLike(userId, threadID)
		}
		topicID, err := DB.RecupTopicByThreadID(threadID)
		if err != nil {
			fmt.Println(Red, "Error recup topic, error : ", err, Reset)
			return
		}
		http.Redirect(w, r, "/forum/topic?id="+strconv.Itoa(topicID.Id), http.StatusSeeOther)
	}
}
