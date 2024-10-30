package functions

import (
	"context"
	"fmt"
	inits "go-mongoose/connection"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindOne(collectionName string, search1 string, find1 string) {
	find := inits.DB.Database(inits.CollectionName).Collection(collectionName)
	var result bson.M

	err := find.FindOne(context.TODO(), bson.M{search1: find1}).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("No results found")
		} else {
			fmt.Println("Database error found", err)
		}
		return
	}

	fmt.Println("Found document:", result)
}