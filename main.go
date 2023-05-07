package main

import (
	"fmt"
	"log"
)

func main() {
	passwd := "devpasswd"
	storage := NewPostgresStorage(fmt.Sprintf("user=postgres password=%s sslmode=disable dbname=unknspec", passwd))
	err := storage.InitTables()
	if err != nil {
		log.Fatal(err)
	}
	server := NewServer("localhost:3000", storage)
	log.Fatalln(server.Run())
}
