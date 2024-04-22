package repository

import "database/sql"

type StubUserRepository struct {
	database *sql.DB
}

func newStubUserRepository(database *sql.DB) *StubUserRepository {
	return &StubUserRepository{database: database}
}
