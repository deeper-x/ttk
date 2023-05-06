package db

import (
	"context"
	"testing"

	"github.com/deeper-x/ttk/settings"
	"github.com/jackc/pgx/v5"
)

func TestPing(t *testing.T) {
	URI, _ := settings.GetDBURI()
	conn, err := pgx.Connect(context.Background(), URI)

	if err != nil {
		t.Error(err)
	}

	e := NewEngine(conn)
	defer e.Close()

	res, err := e.ping()
	if err != nil {
		t.Error(err)
	}

	expected := "pong"

	if res != expected {
		t.Errorf("expected %v, got %v", expected, res)
	}
}
