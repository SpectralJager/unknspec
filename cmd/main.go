package main

import (
	"log"
	"unknspec/core"
)

func main() {
	app := core.InitApp()
	log.Fatal(app.Listen("localhost:8080"))
}
