package cockroach

import (
	"bitbucket.org/elevatt/sirius/internal/data/contracts"
	"database/sql"
)

type OsRepository struct {
	db *sql.DB
	contracts.OsRepositoryContract
}

func GenerateOsRepository(db *sql.DB) *OsRepository {
	return &OsRepository{db: db}
}
