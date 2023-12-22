package factories

import (
	"bitbucket.org/elevatt/sirius/internal/data/services"
	"bitbucket.org/elevatt/sirius/internal/infra/repositories/cockroach"
	"bitbucket.org/elevatt/sirius/internal/presentation/controllers"
	"database/sql"
)

func GenerateAuthController(db *sql.DB) *controllers.AuthController {
	repo := cockroach.GenerateAuthRepository(db)
	service := services.GenerateAuthService(repo)
	return controllers.GenerateAuthController(*service)
}
