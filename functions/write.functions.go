package mongoose

import (
	"context"
	"encoding/json"
	"fmt"
	inits "go-mongoose/connection"

	"go.mongodb.org/mongo-driver/bson"
)

func Create(collectionName string, jsonData string) {
	collection := inits.DB.Database(inits.CollectionName).Collection(collectionName)

	var body bson.M
	err := json.Unmarshal([]byte(jsonData), &body)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	result, err := collection.InsertOne(context.TODO(), body)
	if err != nil {
		fmt.Println("Error Inserting Data")
		return
	}

	fmt.Println(result)
}
