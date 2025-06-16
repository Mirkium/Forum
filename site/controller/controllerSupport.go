package controller

import (
	"fmt"
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
	structHome.Profil = UserConnect
	structHome.ListTopic = repository.RecupTopics(structHome.Profil.Id)
	structHome.ListTag = repository.RecupTags()
	structHome.Post, err = repository.RecupAllThreadsByDateDesc()
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
