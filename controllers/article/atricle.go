package article

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Id      string `json: "id"`
	Title   string `json: "title"`
	Desc    string `json: "desc"`
	Content string `json: "content"`
}

// var Articles []Article

var Articles = []Article{
	{Id: "1", Title: "Hello", Desc: "Article description", Content: "Article content"},
	{Id: "2", Title: "Hello 2", Desc: "Article description", Content: "Article content"},
}

func ReturnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func ReturnSingleArticle(w http.ResponseWriter, r *http.Request) {
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
func CreateNewArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	// fmt.Fprintf(w, "%+v", string(reqBody))
	var article Article

	json.Unmarshal(reqBody, &article)
	// update global Articles array
	Articles = append(Articles, article)

	json.NewEncoder(w).Encode(article)
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
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

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
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
