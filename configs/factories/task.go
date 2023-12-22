package factories

import (
	"bitbucket.org/elevatt/sirius/internal/data/services"
	"bitbucket.org/elevatt/sirius/internal/infra/repositories/cockroach"
	"bitbucket.org/elevatt/sirius/internal/presentation/controllers"
	"database/sql"
)

func GenerateTaskController(db *sql.DB) *controllers.TaskController {
	repo := cockroach.GenerateTaskRepository(db)
	service := services.GenerateTaskService(repo)
	return controllers.GenerateTaskController(*service)
}
