package db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

// Engine core object
type Engine struct {
	conn *pgx.Conn
	ctx  context.Context
}

// NewEngine builder
func NewEngine(c *pgx.Conn) Engine {
	return Engine{
		conn: c,
		ctx:  context.Background(),
	}
}

// ping test db function
func (e *Engine) ping() (string, error) {
	row := e.conn.QueryRow(e.ctx, "SELECT 'pong'")

	var res string

	err := row.Scan(
		&res,
	)

	if err != nil {
		return res, err
	}

	return res, nil
}

// Close connection
func (e *Engine) Close() error {
	err := e.conn.Close(e.ctx)

	if err != nil {
		return err
	}

	return nil
}
