package cockroach

import (
	"bitbucket.org/elevatt/sirius/internal/data/contracts"
	"database/sql"
)

type TaskRepository struct {
	db *sql.DB
	contracts.TaskRepositoryContract
}

func GenerateTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{db: db}
}
