package main

import (
	inits "go-mongoose/connection"
	"go-mongoose/functions"
)

func init() {
	inits.SetupDatabase("mongodb://localhost:27017/fitnessFusion")
}

func main() {
	filter := `{"displayName": "Yajat"}`
	functions.Find("users", filter)
}