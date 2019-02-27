package main

import (
	"encoding/json"
	"log"
	"net/http"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
)

//Book Struct (Model)
type Book struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

func getBooks(w http.ResponseWriter, r *http.Request){

}


func getBook(w http.ResponseWriter, r *http.Request){
	
}

func createBook(w http.ResponseWriter, r *http.Request){
	
}

func updateBook(w http.ResponseWriter, r *http.Request){
	
}

func deleteBook(w http.ResponseWriter, r *http.Request){
	
}

func main(){
	//int router
	r := mux.NewRouter()

	//route handlers
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBooks).Methods("GET")
	r.HandleFunc("/api/books", getBooks).Methods("POST")
	r.HandleFunc("/api/books/{id}", getBooks).Methods("PUT")
	r.HandleFunc("/api/books/{id}", getBooks).Methods("DELETE")
	http.ListenAndServe(":8000", r)
	//log.Fatal()
}