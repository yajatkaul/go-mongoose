package main

import (
	inits "go-mongoose/connection"
	"go-mongoose/functions"
)

func init() {
	inits.SetupDatabase("mongodb://localhost:27017/fitnessFusion")
}

func main() {
	//filter := `{"displayName": "Yajat","email":"mc.crafter1973@gmail.com"}`
	functions.FindOne("users")
}

//functions.Create("users", filter) to Insert a doc
//functions.Find("users", filter) find all matching docs
//functions.FindOne("users", filter) find one matching doc