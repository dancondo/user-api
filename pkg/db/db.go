package db

import (
	"context"
	"time"

	"github.com/dancondo/users-api/common"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mongoURI string
var mongoDatabase string

func ConnectDB() *mongo.Client {
	mongoURI = common.GetEnv("MONGO_BASE_URI")
	mongoDatabase = common.GetEnv("MONGO_DATABASE")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))

	if err != nil {
		common.Log.Error(err)
		panic(err)
	}

	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		common.Log.Error(err)
	}

	common.Log.Info("Connected to the Database")
	return client
}

// DB is the client instance
var db *mongo.Client

// GetCollection gets a database collection
func GetCollection(collectionName string) *mongo.Collection {
	if db == nil {
		db = ConnectDB()
	}

	collection := db.Database(mongoDatabase).Collection(collectionName)
	return collection
}
