package store

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func Open() (*sql.DB, error) {
	db, err := sql.Open("pgx", "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable")
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	log.Println("database connected")

	return db, nil
}
