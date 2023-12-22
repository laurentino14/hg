package contracts

import (
	"bitbucket.org/elevatt/sirius/internal/domain/entities"
	"context"
)

type OsContract = entities.OS

var NewOsContract = entities.NewOS

type NewOsInputContract = entities.NewOSInput

type UpdateOsInputContract = entities.UpdateOSInput

type OsRepositoryContract interface {
	Create(ctx context.Context, input *OsContract) (*OsContract, error)
	FindAll(ctx context.Context) ([]*OsContract, error)
	FindAllByClientID(ctx context.Context, clientID string) ([]*OsContract, error)
	FindAllByWorkerID(ctx context.Context, employeeID string) ([]*OsContract, error)
	FindAllByProjectCOD(ctx context.Context, projectCOD int) ([]*OsContract, error)
	FindByCOD(ctx context.Context, cod int) (*OsContract, error)
	Update(ctx context.Context, input *UpdateOsInputContract) (*OsContract, error)
	Delete(ctx context.Context, cod int) error
}
