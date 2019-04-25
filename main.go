package main

import (
	"net/http"

	"github.com/gokoder/articlemaster/handler"
	"github.com/gorilla/mux"
)

// serves as the starting point of the application
// initializes the mux router and maps the routes to functions
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/articles/{id}", handler.GetArticleByIdHandler()).Methods("GET")
	r.HandleFunc("/articles/{id}", handler.PutArticleHandler()).Methods("PUT")
	r.HandleFunc("/articles", handler.PostArticleHandler()).Methods("POST")
	r.HandleFunc("/tags/{tagName}/{date}", handler.GetArticlesByTagDateHandler()).Methods("GET")
	http.ListenAndServe(":80", r)
}
