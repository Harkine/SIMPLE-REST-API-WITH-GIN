package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Body   string `json:"body"`
}

type articles Article

func getArticles(w http.ResponseWriter, r *http.Request) {
	articles := Article{
		ID: "1", Title: "Introduction to Haematology", Author: "Awe Akinsola", Body: "Hematology is the study of the normal and pathologic aspects of blood and blood elements. Blood is a very unique fluid composed of many cellular elements as well as a liquid portion consisting of proteins, amino acids, carbohydrates, lipids, and other macromolecules and low-molecular-weight precursors.",
	}
	fmt.Println("Article endpoint hit")
	json.NewEncoder(w).Encode(articles)
}

func newArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("New article endpoint hit")
	fmt.Fprintf(w, "New article endpoint hit")
}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server has started")
	fmt.Println("Server started at port 8080")
}

func requestHandler() {
	NewRouter := mux.NewRouter().StrictSlash(true)
	NewRouter.HandleFunc("/", homepage)
	NewRouter.HandleFunc("/articles", getArticles)
	NewRouter.HandleFunc("/new", newArticle).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", NewRouter))
}

func main() {
	requestHandler()
}
