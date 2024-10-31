package mongoose

import (
	"context"
	"encoding/json"
	"fmt"
	MongoDBConnection "go-mongoose/connection"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindOne(collectionName string, filterJSON ...string) bson.M{
	find := MongoDBConnection.DB.Database(MongoDBConnection.CollectionName).Collection(collectionName)
	var result bson.M

	var filter bson.M

	if len(filterJSON) > 0 && filterJSON[0] != "" {
		err := json.Unmarshal([]byte(filterJSON[0]), &filter)
		if err != nil {
			fmt.Println("Error parsing JSON:", err)
			return nil
		}
	} else {
		filter = bson.M{} // Default filter to match all documents
	}

	err := find.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("No results found")
		} else {
			fmt.Println("Database error found", err)
		}
		return nil
	}

	//fmt.Println("Found document:", result)
	return result
}

func Find(collectionName string, filterJSON ...string) []bson.M{
	collection := MongoDBConnection.DB.Database(MongoDBConnection.CollectionName).Collection(collectionName)
	
	var filter bson.M

	if len(filterJSON) > 0 && filterJSON[0] != "" {
		err := json.Unmarshal([]byte(filterJSON[0]), &filter)
		if err != nil {
			fmt.Println("Error parsing JSON:", err)
			return nil
		}
	} else {
		filter = bson.M{} 
	}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("Database error found", err)
		return nil
	}
	defer cursor.Close(context.TODO())

	var results []bson.M

	if err = cursor.All(context.TODO(), &results); err != nil {
		fmt.Println("Error reading cursor:", err)
		return nil
	}

	if len(results) == 0 {
		fmt.Println("No results found")
		return nil
	}

	/*
	for _, result := range results {
		fmt.Println("Found document:", result)
	}
	*/

	return results
}