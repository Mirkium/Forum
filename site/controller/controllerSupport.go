package controller

import (
	"fmt"
	"log"
	models "server/models"
	repository "server/repository"
)

const (
	Red   = "\033[91m"
	Reset = "\033[0m"
)

var UserConnect models.User
var UserList []models.User

var NewTopic models.Topic

var structHome models.Home

var PageTopic models.TopicPage
var TopicID int

var NewThread models.Thread

var structAddTopic models.AddTopic_tag

func ReloadHome() *models.Home {
	var err error
	structHome.ListTopic = repository.RecupTopics(structHome.Profil.Id)
	structHome.ListTag = repository.RecupTags()
	structHome.Post, err = repository.RecupAllThreadsByDateDesc()
	fmt.Println(structHome.Post[0].NbLike, " ", structHome.Post[0].NbDisLike)
	for k := 0; k < len(structHome.Post); k++ {
		exists, isLike, erreur := repository.CheckLike(structHome.Profil.Id, structHome.Post[k].Id)
		if erreur != nil {
			log.Println("Erreur CheckLike:", err)
		} else if exists && isLike {
			structHome.Post[k].IsLike = true
		} else {
			structHome.Post[k].IsLike = false
		}
	}
	if err != nil {
		fmt.Println(Red, "Error with recup post, error : ", err, Reset)
	}
	return &structHome
}

func SubscribeCurrentUserToTopic(topicID int) {
	err := repository.SubscribeToTopic(structHome.Profil.Id, topicID)
	if err != nil {
		fmt.Println(Red, "Erreur abonnement : ", err.Error(), Reset)
	} else {
		structHome = *ReloadHome()
	}
}

func ReloadAddTopic() *models.AddTopic_tag {
	structAddTopic.ListTag = repository.RecupTags()
	return &structAddTopic
}
