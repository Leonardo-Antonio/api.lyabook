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
	category struct {
		collection *mongo.Collection
	}

	ICategory interface {
		Insert(category *entity.Category) (*mongo.InsertOneResult, error)
		Update(category *entity.Category) (*mongo.UpdateResult, error)
	}
)

func NewCategoty(db *mongo.Database) *category {
	return &category{
		collection: db.Collection(enum.Collection.Categories),
	}
}

func (c *category) Insert(category *entity.Category) (*mongo.InsertOneResult, error) {
	category.Id = primitive.NewObjectID()
	category.CreatedAt = time.Now()
	category.Active = true

	result, err := c.collection.InsertOne(context.TODO(), category)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *category) Update(category *entity.Category) (*mongo.UpdateResult, error) {
	category.UpdatedAt = time.Now()
	update := bson.M{
		"$set": category,
	}

	result, err := c.collection.UpdateByID(
		context.TODO(),
		category.Id, update,
	)

	if err != nil {
		return nil, err
	}

	return result, nil
}
