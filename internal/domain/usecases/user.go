package usecases

import (
	"bitbucket.org/elevatt/sirius/internal/domain/entities"
	"context"
)

type UserUseCases interface {
	Create(ctx context.Context, input *entities.NewUserInput) (*entities.User, error)
	FindAllClients(ctx context.Context) ([]*entities.User, error)
	FindAllEmployees(ctx context.Context) ([]*entities.User, error)
	FindAllAdmins(ctx context.Context) ([]*entities.User, error)
	FindByID(ctx context.Context, id string) (*entities.User, error)
	Update(ctx context.Context, input *entities.UpdateUserInput) (*entities.User, error)
	UpdatePassword(ctx context.Context, input *entities.UpdateUserPasswordInput) (*entities.User, error)
	Delete(ctx context.Context, id string) error
}
