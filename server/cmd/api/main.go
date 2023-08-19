package main

import (
	"github.com/Pivetta21/planning-go/internal/infra/db"
	"github.com/Pivetta21/planning-go/internal/transport"
)

func main() {
	db.OpenConnection()
	defer db.CloseConnection()

	transport.StartHttpServer()
}
