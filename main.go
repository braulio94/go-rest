package main

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"math/rand"
	"strconv"
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

//init books var as a slice Book struct

var books [] Book

func getBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}


func getBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //get params
	for _, item := range books {
		if item.ID == params["id"]{
		  json.NewEncoder(w).Encode(item)
		  return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

func createBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"]{
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}

func deleteBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"]{
			books = append(books[:index], books[index+1:]...)
		}
		break
	}
	json.NewEncoder(w).Encode(books)
}

func main(){
	//int router
	r := mux.NewRouter()

	books = append(books, Book{ID: "1", Isbn: "123124", Title: "Book one", Author: &Author{Firstname: "Braulio", Lastname: "Cassule"}})
	books = append(books, Book{ID: "2", Isbn: "123124", Title: "Book Two", Author: &Author{Firstname: "Braulio", Lastname: "Cassule"}})
	books = append(books, Book{ID: "3", Isbn: "123124", Title: "Book Three", Author: &Author{Firstname: "Braulio", Lastname: "Cassule"}})
	books = append(books, Book{ID: "4", Isbn: "123124", Title: "Book Four", Author: &Author{Firstname: "Braulio", Lastname: "Cassule"}})
	books = append(books, Book{ID: "5", Isbn: "123124", Title: "Book Five", Author: &Author{Firstname: "Braulio", Lastname: "Cassule"}})
	books = append(books, Book{ID: "6", Isbn: "123124", Title: "Book Six", Author: &Author{Firstname: "Braulio", Lastname: "Cassule"}})
	books = append(books, Book{ID: "7", Isbn: "123124", Title: "Book Seven", Author: &Author{Firstname: "Braulio", Lastname: "Cassule"}})

	//route handlers
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	http.ListenAndServe(":8000", r)
	//log.Fatal()
}