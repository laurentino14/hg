package usecases

import (
	"bitbucket.org/elevatt/sirius/internal/domain/entities"
	"context"
	"github.com/gofiber/fiber/v2"
)

type AuthUseCases interface {
	SignInCLIENT(ctx *fiber.Ctx, email string, password string) (*entities.Auth, error)
	SignUpCLIENT(ctx *fiber.Ctx, input entities.NewUserInput) error
	RefreshTokenCLIENT(ctx *fiber.Ctx, rt string) (*entities.Auth, error)
	SignInADMIN(ctx *fiber.Ctx, email string, password string) (*entities.Auth, error)
	FirstAdmin(ctx context.Context) (bool, error)
	RefreshTokenADMIN(ctx *fiber.Ctx, rt string) (*entities.Auth, error)
	SignInWORKER(ctx *fiber.Ctx, email string, password string) (*entities.Auth, error)
	RefreshTokenWORKER(ctx *fiber.Ctx, rt string) (*entities.Auth, error)
}
