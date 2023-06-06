package main

import (
	"unknspec/src/database"
	"unknspec/src/server"
)

func main() {
	db := database.NewMongoStorage("mongodb://localhost:27017/")
	server := server.NewServer(":3000", db)
	server.Run()
}
