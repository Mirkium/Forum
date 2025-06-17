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
	http.HandleFunc("/forum/topic/unfollow/", controller.UnFollowTopic)
	http.HandleFunc("/forum/topic/follow/", controller.FollowTopic)
	http.HandleFunc("/forum/topic/name/", controller.TopicHandlerByName)

	// Thread
	http.HandleFunc("/forum/topic/thread/get_add/", controller.AddThread_Get)
	http.HandleFunc("/forum/topic/thread/post_add", controller.AddThread_Post)

	// Likes
	http.HandleFunc("/forum/thread/like/", controller.LikeThreadHandler)
	http.HandleFunc("/forum/thread/unlike/", controller.LikeThreadHandler)
	http.HandleFunc("/forum/topic/thread/like/", controller.LikeThreadTopic)

	// Tag
	http.HandleFunc("/forum/tag/get_add", controller.AddTag_get)
	http.HandleFunc("/forum/tag/post_add", controller.AddTag_Post)

	// User
	http.HandleFunc("/forum/user", controller.ProfilHandler)
	http.HandleFunc("/forum/user/", controller.ConsultProfil)
	http.HandleFunc("/forum/user/follow", controller.FollowUser)
	http.HandleFunc("/logout", controller.Logout)
}
