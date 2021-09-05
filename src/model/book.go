package model

import (
	"context"
	"time"

	"github.com/Leonardo-Antonio/api.lyabook/src/entity"
	"github.com/Leonardo-Antonio/api.lyabook/src/utils/enum"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	book struct {
		collection *mongo.Collection
	}

	Ibook interface {
		Insert(book *entity.Book) (*mongo.InsertOneResult, error)
		Update(book *entity.Book) (*mongo.UpdateResult, error)
		UpdatePriceCurrent(id primitive.ObjectID, priceCurrent float64) (*mongo.UpdateResult, error)
		FindBookById(id primitive.ObjectID) (book entity.Book, err error)
		DeleteById(id primitive.ObjectID) (*mongo.UpdateResult, error)
	}
)

func NewBook(db *mongo.Database) *book {
	return &book{
		collection: db.Collection(enum.Collection.Books),
	}
}

func (b *book) Insert(book *entity.Book) (*mongo.InsertOneResult, error) {
	book.Id = primitive.NewObjectID()
	book.PriceBefore = book.PriceCurrent
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

func (b *book) Update(book *entity.Book) (*mongo.UpdateResult, error) {
	book.UpdatedAt = time.Now()

	update := bson.M{
		"$set": book,
	}

	result, err := b.collection.UpdateByID(
		context.TODO(),
		book.Id,
		update,
	)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (b *book) UpdatePriceCurrent(id primitive.ObjectID, priceCurrent float64) (*mongo.UpdateResult, error) {
	update := bson.M{
		"$set": bson.M{
			"price_current": priceCurrent,
			"updated_at":    time.Now(),
		},
	}

	result, err := b.collection.UpdateByID(
		context.TODO(),
		id,
		update,
	)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (b *book) FindBookById(id primitive.ObjectID) (book entity.Book, err error) {
	filter := bson.M{
		"_id":    id,
		"active": true,
	}

	if err := b.collection.FindOne(context.TODO(), filter).Decode(&book); err != nil {
		return book, err
	}

	return
}

func (b *book) DeleteById(id primitive.ObjectID) (*mongo.UpdateResult, error) {
	update := bson.M{
		"$set": bson.M{
			"deleted_at": time.Now(),
			"active":     false,
		},
	}
	result, err := b.collection.UpdateByID(context.TODO(), id, update)
	if err != nil {
		return nil, err
	}

	return result, nil
}
