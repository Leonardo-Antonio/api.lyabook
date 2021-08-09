package model

import (
	"context"
	"time"

	"github.com/Leonardo-Antonio/api.lyabook/src/entity"
	"github.com/Leonardo-Antonio/api.lyabook/src/utils/enum"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	book struct {
		collection *mongo.Collection
	}

	Ibook interface {
		Insert(book *entity.Book) (*mongo.InsertOneResult, error)
	}
)

func NewBook(db *mongo.Database) *book {
	return &book{
		collection: db.Collection(enum.Collection.Books),
	}
}

func (b *book) Insert(book *entity.Book) (*mongo.InsertOneResult, error) {
	book.Id = primitive.NewObjectID()
	book.CreatedAt = time.Now()
	book.Active = true

	result, err := b.collection.InsertOne(
		context.TODO(),
		&book,
	)

	if err != nil {
		return nil, err
	}

	return result, nil
}
