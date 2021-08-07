package dbutil

import (
	"context"
	"log"

	"github.com/Leonardo-Antonio/api.lyabook/src/utils/enum"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type collection struct {
	db *mongo.Database
}

func NewCollectionIndex(db *mongo.Database) *collection {
	return &collection{db}
}

func (c *collection) createIndexUser() {
	_, err := c.db.Collection(enum.Collection.Users).
		Indexes().
		CreateOne(
			context.TODO(),
			mongo.IndexModel{
				Keys: bson.D{
					{Key: "email", Value: 1},
					{Key: "dni", Value: 1},
				},
				Options: options.Index().SetUnique(true),
			},
		)
	if err != nil {
		log.Fatalln(err)
	}
}
