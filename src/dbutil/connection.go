package dbutil

import (
	"context"
	"log"

	"github.com/Leonardo-Antonio/api.lyabook/src/utils/env"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetConnection() (db *mongo.Database) {
	clientOptions := options.Client().ApplyURI(env.Data.UrlMongo)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalln(err)
	}

	if err := client.Ping(context.TODO(), nil); err != nil {
		log.Fatalln(err)
	}

	db = client.Database(env.Data.DBName)
	fildsUnique := NewCollectionIndex(db)
	fildsUnique.createIndexUser()
	fildsUnique.createIndexCategory()
	return
}
