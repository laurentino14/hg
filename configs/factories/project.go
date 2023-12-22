package factories

import (
	"bitbucket.org/elevatt/sirius/internal/data/services"
	"bitbucket.org/elevatt/sirius/internal/infra/repositories/cockroach"
	"bitbucket.org/elevatt/sirius/internal/presentation/controllers"
	"database/sql"
)

func GenerateProjectController(db *sql.DB) *controllers.ProjectController {
	repo := cockroach.GenerateProjectRepository(db)
	service := services.GenerateProjectService(repo)
	return controllers.GenerateProjectController(*service)
}
