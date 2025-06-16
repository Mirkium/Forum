package route

import (
	"net/http"
	controller "server/controller"
)

func InitRoutes() {
	// Accueil
	http.HandleFunc("/", controller.AccueilHandler)
	http.HandleFunc("/forum/home", controller.HomeHandler)

	// Authentification
	http.HandleFunc("/forum/register", controller.Register)
	http.HandleFunc("/forum/login", controller.Login)
	http.HandleFunc("/forum/connect", controller.Connect)

	// Topic
	http.HandleFunc("/forum/topic/get_add", controller.AddTopic_Get)
	http.HandleFunc("/forum/topic/post_add", controller.AddTopic_Post)
	http.HandleFunc("/forum/topic/", controller.TopicHandler)

	// Thread
	http.HandleFunc("/forum/topic/thread/get_add/", controller.AddThread_Get)
	http.HandleFunc("/forum/topic/thread/post_add/", controller.AddThread_Post)

	// Likes
	http.HandleFunc("/forum/thread/like", controller.LikeThreadHandler)

	// Tag
	http.HandleFunc("/forum/tag/get_add", controller.AddTag_get)
	http.HandleFunc("/forum/tag/post_add", controller.AddTag_Post)
}
