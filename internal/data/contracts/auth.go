package contracts

import (
	"bitbucket.org/elevatt/sirius/internal/domain/entities"
	"context"
)

type AuthContract = entities.Auth

var NewAuthContract = entities.NewAuth

type AuthRepositoryContract interface {
	SignInCLIENT(ctx context.Context, email string, password string) (*UserContract, error)
	SignUpCLIENT(ctx context.Context, input UserContract) (bool, error)
	FirstAdmin(ctx context.Context) (bool, error)
	SignInADMIN(ctx context.Context, email string, password string) (*UserContract, error)
	SignInWORKER(ctx context.Context, email string, password string) (*UserContract, error)
	GetUserByID(ctx context.Context, id string) (*UserContract, error)
}
