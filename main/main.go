package main

import (
	inits "go-mongoose/connection"
	"go-mongoose/functions"
)

func init() {
	inits.SetupDatabase("mongodb://localhost:27017/test")
}

func main() {
	functions.FindOne("users", "displayName","Yajat")
}