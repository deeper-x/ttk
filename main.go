package main

import (
	"context"
	"log"

	"github.com/deeper-x/ttk/settings"
	"github.com/jackc/pgx/v5"
)

func main() {
	URI, err := settings.GetDBURI()
	if err != nil {
		panic(err)
	}
	conn, err := pgx.Connect(context.Background(), URI)
	if err != nil {
		panic(err)
	}

	defer conn.Close(context.Background())

	var email string
	err = conn.QueryRow(context.Background(), "SELECT 'pong'").Scan(&email)
	if err != nil {
		panic(err)
	}

	log.Println(email)
}
