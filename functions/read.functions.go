package mongoose

import (
	"context"
	"encoding/json"
	"fmt"
	inits "go-mongoose/connection"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindOne(collectionName string, filterJSON ...string) {
	find := inits.DB.Database(inits.CollectionName).Collection(collectionName)
	var result bson.M

	var filter bson.M

	if len(filterJSON) > 0 && filterJSON[0] != "" {
		err := json.Unmarshal([]byte(filterJSON[0]), &filter)
		if err != nil {
			fmt.Println("Error parsing JSON:", err)
			return
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
		return
	}

	fmt.Println("Found document:", result)
}

func Find(collectionName string, filterJSON ...string) {
	find := inits.DB.Database(inits.CollectionName).Collection(collectionName)
	
	var filter bson.M

	if len(filterJSON) > 0 && filterJSON[0] != "" {
		err := json.Unmarshal([]byte(filterJSON[0]), &filter)
		if err != nil {
			fmt.Println("Error parsing JSON:", err)
			return
		}
	} else {
		filter = bson.M{} 
	}

	cursor, err := find.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("Database error found", err)
		return
	}
	defer cursor.Close(context.TODO())

	var results []bson.M

	if err = cursor.All(context.TODO(), &results); err != nil {
		fmt.Println("Error reading cursor:", err)
		return
	}

	if len(results) == 0 {
		fmt.Println("No results found")
		return
	}

	for _, result := range results {
		fmt.Println("Found document:", result)
	}
}