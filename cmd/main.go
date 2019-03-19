package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

var books [] Book

type Book struct {
	ID 		string `json:"id"`
	Isbn 	string `json:"isbn"`
	Title 	string `json:"title"`
	Author  *Author `json:"author"`
}

type Author struct {
	FirstName  string `json:"firstname"`
	LastName   string `json:"lastname"`
}

//GET ALL BOOKS
func getBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(books)
}

//Get a book by ID
func getBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params :=mux.Vars(r)

	for _,item:= range books{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

//Create a book
func createBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(100000))
	books = append(books,book)
	json.NewEncoder(w).Encode(book)
}

//Update a book
func updateBook(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)

	for index,item :=range books{
		if item.ID == params["id"]{
			books = append(books[:index],books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID =  params["id"]
			books = append(books,book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}

//Delete a book
func deleteBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)

	for index,item :=range books{
		if item.ID == params["id"]{
			books = append(books[:index],books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main()  {

	r :=mux.NewRouter()

	// Mock data. TODO :Implement Database

	books = append(books, Book{ID:"1",Isbn:"66676",Title:"Book 1",Author:&Author{FirstName:"John",LastName:"Doe"}})
	books = append(books, Book{ID:"2",Isbn:"66776",Title:"Book 2",Author:&Author{FirstName:"John",LastName:"Paul"}})

	//Route Handlers
	r.HandleFunc("/api/books",getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}",getBook).Methods("GET")
	r.HandleFunc("/api/books",createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}",updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}",deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8001",r))

}