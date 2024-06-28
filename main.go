package main

import (
	"github.com/odrianoaliveira/go-service/internal/database"
	"github.com/odrianoaliveira/go-service/internal/server"
	"log"
)

func main() {
	client, err := database.NewClient()
	if err != nil {
		log.Fatalf("error creating database client: %s", err)
	}

	srv := server.NewEchoServer(client)
	if err := srv.Start(); err != nil {
		log.Fatalf("error starting server: %s", err)
	}
}
