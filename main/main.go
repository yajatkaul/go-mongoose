package main

import (
	"fmt"
	MongoDBConnection "go-mongoose/connection"
	mongoose "go-mongoose/functions"
)

func init() {
	//Connect to the db
	MongoDBConnection.Connect("mongodb://localhost:27017/Tourney")
}

func main() {
	filter := `{"gender": "Male"}`
	//update := `{"sports": "Football"}`
	val := mongoose.FindOne("users", filter)
	fmt.Println(val)
}

//mongoose.Create("users", filter) to Insert a doc
//mongoose.Find("users", filter) find all matching docs
//mongoose.FindOne("users", filter) find one matching doc
//mongoose.FindOneandUpdate("users",filter,update) find one and update mathing doc
//mongoose.FindOneandDelete("users", filter) find one and delete matching doc
//mongoose.FindAllandDelete("turves", filter) find all matching docs and delete
//mongoose.FindAllandUpdate("turves", filter,update) find all mathcing docs and update them

/*
val := mongoose.Find("users", filter)
fmt.Println(val[0]["displayName"])
*/