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
		InsertMany(categories []*entity.Category) (*mongo.InsertManyResult, error)
		Update(category *entity.Category) (*mongo.UpdateResult, error)
		DeleteById(id primitive.ObjectID) (*mongo.UpdateResult, error)
		SearchById(id primitive.ObjectID) (entity.Category, error)
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

func (c *category) InsertMany(categories []*entity.Category) (*mongo.InsertManyResult, error) {
	var i []interface{}
	for _, category := range categories {
		category.Id = primitive.NewObjectID()
		category.CreatedAt = time.Now()
		category.Active = true
		i = append(i, category)
	}
	resultMany, err := c.collection.InsertMany(context.TODO(), i)
	if err != nil {
		return nil, err
	}
	return resultMany, nil
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

func (c *category) DeleteById(id primitive.ObjectID) (*mongo.UpdateResult, error) {
	delete := bson.M{
		"$set": bson.M{
			"deleted_at": time.Now(),
			"active":     false,
		},
	}

	result, err := c.collection.UpdateByID(
		context.TODO(),
		id,
		delete,
	)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *category) SearchById(id primitive.ObjectID) (entity.Category, error) {
	filter := bson.M{
		"_id":    &id,
		"active": true,
	}

	var data entity.Category
	if err := c.collection.FindOne(context.TODO(), filter).Decode(&data); err != nil {
		return data, err
	}

	return data, nil
}
