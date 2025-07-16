package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Service interface {
	GetTest() (string, error)
}

type PostgresService struct {
	db *sql.DB
}

func NewPostgresService() (*PostgresService, error) {
	connStr := "postgres://postgres:1@localhost:5432/psqldev?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	service := &PostgresService{
		db: db,
	}

	return service, nil
}

func (s *PostgresService) Init() error {
	return nil
}

func (s *PostgresService) GetTest() (string, error) {
	var result string
	rows, err := s.db.Query("select * from test")
	if err != nil {
		return "n-avem", err
	}

	for rows.Next() {
		rows.Scan(&result)
	}

	return result, nil
}
