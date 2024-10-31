package mongoose

import (
	"context"
	"encoding/json"
	"fmt"

	MongoDBConnection "github.com/yajatkaul/go-mongoose/src/src/connection"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Create(collectionName string, jsonData string) {
	collection := MongoDBConnection.DB.Database(MongoDBConnection.DBName).Collection(collectionName)

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
	collection := MongoDBConnection.DB.Database(MongoDBConnection.DBName).Collection(collectionName)

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

func FindAllandUpdate(collectionName string, filterJSON string, updateJSON string){
	collection := MongoDBConnection.DB.Database(MongoDBConnection.DBName).Collection(collectionName)

	//Parsing filter to json
	var filter bson.M
	err := json.Unmarshal([]byte(filterJSON), &filter)
	if err != nil {
		fmt.Println("Error parsing filter JSON:", err)
		return
	}

	//Parsing update to json
	var update bson.M
	err = json.Unmarshal([]byte(updateJSON), &update)
	if err != nil {
		fmt.Println("Error parsing update JSON:", err)
		return
	}

	updateDoc := bson.M{"$set": update}

	updateResult, err := collection.UpdateMany(context.TODO(), filter, updateDoc)
	if err != nil {
		fmt.Println("Error updating documents:", err)
		return
	}

	if updateResult.MatchedCount == 0 {
		fmt.Println("No matching documents found to update")
	} else {
		fmt.Printf("Matched %d documents and updated %d documents\n", updateResult.MatchedCount, updateResult.ModifiedCount)
	}
}
