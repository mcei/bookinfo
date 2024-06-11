package storage

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"bookinfo/domain/book"
)

type BookDB struct{}

func (b BookDB) GetBooks(ctx context.Context, c *mongo.Collection, books []book.Book) ([]book.Book, error) {
	cursor, err := c.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &books)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	return books, nil
}

func (b BookDB) GetBook(ctx context.Context, c *mongo.Collection, book book.Book, id int) (book.Book, error) {
	filter := bson.D{{"id", id}}
	err := c.FindOne(ctx, filter).Decode(&book)
	if err != nil {
		return book, err
	}
	// failed if id does not exist
	return book, nil
}

func (b BookDB) AddBook(ctx context.Context, c *mongo.Collection, book book.Book) (string, error) {
	insertResult, err := c.InsertOne(ctx, book)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", insertResult.InsertedID), nil
}

func (b BookDB) UpdateBook(ctx context.Context, c *mongo.Collection, book book.Book) (string, error) {
	filter := bson.D{{"id", book.ID}}
	update := bson.D{{"$set", bson.D{{"title", book.Title}, {"author", book.Author}, {"year", book.Year}}}}
	updateResult, err := c.UpdateOne(ctx, filter, update)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", updateResult), nil
}

func (b BookDB) RemoveBook(ctx context.Context, c *mongo.Collection, id int) (string, error) {
	filter := bson.D{{"id", id}}
	delResult, err := c.DeleteOne(ctx, filter)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", delResult), nil
}
