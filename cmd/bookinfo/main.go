package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"bookinfo/cmd/rest"
)

func main() {
	router := mux.NewRouter()
	handler := rest.NewHandler()
	ctx := context.Background()

	// Как здесь использовать функцию? "сразу методы объекта в маршрутизатор привязывать"
	router.HandleFunc("/books", handler.GetBooks(ctx)).Methods("GET")
	router.HandleFunc("/books/{id}", handler.GetBook(ctx)).Methods("GET")
	router.HandleFunc("/books", handler.AddBook(ctx)).Methods("POST")
	router.HandleFunc("/books", handler.UpdateBook(ctx)).Methods("PUT")
	router.HandleFunc("/books/{id}", handler.RemoveBook(ctx)).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
