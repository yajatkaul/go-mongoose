# Go Mongoose

[![Go Report Card](https://goreportcard.com/badge/github.com/yajatkaul/go-mongoose)](https://goreportcard.com/report/github.com/yajatkaul/go-mongoose)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

**Go Mongoose** is a lightweight MongoDB wrapper for Go, inspired by Mongoose, designed to simplify database operations and improve code readability.

## Features

- **Simple API**: Intuitive functions for common database operations.
- **Populate**: Easily reference and populate fields from related collections.
- **Flexible**: Supports various MongoDB operations including Create, Read, Update, and Delete (CRUD).

## Installation

To install Go Mongoose, use the following command:

```bash
go get github.com/yajatkaul/go-mongoose
go mod tidy
```

## Getting Started
Here’s a quick guide to get you started with Go Mongoose.

# 1. Connect to the Database
Before performing any operations, you need to connect to your MongoDB database. Here's an example of how to connect:
```bash
package main

import (
    MongoDBConnection "github.com/yajatkaul/go-mongoose/MongoDBConnection"
)

func init() {
    // Connect to the MongoDB database
    MongoDBConnection.Connect("uri", "Database Name")
}
```

# 2. Performing Operations
Go Mongoose provides several functions to interact with your collections. Below are some common operations:
  ```bash
  //filter and updates are just JSON formats
  filter := `{"displayName": "Yajat","email":"def@sds.com"}`
  update := `{"displayName": "Yajat2", "email":"abc@sds.com"}`
  ```
* Create a Document: Insert a document into a collection.
  ```bash
  mongoose.Create("collectionName", filter) // filter is the document to insert
  ```
* Find Documents: Retrieve documents from a collection.
  ```bash
  val := mongoose.Find("collectionName", filter) // Returns []bson.M
  ```
* Find One Document: Retrieve a single document.
  ```bash
  val := mongoose.FindOne("collectionName", filter) // Returns bson.M
  ```
* Update a Document: Update a single document.
  ```bash
  mongoose.FindOneAndUpdate("collectionName", filter, update)
  ```
* Delete a Document: Remove a single document.
  ```bash
  mongoose.FindOneAndDelete("collectionName", filter)
  ```
* Find and Delete Multiple Documents: Find and delete all matching documents.
  ```bash
  mongoose.FindAllAndDelete("collectionName", filter)
  ```
* Find and Update Multiple Documents: Find and update all matching documents.
  ```bash
  mongoose.FindAllAndUpdate("collectionName", filter, update)
  ```
# 3. Populating References
You can easily populate fields from related collections:
* Populate a Single Reference:
  ```bash
  document := mongoose.FindOne("collectionName")
  fmt.Println(mongoose.Populate("Name of the field to Populate", "Refference with which you want to populate", document))
  ```
* Populate All References:
  ```bash
  document := mongoose.Find("collectionName")
  fmt.Println(mongoose.PopulateAll("Name of the field to Populate", "Refference with which you want to populate", document))
  ```
## Example Usage
Here’s a complete example demonstrating the usage of Go Mongoose:
```bash
package main

import (
    "fmt"
    MongoDBConnection "github.com/yajatkaul/go-mongoose/MongoDBConnection"
	"github.com/yajatkaul/go-mongoose/mongoose"
)

func init() {
    // Connect to the MongoDB database
    MongoDBConnection.Connect("mongodb://localhost:27017/", "Tourney")
}

func main() {
    // Find all tourneys
    document := mongoose.Find("users")
    
    // Populate host information
    fmt.Println(mongoose.PopulateAll("followers", "users", document))
}
```
## Contributing
Contributions are welcome! Please feel free to submit a Pull Request or open an Issue if you find bugs or have suggestions for improvements.

## License
This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgements
* Inspired by the Mongoose library for Node.js.
* Thanks to the MongoDB community for their support and resources.

## Changes required

> -Add populate funcionality (This has been done but this is not very efficient and is not exactly like mongoose maybe someday, for now maybe use it or just nested Find functions)

> -Struct funcionalty idk atp

> ~~-Few more functions like findoneandupdate, findonebyid etc~~

> -IDK

