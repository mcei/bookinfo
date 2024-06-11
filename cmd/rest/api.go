package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"

	"bookinfo/domain/book"
	"bookinfo/storage"
)

type Handler struct {
	collection *mongo.Collection
}

func NewHandler() *Handler {
	return &Handler{
		storage.Connect(),
	}
}

var bookRepo = storage.BookDB{}

// как сюда передавать контекст из мейна?

func (h Handler) GetBooks(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		books, err := bookRepo.GetBooks(ctx, h.collection, []book.Book{})
		if err != nil {
			log.Fatal(err) // TODO how to handle error?
		}
		//fmt.Println(r.Header.Get("User-Agent"))
		json.NewEncoder(w).Encode(books)
	}
}

func (h Handler) GetBook(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var b book.Book
		params := mux.Vars(r)
		id, _ := strconv.Atoi(params["id"])
		b, err := bookRepo.GetBook(ctx, h.collection, b, id)
		if err != nil {
			log.Fatal(err) // TODO how to handle error?
		}
		json.NewEncoder(w).Encode(b)
	}
}

func (h Handler) AddBook(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var b book.Book
		_ = json.NewDecoder(r.Body).Decode(&b)
		id, err := bookRepo.AddBook(ctx, h.collection, b)
		if err != nil {
			log.Fatal(err) // TODO how to handle error?
		}
		json.NewEncoder(w).Encode(fmt.Sprintf("Inserted a single document: %v ", id))
	}
}

func (h Handler) UpdateBook(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var b book.Book
		_ = json.NewDecoder(r.Body).Decode(&b)
		updateResult, err := bookRepo.UpdateBook(ctx, h.collection, b)
		if err != nil {
			log.Fatal(err) // TODO how to handle error?
		}
		json.NewEncoder(w).Encode(updateResult)
	}
}

func (h Handler) RemoveBook(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		book_id, _ := strconv.Atoi(params["id"])
		delResult, err := bookRepo.RemoveBook(ctx, h.collection, book_id)
		if err != nil {
			log.Fatal(err) // TODO how to handle error?
		}
		json.NewEncoder(w).Encode(fmt.Sprintf("Deleted: %v", delResult))
	}
}
