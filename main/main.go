package main

import (
	MongoDBConnection "go-mongoose/connection"
	mongoose "go-mongoose/functions"
)

func init() {
	//Connect to the db
	MongoDBConnection.Connect("mongodb://localhost:27017/fitnessFusion")
}

func main() {
	//filter := `{"displayName": "Yajat","email":"mc.crafter1973@gmail.com"}`
	mongoose.FindOne("users")
}

//functions.Create("users", filter) to Insert a doc
//functions.Find("users", filter) find all matching docs
//functions.FindOne("users", filter) find one matching doc