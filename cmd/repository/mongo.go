package repository

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"bookinfo/cmd/models"
)

type BookRepository struct{}

func (b BookRepository) GetBooks(collection *mongo.Collection, book models.Book, books []models.Book) []models.Book {
	ctx := context.TODO()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	err = cursor.All(ctx, &books)
	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(ctx)

	return books
}

func (b BookRepository) GetBook(collection *mongo.Collection, book models.Book, id int) models.Book {
	filter := bson.D{{"id", id}}

	err := collection.FindOne(context.TODO(), filter).Decode(&book)
	if err != nil {
		log.Fatal(err)
	}
	// failed if id does not exist
	return book
}

func (b BookRepository) AddBook(collection *mongo.Collection, book models.Book) string {
	insertResult, err := collection.InsertOne(context.TODO(), book)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%v", insertResult.InsertedID)
}

func (b BookRepository) UpdateBook(collection *mongo.Collection, book models.Book) string {
	filter := bson.D{{"id", book.ID}}
	update := bson.D{{"$set", bson.D{{"title", book.Title}, {"author", book.Author}, {"year", book.Year}}}}

	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%v", updateResult)
}

func (b BookRepository) RemoveBook(collection *mongo.Collection, id int) string {
	filter := bson.D{{"id", id}}
	delResult, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%v", delResult)
}
