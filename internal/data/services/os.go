package services

import (
	"bitbucket.org/elevatt/sirius/internal/data/contracts"
	"bitbucket.org/elevatt/sirius/internal/domain/usecases"
	"context"
)

type OsService struct {
	repo contracts.OsRepositoryContract
	usecases.OsUseCases
}

func GenerateOsService(repo contracts.OsRepositoryContract) *OsService {
	return &OsService{repo: repo}
}

func (s *OsService) Create(ctx context.Context, input *contracts.NewOsInputContract) (*contracts.OsContract, error) {
	data, err := contracts.NewOsContract(input)
	if err != nil {
		return nil, err
	}
	return s.repo.Create(ctx, data)
}

func (s *OsService) FindAll(ctx context.Context) ([]*contracts.OsContract, error) {
	return s.repo.FindAll(ctx)
}

func (s *OsService) FindAllByClientID(ctx context.Context, clientID string) ([]*contracts.OsContract, error) {
	return s.repo.FindAllByClientID(ctx, clientID)
}

func (s *OsService) FindAllByWorkerID(ctx context.Context, workerID string) ([]*contracts.OsContract, error) {
	return s.repo.FindAllByWorkerID(ctx, workerID)
}

func (s *OsService) FindAllByProjectCOD(ctx context.Context, projectCOD int) ([]*contracts.OsContract, error) {
	return s.repo.FindAllByProjectCOD(ctx, projectCOD)
}

func (s *OsService) FindByCOD(ctx context.Context, id int) (*contracts.OsContract, error) {
	return s.repo.FindByCOD(ctx, id)
}

func (s *OsService) Update(ctx context.Context, input *contracts.UpdateOsInputContract) (*contracts.OsContract, error) {
	return s.repo.Update(ctx, input)
}

func (s *OsService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
