package postgres

import (
	"database/sql"
	"fmt"
	"io"
	"os"

	_ "github.com/lib/pq" // init postgres driver
)

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "postgres"
	dbname   = "postgresProTask"
)

type Storage struct {
	db *sql.DB
}

func New() (*Storage, error) {
	const fn = "storage.postgres.New"
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	initQuery, err := readSQLFile()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	_, err = db.Query(initQuery)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	return &Storage{db: db}, nil
}

func readSQLFile() (string, error) {
	file, err := os.Open("./resources/database/table.sql")
	if err != nil {
		return "", fmt.Errorf("readSQLFile: %w", err)
	}

	b, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("readSQLFile: %w", err)
	}

	return string(b), nil
}

func (s *Storage) InsertIntoStorage(username, password string) error {
	const fn = "storage.postgres.InsertIntoStorage"

	stmt, err := s.db.Prepare("INSERT INTO users(username, password) VALUES ($1, $2)")
	if err != nil {
		return fmt.Errorf("%s: %w", fn, err)
	}

	_, err = stmt.Exec(username, password)
	if err != nil {
		return fmt.Errorf("%s: %w", fn, err)
	}

	return nil
}
