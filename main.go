package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {

	articles := Articles{
		Article{Title: "Test Title", Desc: "Test Description", Content: "Hello World"},
	}

	fmt.Println("Endpoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}

func testPostArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test POST endpoint worked")
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", handler)
	router.HandleFunc("/articles", allArticles).Methods("GET")
	router.HandleFunc("/articles", testPostArticles).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}
