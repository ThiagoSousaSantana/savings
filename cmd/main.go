package main

import (
	"log"

	"github.com/ThiagoSousaSantana/saving/cmd/api"
)

func main() {
	server := api.NewAPIServer(":8080", nil)

	if err := server.Run(); err != nil {
		log.Fatal("Error initiating server")
	}
}
