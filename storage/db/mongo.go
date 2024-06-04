package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"bookinfo/entity"
)

type BookDB struct{}

func (b BookDB) GetBooks(collection *mongo.Collection, book entity.Book, books []entity.Book) []entity.Book {
	ctx := context.Background()

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

func (b BookDB) GetBook(collection *mongo.Collection, book entity.Book, id int) entity.Book {
	filter := bson.D{{"id", id}}

	err := collection.FindOne(context.Background(), filter).Decode(&book)
	if err != nil {
		log.Fatal(err)
	}
	// failed if id does not exist
	return book
}

func (b BookDB) AddBook(collection *mongo.Collection, book entity.Book) string {
	insertResult, err := collection.InsertOne(context.Background(), book)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%v", insertResult.InsertedID)
}

func (b BookDB) UpdateBook(collection *mongo.Collection, book entity.Book) string {
	filter := bson.D{{"id", book.ID}}
	update := bson.D{{"$set", bson.D{{"title", book.Title}, {"author", book.Author}, {"year", book.Year}}}}

	updateResult, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%v", updateResult)
}

func (b BookDB) RemoveBook(collection *mongo.Collection, id int) string {
	filter := bson.D{{"id", id}}
	delResult, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%v", delResult)
}
