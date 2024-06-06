package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"bookinfo/cmd/rest"
	"bookinfo/storage/db"
)

func main() {
	router := mux.NewRouter()

	handler := rest.Handler{}

	collection := db.Connect()

	router.HandleFunc("/books", handler.GetBooks(collection)).Methods("GET")
	router.HandleFunc("/books/{id}", handler.GetBook(collection)).Methods("GET")
	router.HandleFunc("/books", handler.AddBook(collection)).Methods("POST")
	router.HandleFunc("/books", handler.UpdateBook(collection)).Methods("PUT")
	router.HandleFunc("/books/{id}", handler.RemoveBook(collection)).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
