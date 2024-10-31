package mongoose

import (
	"context"
	"encoding/json"
	"fmt"
	MongoDBConnection "go-mongoose/connection"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindOneandDelete(collectionName string, filterJSON string) {
	collection := MongoDBConnection.DB.Database(MongoDBConnection.CollectionName).Collection(collectionName)

	var filter bson.M
	err := json.Unmarshal([]byte(filterJSON), &filter)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}
	var result bson.M

	err = collection.FindOneAndDelete(context.TODO(), filter).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("No results found to delete")
		} else {
			fmt.Println("Database error found", err)
		}
		return
	}
}