package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/ThiagoSousaSantana/saving/cmd/api"
)

func main() {
	dbUri := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"localhost", "5432", "postgres", "postgres", "savings")

	d, err := sql.Open("postgres", dbUri)
	if err != nil {
		return
	}
	defer d.Close()

	server := api.NewAPIServer(":8080", d)

	if err := server.Run(); err != nil {
		log.Fatal("Error initiating server")
	}
}
