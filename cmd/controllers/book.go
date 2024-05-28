package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"

	"bookinfo/cmd/models"
	"bookinfo/cmd/repository"
)

type Controller struct{}

var books []models.Book

var bookRepo = repository.BookRepository{}

func (c Controller) GetBooks(collection *mongo.Collection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book

		books = bookRepo.GetBooks(collection, book, []models.Book{})

		//fmt.Println(r.Header.Get("User-Agent"))
		json.NewEncoder(w).Encode(books)
	}
}

func (c Controller) GetBook(collection *mongo.Collection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book

		params := mux.Vars(r)
		id, _ := strconv.Atoi(params["id"])

		book = bookRepo.GetBook(collection, book, id)
		json.NewEncoder(w).Encode(book)
	}
}

func (c Controller) AddBook(collection *mongo.Collection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		_ = json.NewDecoder(r.Body).Decode(&book)

		id := bookRepo.AddBook(collection, book)
		json.NewEncoder(w).Encode(fmt.Sprintf("Inserted a single document: %v ", id))
	}
}

func (c Controller) UpdateBook(collection *mongo.Collection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		_ = json.NewDecoder(r.Body).Decode(&book)

		updateResult := bookRepo.UpdateBook(collection, book)
		json.NewEncoder(w).Encode(updateResult)
	}
}

func (c Controller) RemoveBook(collection *mongo.Collection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		book_id, _ := strconv.Atoi(params["id"])

		delResult := bookRepo.RemoveBook(collection, book_id)
		json.NewEncoder(w).Encode(fmt.Sprintf("Deleted: %v", delResult))
	}
}
