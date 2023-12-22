package factories

import (
	"bitbucket.org/elevatt/sirius/internal/data/services"
	"bitbucket.org/elevatt/sirius/internal/infra/repositories/cockroach"
	"bitbucket.org/elevatt/sirius/internal/presentation/controllers"
	"database/sql"
)

func GenerateOsController(db *sql.DB) *controllers.OsController {
	repo := cockroach.GenerateOsRepository(db)
	service := services.GenerateOsService(repo)
	return controllers.GenerateOsController(*service)
}
