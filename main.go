package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Id      string `json: "id"`
	Title   string `json: "title"`
	Desc    string `json: "desc"`
	Content string `json: "content"`
}

var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page!")
	fmt.Println("Endpoint hit: homePage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	fmt.Fprintf(w, "Key: "+key)
	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

// {
//     "Id": "3",
//     "Title": "Newly Created Post",
//     "desc": "The description for my new post",
//     "content": "my articles content"
// }
func createNewArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	// fmt.Fprintf(w, "%+v", string(reqBody))
	var article Article

	json.Unmarshal(reqBody, &article)
	// update global Articles array
	Articles = append(Articles, article)

	json.NewEncoder(w).Encode(article)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	// pars parameters
	vars := mux.Vars(r)
	// id of the article to delete
	id := vars["id"]

	for index, article := range Articles {
		if article.Id == id {
			// update Articles array
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}
	// return all articles to see the result
	json.NewEncoder(w).Encode(Articles)
}

func updateArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	reqBody, _ := ioutil.ReadAll(r.Body)
	var newArticle Article

	json.Unmarshal(reqBody, &newArticle)

	for index, article := range Articles {
		if article.Id == id {
			// update Articles array
			article.Title = newArticle.Title
			article.Desc = newArticle.Desc
			article.Content = newArticle.Content

			Articles[index] = article
		}
	}
	// return all articles to see the result
	json.NewEncoder(w).Encode(Articles)
}

func handleRequests() {

	// http.HandleFunc("/", homePage)
	// http.HandleFunc("/articles", returnAllArticles)
	// log.Fatal(http.ListenAndServe(":8080", nil))

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/articles", returnAllArticles)
	// NOTE: should be defined before other article endpoints
	myRouter.HandleFunc("/articles/{id}", deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/articles/{id}", updateArticle).Methods("PUT")
	myRouter.HandleFunc("/articles/{id}", returnSingleArticle)
	fmt.Println("Server start at localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	Articles = []Article{
		Article{Id: "1", Title: "Hello", Desc: "Article description", Content: "Article content"},
		Article{Id: "2", Title: "Hello 2", Desc: "Article description", Content: "Article content"},
	}

	handleRequests()
}
