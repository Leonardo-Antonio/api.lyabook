package model

import (
	"context"
	"time"

	"github.com/Leonardo-Antonio/api.lyabook/src/entity"
	"github.com/Leonardo-Antonio/api.lyabook/src/utils/enum"
	"github.com/Leonardo-Antonio/api.lyabook/src/utils/errores"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	user struct {
		collection *mongo.Collection
	}
)

func NewUser(db *mongo.Database) *user {
	return &user{
		collection: db.Collection(enum.Collection.Users),
	}
}

func (u *user) Insert(user *entity.User) (*mongo.InsertOneResult, error) {
	user.Id = primitive.NewObjectID()
	user.Active = true
	user.CreateAt = time.Now()

	result, err := u.collection.InsertOne(
		context.TODO(),
		user,
	)

	if err != nil {
		return result, err
	}
	return result, nil
}

func (u *user) Find(credentialsUser *entity.User, flag string) (*entity.User, error) {
	var user *entity.User

	switch flag {
	case enum.TypeLogin.Dni:
		filter := bson.M{
			"dni":      credentialsUser.Dni,
			"password": credentialsUser.Password,
			"active":   true,
		}
		if err := u.collection.FindOne(
			context.TODO(),
			filter,
		).Decode(user); err != nil {
			return nil, err
		}
		return user, nil

	case enum.TypeLogin.Email:
		filter := bson.M{
			"email":    credentialsUser.Email,
			"password": credentialsUser.Password,
			"active":   true,
		}
		if err := u.collection.FindOne(
			context.TODO(),
			filter,
		).Decode(&user); err != nil {
			return nil, err
		}
	default:
		return nil, errores.ErrTypeLogin
	}
	return user, nil
}
