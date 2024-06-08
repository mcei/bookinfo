package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"

	"bookinfo/domain/book"
	"bookinfo/storage"
)

type Handler struct{}

var books []book.Book

var bookRepo = storage.BookDB{}

func (h Handler) GetBooks(collection *mongo.Collection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var b book.Book

		books = bookRepo.GetBooks(collection, b, []book.Book{})

		//fmt.Println(r.Header.Get("User-Agent"))
		json.NewEncoder(w).Encode(books)
	}
}

func (h Handler) GetBook(collection *mongo.Collection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var b book.Book

		params := mux.Vars(r)
		id, _ := strconv.Atoi(params["id"])

		b = bookRepo.GetBook(collection, b, id)
		json.NewEncoder(w).Encode(b)
	}
}

func (h Handler) AddBook(collection *mongo.Collection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var b book.Book
		_ = json.NewDecoder(r.Body).Decode(&b)

		id := bookRepo.AddBook(collection, b)
		json.NewEncoder(w).Encode(fmt.Sprintf("Inserted a single document: %v ", id))
	}
}

func (h Handler) UpdateBook(collection *mongo.Collection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var b book.Book
		_ = json.NewDecoder(r.Body).Decode(&b)

		updateResult := bookRepo.UpdateBook(collection, b)
		json.NewEncoder(w).Encode(updateResult)
	}
}

func (h Handler) RemoveBook(collection *mongo.Collection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		book_id, _ := strconv.Atoi(params["id"])

		delResult := bookRepo.RemoveBook(collection, book_id)
		json.NewEncoder(w).Encode(fmt.Sprintf("Deleted: %v", delResult))
	}
}
