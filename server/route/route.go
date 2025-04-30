package route

import (
    "net/http"
    "server/controller"
)

func InitRoutes() {
    // Authentification
    http.HandleFunc("/forum/register", controller.Register)
    http.HandleFunc("/forum/login", controller.Login)

    // Threads
    http.HandleFunc("/forum/threads", controller.ThreadsHandler)            // GET, POST
    http.HandleFunc("/forum/thread/", controller.ThreadByIDHandler)         // GET, PUT, DELETE

    // Messages
    http.HandleFunc("/forum/thread/", controller.MessageHandler)            // POST (ajout de message)
    http.HandleFunc("/forum/message/", controller.MessageByIDHandler)       // PUT, DELETE

    // Likes / Dislikes
    http.HandleFunc("/forum/message/", controller.LikeDislikeHandler)       // POST pour like/dislike

    // Recherche & tags
    http.HandleFunc("/forum/search", controller.SearchThreads)
    http.HandleFunc("/forum/threads/tags/", controller.ThreadsByTag)

    // Admin
    http.HandleFunc("/forum/admin/thread/", controller.AdminThreadHandler)  // PUT, DELETE
    http.HandleFunc("/forum/admin/message/", controller.AdminDeleteMessage)
    http.HandleFunc("/forum/admin/ban/", controller.AdminBanUser)

    /* (Bonus) Profil
    http.HandleFunc("/forum/profile/", controller.ProfileHandler)           // GET, PUT

    // (Bonus) Amis
    http.HandleFunc("/forum/friends/", controller.FriendHandler)            // POST demandes/accept
    http.HandleFunc("/forum/threads/private", controller.GetPrivateThreads) // GET

    // Lancer le serveur
    http.ListenAndServe(":8080", nil)*/
}
