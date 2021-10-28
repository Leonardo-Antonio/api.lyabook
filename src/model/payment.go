package model

import (
	"context"

	"github.com/Leonardo-Antonio/api.lyabook/src/entity"
	"github.com/Leonardo-Antonio/api.lyabook/src/utils/enum"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	payment struct {
		collection *mongo.Collection
	}

	IPayment interface {
		GetById(id primitive.ObjectID) ([]*entity.PaymentClient, error)
	}
)

func NewPayment(db *mongo.Database) *payment {
	return &payment{collection: db.Collection(enum.Collection.Payment)}
}

func (p *payment) GetById(id primitive.ObjectID) ([]*entity.PaymentClient, error) {
	pipeline := []bson.M{
		{
			"$match": bson.M{
				"_id": id,
			},
		},
		{
			"$lookup": bson.M{
				"from":         "users",
				"localField":   "id_client",
				"foreignField": "_id",
				"as":           "client",
			},
		},
	}

	cursor, err := p.collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var paymentClient []*entity.PaymentClient
	if err := cursor.All(context.TODO(), &paymentClient); err != nil {
		return nil, err
	}

	return paymentClient, nil
}
