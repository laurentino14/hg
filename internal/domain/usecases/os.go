package usecases

import (
	"bitbucket.org/elevatt/sirius/internal/domain/entities"
	"context"
)

type OsUseCases interface {
	Create(ctx context.Context, input *entities.NewOSInput) (*entities.OS, error)
	FindAll(ctx context.Context) ([]*entities.OS, error)
	FindAllByClientID(ctx context.Context, clientID string) ([]*entities.OS, error)
	FindAllByWorkerID(ctx context.Context, workerID string) ([]*entities.OS, error)
	FindAllByProjectCOD(ctx context.Context, projectCOD int) ([]*entities.OS, error)
	FindByID(ctx context.Context, id int) (*entities.OS, error)
	Update(ctx context.Context, input *entities.UpdateOSInput) (*entities.OS, error)
	Delete(ctx context.Context, id int) error
}
