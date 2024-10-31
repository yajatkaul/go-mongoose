package mongoose

import (
	"context"
	"encoding/json"
	"fmt"
	inits "go-mongoose/connection"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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

func FindOneandUpdate(collectionName string, filterJson string, updateJson string){
	collection := inits.DB.Database(inits.CollectionName).Collection(collectionName)

	var filter bson.M
	err := json.Unmarshal([]byte(filterJson), &filter)
	if err != nil {
		fmt.Println("Error parsing filter JSON:", err)
		return
	}

	var update bson.M
	err = json.Unmarshal([]byte(updateJson), &update)
	if err != nil {
		fmt.Println("Error parsing update JSON:", err)
		return
	}

	updateDoc := bson.M{"$set": update}

	var updatedDoc bson.M
	err = collection.FindOneAndUpdate(context.TODO(), filter, updateDoc).Decode(&updatedDoc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("No document found to update")
		} else {
			fmt.Println("Error updating document:", err)
		}
		return
	}

	fmt.Println("Updated document:", updatedDoc)
}
