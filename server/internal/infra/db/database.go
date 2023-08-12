package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/Pivetta21/planning-go/internal/configs"
	_ "github.com/lib/pq"
)

type Context struct {
	// Database connection pool
	Conn *sql.DB

	// Default duration timeout (i.e., in seconds) for contexts
	DefaultTimeout time.Duration
}

var Ctx *Context

func OpenConnection() {
	cfg := configs.DBConfig.Postgres
	connectionString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name, cfg.SslMode,
	)

	conn, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if err := conn.PingContext(ctx); err != nil {
		log.Fatal(err)
	}

	Ctx = &Context{
		Conn:           conn,
		DefaultTimeout: 3 * time.Second,
	}

	log.Println("database connection established")
}

func CloseConnection() {
	err := Ctx.Conn.Close()
	if err != nil {
		log.Fatal("database connection could not be closed")
	}

	log.Println("database connection closed")
}
