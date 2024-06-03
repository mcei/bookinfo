package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"bookinfo/controllers"
	"bookinfo/storage/db"
)

func main() {
	router := mux.NewRouter()

	controller := controllers.Controller{}

	collection := db.Connect()

	router.HandleFunc("/books", controller.GetBooks(collection)).Methods("GET")
	router.HandleFunc("/books/{id}", controller.GetBook(collection)).Methods("GET")
	router.HandleFunc("/books", controller.AddBook(collection)).Methods("POST")
	router.HandleFunc("/books", controller.UpdateBook(collection)).Methods("PUT")
	router.HandleFunc("/books/{id}", controller.RemoveBook(collection)).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
