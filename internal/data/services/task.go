package services

import (
	"bitbucket.org/elevatt/sirius/internal/data/contracts"
	"bitbucket.org/elevatt/sirius/internal/domain/usecases"
	"context"
)

type TaskService struct {
	repo contracts.TaskRepositoryContract
	usecases.TaskUseCases
}

func GenerateTaskService(repo contracts.TaskRepositoryContract) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) Create(ctx context.Context, input *contracts.NewTaskInputContract) (*contracts.TaskContract, error) {
	data, err := contracts.NewTaskContract(input)
	if err != nil {
		return nil, err
	}
	return s.repo.Create(ctx, data)
}

func (s *TaskService) FindAll(ctx context.Context) ([]*contracts.TaskContract, error) {
	return s.repo.FindAll(ctx)
}

func (s *TaskService) FindAllByProjectCOD(ctx context.Context, projectCOD int) ([]*contracts.TaskContract, error) {
	return s.repo.FindAllByProjectCOD(ctx, projectCOD)
}

func (s *TaskService) FindAllByWorkerID(ctx context.Context, workerID string) ([]*contracts.TaskContract, error) {
	return s.repo.FindAllByWorkerID(ctx, workerID)
}

func (s *TaskService) FindByCOD(ctx context.Context, cod int) (*contracts.TaskContract, error) {
	return s.repo.FindByCOD(ctx, cod)
}

func (s *TaskService) Update(ctx context.Context, input *contracts.UpdateTaskInputContract) (*contracts.TaskContract, error) {
	return s.repo.Update(ctx, input)
}

func (s *TaskService) Delete(ctx context.Context, cod int) error {
	return s.repo.Delete(ctx, cod)
}
