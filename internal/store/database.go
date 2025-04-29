package store

import (
	"database/sql"
	"fmt"
	"io/fs"
	"os"

	"github.com/pressly/goose/v3"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func Open() (*sql.DB, error) {
	url := os.Getenv("TURSO_DATABASE_URL")
	token := os.Getenv("TURSO_AUTH_TOKEN")
	connectionString := fmt.Sprintf("%s?authToken=%s", url, token)
	db, err := sql.Open("libsql", connectionString)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	fmt.Println("connected to database...")

	return db, nil
}

func MigrateFS(db *sql.DB, migrationsFS fs.FS, dir string) error {
	goose.SetBaseFS(migrationsFS)
	defer func() {
		goose.SetBaseFS(nil)
	}()

	return Migrate(db, dir)
}

func Migrate(db *sql.DB, dir string) error {
	err := goose.SetDialect("sqlite3")
	if err != nil {
		return fmt.Errorf("error setting dialect: %w", err)
	}

	err = goose.Up(db, dir)
	if err != nil {
		return fmt.Errorf("error running migrations: %w", err)
	}

	return nil
}
