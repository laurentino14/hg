package services

import (
	"bitbucket.org/elevatt/sirius/internal/data/contracts"
	"bitbucket.org/elevatt/sirius/internal/domain/usecases"
	"context"
)

type ProjectService struct {
	repo contracts.ProjectRepositoryContract
	usecases.ProjectUseCases
}

func GenerateProjectService(repo contracts.ProjectRepositoryContract) *ProjectService {
	return &ProjectService{repo: repo}
}

func (s *ProjectService) Create(ctx context.Context, input *contracts.NewProjectInputContract) (*contracts.ProjectContract, error) {
	data, err := contracts.NewProjectContract(input)
	if err != nil {
		return nil, err
	}
	return s.repo.Create(ctx, data)
}

func (s *ProjectService) FindAll(ctx context.Context) ([]*contracts.ProjectContract, error) {
	return s.repo.FindAll(ctx)
}

func (s *ProjectService) FindAllByClientID(ctx context.Context, clientID string) ([]*contracts.ProjectContract, error) {
	return s.repo.FindAllByClientID(ctx, clientID)
}

func (s *ProjectService) FindAllByWorkerID(ctx context.Context, workerID string) ([]*contracts.ProjectContract, error) {
	return s.repo.FindAllByWorkerID(ctx, workerID)
}

func (s *ProjectService) FindByID(ctx context.Context, id string) (*contracts.ProjectContract, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *ProjectService) Update(ctx context.Context, input *contracts.UpdateProjectInputContract) (*contracts.ProjectContract, error) {
	return s.repo.Update(ctx, input)
}

func (s *ProjectService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
