package main

import (
	"log"

	"github.com/Pivetta21/planning-go/internal/infra/db"
	"github.com/Pivetta21/planning-go/internal/transport"
)

func main() {
	db.OpenConnection()
	defer db.CloseConnection()

	err := transport.StartHttpServer()
	log.Fatal(err)
}
