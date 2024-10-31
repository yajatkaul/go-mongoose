package main

import (
	MongoDBConnection "go-mongoose/connection"
)

func init() {
	//Connect to the db
	MongoDBConnection.Connect("mongodb://localhost:27017/fitnessFusion")
}

func main() {
	//filter := `{"displayName": "Yajat","email":"mc.crafter1973@gmail.com"}`
	//update := `{"displayName": "Yajat2", "email":"skibidisigma"}`
	
}

//mongoose.Create("users", filter) to Insert a doc
//mongoose.Find("users", filter) find all matching docs
//mongoose.FindOne("users", filter) find one matching doc
//mongoose.FindOneandUpdate("users",filter,update) find one and update mathing doc
//mongoose.FindOneandDelete("users", filter) find one and delete matching doc