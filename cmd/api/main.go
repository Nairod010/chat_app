package main

import (
	"log"

	"github.com/nairod010/chat_app/internal/database"
	"github.com/nairod010/chat_app/internal/server"
)

func main() {
	db, err := database.NewPostgresService()
	if err != nil {
		log.Fatal(err)
	}
	api := server.NewAPIServer(":3000", db)
	api.Server()
}
