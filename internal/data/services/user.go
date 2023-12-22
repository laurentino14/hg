package services

import (
	"bitbucket.org/elevatt/sirius/internal/data/contracts"
	"bitbucket.org/elevatt/sirius/internal/domain/usecases"
	"context"
	"errors"
)

type UserService struct {
	repo contracts.UserRepositoryContract
	usecases.UserUseCases
}

func GenerateUserService(repo contracts.UserRepositoryContract) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(ctx context.Context, input *contracts.NewUserInputContract) (*contracts.UserContract, error) {
	data, err := contracts.NewUserContract(input)
	if err != nil {
		return nil, err
	}
	return s.repo.Create(ctx, data)
}

func (s *UserService) FindAllClients(ctx context.Context) ([]*contracts.UserContract, error) {
	return s.repo.FindAllClients(ctx)
}

func (s *UserService) FindAllEmployees(ctx context.Context) ([]*contracts.UserContract, error) {
	return s.repo.FindAllEmployees(ctx)
}

func (s *UserService) FindAllAdmins(ctx context.Context) ([]*contracts.UserContract, error) {
	return s.repo.FindAllAdmins(ctx)
}

func (s *UserService) FindByID(ctx context.Context, id string) (*contracts.UserContract, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *UserService) Update(ctx context.Context, input *contracts.NewUserInputContract) (*contracts.UserContract, error) {
	return s.repo.Update(ctx, input)
}

func (s *UserService) UpdatePassword(ctx context.Context, input *contracts.UpdateUserPasswordContract) (*contracts.UserContract, error) {
	if input.Password != input.ConfirmPassword {
		return nil, errors.New("passwords do not match")
	}

	return s.repo.UpdatePassword(ctx, input)
}

func (s *UserService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
