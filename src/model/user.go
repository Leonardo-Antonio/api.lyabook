package model

import (
	"context"
	"time"

	"github.com/Leonardo-Antonio/api.lyabook/src/entity"
	"github.com/Leonardo-Antonio/api.lyabook/src/utils/enum"
	"github.com/Leonardo-Antonio/api.lyabook/src/utils/errores"
	"github.com/Leonardo-Antonio/api.lyabook/src/utils/key"
	"github.com/Leonardo-Antonio/gobcrypt/gobcrypt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	user struct {
		collection *mongo.Collection
	}

	IUser interface {
		Insert(user *entity.User) (*mongo.InsertOneResult, error)
		Find(credentialsUser *entity.User, flag string) (entity.User, error)
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

func (u *user) Find(credentialsUser *entity.User, flag string) (entity.User, error) {
	var user entity.User

	switch flag {
	case enum.TypeLogin.Dni:
		filter := bson.M{
			"dni":    credentialsUser.Dni,
			"active": true,
		}
		if err := u.collection.FindOne(
			context.TODO(),
			filter,
		).Decode(&user); err != nil {
			return user, err
		}
		return user, nil

	case enum.TypeLogin.Email:
		filter := bson.M{
			"email":  credentialsUser.Email,
			"active": true,
		}
		if err := u.collection.FindOne(
			context.TODO(),
			filter,
		).Decode(&user); err != nil {
			return user, err
		}
	default:
		return user, errores.ErrTypeLogin
	}

	passDecode, err := gobcrypt.Decrypt(user.Password, []byte(key.Password))
	if err != nil {
		return user, err
	}

	if passDecode != credentialsUser.Password {
		return user, errores.ErrInvalidPassword
	}

	return user, nil
}
