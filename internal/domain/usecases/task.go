package usecases

import (
	"bitbucket.org/elevatt/sirius/internal/domain/entities"
	"context"
)

type TaskUseCases interface {
	Create(ctx context.Context, input *entities.NewTaskInput) (*entities.Task, error)
	FindAll(ctx context.Context) ([]*entities.Task, error)
	FindAllByProjectCOD(ctx context.Context, projectCOD int) ([]*entities.Task, error)
	FindAllByWorkerID(ctx context.Context, workerID string) ([]*entities.Task, error)
	FindByCOD(ctx context.Context, cod int) (*entities.Task, error)
	Update(ctx context.Context, input *entities.UpdateTaskInput) (*entities.Task, error)
	Delete(ctx context.Context, cod int) error
}
