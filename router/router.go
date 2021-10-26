package router

import (
	"fmt"
	"log"
	"net/http"

	Article "vlad/rest/controllers/article"
	Home "vlad/rest/controllers/home"

	"github.com/gorilla/mux"
)

func HandleRequests() {

	// http.HandleFunc("/", homePage)
	// http.HandleFunc("/articles", returnAllArticles)
	// log.Fatal(http.ListenAndServe(":8080", nil))

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", Home.HomePage)
	myRouter.HandleFunc("/articles", Article.CreateNewArticle).Methods("POST")
	myRouter.HandleFunc("/articles", Article.ReturnAllArticles)
	// NOTE: should be defined before other article endpoints
	myRouter.HandleFunc("/articles/{id}", Article.DeleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/articles/{id}", Article.UpdateArticle).Methods("PUT")
	myRouter.HandleFunc("/articles/{id}", Article.ReturnSingleArticle)
	fmt.Println("Server start at localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
