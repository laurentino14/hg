package factories

import (
	"bitbucket.org/elevatt/sirius/internal/data/services"
	"bitbucket.org/elevatt/sirius/internal/infra/repositories/cockroach"
	"bitbucket.org/elevatt/sirius/internal/presentation/controllers"
	"database/sql"
)

func GenerateUserController(db *sql.DB) *controllers.UserController {
	repo := cockroach.GenerateUserRepository(db)
	service := services.GenerateUserService(repo)
	return controllers.GenerateUserController(*service)
}
