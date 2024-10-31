package MongoDBConnection

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client
var DBName string

func Connect(uri string, dbName string) *mongo.Client {
	DBName = dbName
	var err error
	clientOptions := options.Client().ApplyURI(uri)

	DB, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB!")
	return DB
}
