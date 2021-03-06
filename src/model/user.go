package model

import (
	"context"
	"errors"
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
		FindById(id primitive.ObjectID) (entity.User, error)
		Insert(user *entity.User) (*mongo.InsertOneResult, error)
		Find(credentialsUser *entity.User, flag string) (entity.User, error)
		VerifyAccount(email, code string) (entity.User, error)
		VerifyAccountById(id primitive.ObjectID, code string) (entity.User, error)
		Update(user *entity.User) (*mongo.UpdateResult, error)
		FindUsersWithEmail() (entity.Users, error)
		FindAllUsersByRol(rol string) (entity.Users, error)
		DeleteById(id primitive.ObjectID) (*mongo.UpdateResult, error)
		DeleteByIdFisico(id primitive.ObjectID) (*mongo.DeleteResult, error)
		CountUserByRol(rol string) (int64, error)
	}
)

func NewUser(db *mongo.Database) *user {
	return &user{
		collection: db.Collection(enum.Collection.Users),
	}
}

func (u *user) Insert(user *entity.User) (*mongo.InsertOneResult, error) {
	if user.Id.IsZero() {
		user.Id = primitive.NewObjectID()
	}
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

func (u *user) FindById(id primitive.ObjectID) (entity.User, error) {
	filter := bson.M{
		"_id":    id,
		"active": true,
	}

	var dataUser entity.User
	if err := u.collection.FindOne(context.TODO(), filter).Decode(&dataUser); err != nil {
		return entity.User{}, err
	}

	return dataUser, nil
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

func (u *user) VerifyAccount(email, code string) (entity.User, error) {
	filter := bson.M{
		"email":             email,
		"verification_code": code,
	}
	var dataUser entity.User
	if err := u.collection.FindOne(
		context.TODO(), filter,
	).Decode(&dataUser); err != nil {
		return dataUser, err
	}

	dataUser.Active = true
	result, err := u.Update(&dataUser)
	if err != nil {
		return dataUser, err
	}

	if result.ModifiedCount != 1 {
		return dataUser, errors.New("no se pudo activar la cuenta <" + dataUser.Email + ">")
	}

	return dataUser, nil
}

func (u *user) VerifyAccountById(id primitive.ObjectID, code string) (entity.User, error) {
	filter := bson.M{
		"_id":               id,
		"verification_code": code,
	}
	var dataUser entity.User
	if err := u.collection.FindOne(
		context.TODO(), filter,
	).Decode(&dataUser); err != nil {
		return dataUser, err
	}

	dataUser.Active = true
	result, err := u.Update(&dataUser)
	if err != nil {
		return dataUser, err
	}

	if result.ModifiedCount != 1 {
		return dataUser, errors.New("no se pudo activar la cuenta <" + dataUser.Email + ">")
	}

	return dataUser, nil
}

func (u *user) Update(user *entity.User) (*mongo.UpdateResult, error) {
	user.UpdateAt = time.Now()
	update := bson.M{
		"$set": user,
	}
	result, err := u.collection.UpdateByID(context.TODO(), user.Id, update)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (b *user) FindUsersWithEmail() (entity.Users, error) {
	filter := bson.M{
		"email": bson.M{
			"$exists": true,
		},
		"rol":    "Client",
		"active": true,
	}

	cursor, err := b.collection.Find(
		context.TODO(), filter,
	)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.TODO())

	var users entity.Users
	if err := cursor.All(context.TODO(), &users); err != nil {
		return nil, err
	}

	return users, nil
}

func (u *user) FindAllUsersByRol(rol string) (entity.Users, error) {
	filter := bson.M{
		"rol": rol,
	}
	cursor, err := u.collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var users entity.Users
	if err := cursor.All(context.TODO(), &users); err != nil {
		return nil, err
	}

	return users, nil
}

func (u *user) DeleteById(id primitive.ObjectID) (*mongo.UpdateResult, error) {
	delete := bson.M{
		"$set": bson.M{
			"deleted_at": time.Now(),
			"active":     false,
		},
	}

	result, err := u.collection.UpdateByID(
		context.TODO(),
		id,
		delete,
	)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u *user) DeleteByIdFisico(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	filter := bson.M{
		"_id": id,
	}
	result, err := u.collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u *user) CountUserByRol(rol string) (int64, error) {
	filter := bson.M{
		"rol": rol,
	}
	amount, err := u.collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		return 0, err
	}

	return amount, nil
}
