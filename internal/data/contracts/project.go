package contracts

import (
	"bitbucket.org/elevatt/sirius/internal/domain/entities"
	"context"
)

type ProjectContract = entities.Project

var NewProjectContract = entities.NewProject

type NewProjectInputContract = entities.NewProjectInput

type UpdateProjectInputContract = entities.UpdateProjectInput

type ProjectRepositoryContract interface {
	Create(ctx context.Context, input *ProjectContract) (*ProjectContract, error)
	FindAll(ctx context.Context) ([]*ProjectContract, error)
	FindAllByClientID(ctx context.Context, clientID string) ([]*ProjectContract, error)
	FindAllByWorkerID(ctx context.Context, workerID string) ([]*ProjectContract, error)
	FindByID(ctx context.Context, id string) (*ProjectContract, error)
	Update(ctx context.Context, input *UpdateProjectInputContract) (*ProjectContract, error)
	Delete(ctx context.Context, id string) error
}
