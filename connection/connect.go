package MongoDBConnection

import (
	"context"
	"log"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client
var CollectionName string

func Connect(uri string) *mongo.Client{
    CollectionName = GetCollection(uri)
	var err error
	clientOptions := options.Client().ApplyURI(GetBaseURI(uri))

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

func GetBaseURI(collectionUri string) string {
	slice := strings.Split(collectionUri, "/")
	return strings.Join(slice[:len(slice)-1], "/") 
}

func GetCollection(collectionUri string) string{
    slice := strings.Split(collectionUri, "/")
    return slice[len(slice)-1]
}