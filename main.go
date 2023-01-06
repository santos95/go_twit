package main

import (
	"github.com/santos95/go_twitter/bd"
	"github.com/santos95/go_twitter/handlers"
	"log"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("Error, is not able to connect with the database")
		return
	}

	handlers.Handlers()
}
