package controller

import (
	"fmt"
	"net/http"
	"text/template"
)

var temp, err = template.ParseGlob("../../site/templates/*.html")

//=================================Login & Register================================

func Login(w http.ResponseWriter, r *http.Request) {

	errLogin := temp.ExecuteTemplate(w, "login", nil)
	if errLogin != nil {
		fmt.Println("Error : ", Red, errLogin, Reset)
	}
}

func Register(w http.ResponseWriter, r *http.Request) {

	errRegister := temp.ExecuteTemplate(w, "register", nil)
	if errRegister != nil {
		fmt.Println("Error : ", Red, errRegister, Reset)
	}
}

//================================================================================

//===============================Threads=========================================

func ThreadsHandler(w http.ResponseWriter, r *http.Request) {

}

func ThreadByIDHandler(w http.ResponseWriter, r *http.Request) {

}

//===============================================================================

//===================================Messages====================================

func MessageHandler(w http.ResponseWriter, r *http.Request) {

}

func MessageByIDHandler(w http.ResponseWriter, r *http.Request) {

}

//===============================================================================

//================================Like / DisLike=================================

func LikeDislikeHandler(w http.ResponseWriter, r *http.Request) {

}

//===============================================================================

//===============================Search & Tags===================================

func SearchThreads(w http.ResponseWriter, r *http.Request) {

}

func ThreadsByTag(w http.ResponseWriter, r *http.Request) {

}

//===============================================================================

//===================================Admin=======================================

func AdminThreadHandler(w http.ResponseWriter, r *http.Request) {

}

func AdminDeleteMessage(w http.ResponseWriter, r *http.Request) {

}

func AdminBanUser(w http.ResponseWriter, r *http.Request) {

}

//===============================================================================
