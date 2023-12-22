package cockroach

import (
	"bitbucket.org/elevatt/sirius/internal/data/contracts"
	"database/sql"
)

type ProjectRepository struct {
	db *sql.DB
	contracts.ProjectRepositoryContract
}

func GenerateProjectRepository(db *sql.DB) *ProjectRepository {
	return &ProjectRepository{db: db}
}
