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
	complaintsBook struct {
		collection *mongo.Collection
	}
	IComplaintsBook interface {
		InsertOne(data *entity.ComplaintsBook) (*mongo.InsertOneResult, error)
		FindAll() (entity.ComplaintsBooks, error)
		CountClaims() (int64, error)
	}
)

func NewComplaintsBook(db *mongo.Database) *complaintsBook {
	return &complaintsBook{
		collection: db.Collection(enum.Collection.ComplaintsBooks),
	}
}

func (c *complaintsBook) InsertOne(data *entity.ComplaintsBook) (*mongo.InsertOneResult, error) {
	data.Id = primitive.NewObjectID()
	data.CreatedAt = time.Now()
	data.Active = true

	result, err := c.collection.InsertOne(
		context.TODO(),
		data,
	)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *complaintsBook) FindAll() (entity.ComplaintsBooks, error) {
	filter := bson.M{
		"active": true,
	}
	cursor, err := c.collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var data entity.ComplaintsBooks
	if err := cursor.All(context.TODO(), &data); err != nil {
		return nil, err
	}

	return data, nil
}

func (c *complaintsBook) CountClaims() (int64, error) {
	filter := bson.M{
		"active": true,
	}
	amount, err := c.collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		return 0, err
	}

	return amount, nil
}
