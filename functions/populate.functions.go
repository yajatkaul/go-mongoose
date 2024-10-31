package mongoose

import (
	"context"
	"fmt"
	MongoDBConnection "go-mongoose/connection"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


func Populate(fieldName string, collectionRef string, document bson.M) bson.M{
	if val, ok := document[fieldName]; ok {
        // Convert val to ObjectId if necessary, depending on how you store references
        if objectId, ok := val.(primitive.ObjectID); ok {
            collection := MongoDBConnection.DB.Database(MongoDBConnection.DBName).Collection(collectionRef)

            // Fetch the referenced document
            var populatedDoc bson.M
            err := collection.FindOne(context.TODO(), bson.M{"_id": objectId}).Decode(&populatedDoc)
            if err != nil {
                if err == mongo.ErrNoDocuments {
                    fmt.Println("No referenced document found")
                } else {
                    fmt.Println("Database error found", err)
                }
                return nil
            }

            return populatedDoc
        }
    }
    return nil
}

func PopulateAll(fieldName string, collectionRef string, documents []bson.M) []bson.M {
    populatedDocuments := make([]bson.M, 0)

    collection := MongoDBConnection.DB.Database(MongoDBConnection.DBName).Collection(collectionRef)

    for _, doc := range documents {
        if val, ok := doc[fieldName]; ok {
			// Convert val to ObjectId if necessary, depending on how you store references
			if objectId, ok := val.(primitive.ObjectID); ok {
				// Fetch the referenced document
				var populatedDoc bson.M
				err := collection.FindOne(context.TODO(), bson.M{"_id": objectId}).Decode(&populatedDoc)
				if err != nil {
					if err == mongo.ErrNoDocuments {
						fmt.Println("No referenced document found")
					} else {
						fmt.Println("Database error found", err)
					}
				}
				populatedDocuments = append(populatedDocuments, populatedDoc)
			}
		}
    }

    return populatedDocuments
}