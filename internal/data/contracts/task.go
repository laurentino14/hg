package contracts

import (
	"bitbucket.org/elevatt/sirius/internal/domain/entities"
	"context"
)

type TaskContract = entities.Task

var NewTaskContract = entities.NewTask

type NewTaskInputContract = entities.NewTaskInput

type UpdateTaskInputContract = entities.UpdateTaskInput

type TaskRepositoryContract interface {
	Create(ctx context.Context, input *TaskContract) (*TaskContract, error)
	FindAll(ctx context.Context) ([]*TaskContract, error)
	FindAllByProjectCOD(ctx context.Context, projectCOD int) ([]*TaskContract, error)
	FindAllByWorkerID(ctx context.Context, workerID string) ([]*TaskContract, error)
	FindByCOD(ctx context.Context, cod int) (*TaskContract, error)
	Update(ctx context.Context, input *UpdateTaskInputContract) (*TaskContract, error)
	Delete(ctx context.Context, cod int) error
}
