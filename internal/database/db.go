package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Test struct {
	gorm.Model
	Check string
}

type Service interface {
	GetTest() (string, error)
}

type PostgresService struct {
	db *gorm.DB
}

func NewPostgresService() (*PostgresService, error) {
	dsn := "host=localhost user=postgres password=1 dbname=psqldev port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	service := &PostgresService{
		db: db,
	}

	return service, nil
}

func (s *PostgresService) Init() error {
	s.db.AutoMigrate(&Test{})
	insertTest := &Test{Check: "test"}
	s.db.Create(insertTest)
	return nil
}

func (s *PostgresService) GetTest() (string, error) {
	var result string
	readTest := &Test{}

	s.db.First(readTest)
	result = readTest.Check

	return result, nil
}
