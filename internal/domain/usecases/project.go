package usecases

import (
	"bitbucket.org/elevatt/sirius/internal/domain/entities"
	"context"
)

type ProjectUseCases interface {
	Create(ctx context.Context, input *entities.NewProjectInput) (*entities.Project, error)
	FindAll(ctx context.Context) ([]*entities.Project, error)
	FindAllByClientID(ctx context.Context, clientID string) ([]*entities.Project, error)
	FindAllByWorkerID(ctx context.Context, workerID string) ([]*entities.Project, error)
	FindByID(ctx context.Context, id string) (*entities.Project, error)
	Update(ctx context.Context, id string, input *entities.NewProjectInput) (*entities.Project, error)
	Delete(ctx context.Context, id string) error
}
