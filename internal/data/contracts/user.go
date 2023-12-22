package contracts

import (
	"bitbucket.org/elevatt/sirius/internal/domain/entities"
	"context"
)

type UserContract = entities.User

var NewUserContract = entities.NewUser

type NewUserInputContract = entities.NewUserInput

type UpdateUserContract = entities.UpdateUserInput

type UpdateUserPasswordContract = entities.UpdateUserPasswordInput

type UserRepositoryContract interface {
	Create(ctx context.Context, input *UserContract) (*UserContract, error)
	FindAllClients(ctx context.Context) ([]*UserContract, error)
	FindAllEmployees(ctx context.Context) ([]*UserContract, error)
	FindAllAdmins(ctx context.Context) ([]*UserContract, error)
	FindByID(ctx context.Context, id string) (*UserContract, error)
	Update(ctx context.Context, input *NewUserInputContract) (*UserContract, error)
	UpdatePassword(ctx context.Context, input *UpdateUserPasswordContract) (*UserContract, error)
	Delete(ctx context.Context, id string) error
}
