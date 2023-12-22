package services

import (
	"bitbucket.org/elevatt/sirius/internal/data/contracts"
	"bitbucket.org/elevatt/sirius/internal/data/utils"
	"bitbucket.org/elevatt/sirius/internal/domain/usecases"
	"context"
	"errors"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"time"
)

type AuthService struct {
	repo contracts.AuthRepositoryContract
	usecases.AuthUseCases
}

func GenerateAuthService(repo contracts.AuthRepositoryContract) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) SignInCLIENT(ctx *fiber.Ctx, email string, password string) (*contracts.AuthContract, error) {
	exec, err := s.repo.SignInCLIENT(ctx.Context(), email, password)

	if err != nil {
		return nil, err
	}

	AT, err := utils.GenerateAccessToken(*exec)

	if err != nil {
		return nil, err
	}

	RT, err := utils.GenerateRefreshToken(*exec)

	if err != nil {
		return nil, err
	}
	ctx.Cookie(&fiber.Cookie{
		Name:     "rt",
		Value:    "Bearer " + RT,
		HTTPOnly: true,
		Expires:  time.Now().Add(24 * time.Hour * 7),
		Path:     "/",
	})
	ctx.Cookie(&fiber.Cookie{
		Name:     "at",
		Value:    "Bearer " + AT,
		HTTPOnly: true,
		Expires:  time.Now().Add(1 * time.Hour),
		Path:     "/",
	})

	return &contracts.AuthContract{
		User: *exec,
	}, nil
}

func (s *AuthService) SignUpCLIENT(ctx *fiber.Ctx, input contracts.NewUserInputContract) (bool, error) {
	user, err := contracts.NewUserContract(&input)
	if err != nil {
		return false, err
	}
	return s.repo.SignUpCLIENT(ctx.Context(), *user)
}

func (s *AuthService) RefreshTokenCLIENT(ctx *fiber.Ctx, rt string) (*contracts.AuthContract, error) {
	bearer := "Bearer "
	rt = rt[len(bearer):]

	validate, err := utils.ValidateRT(rt)
	if err != nil {
		return nil, err
	}

	if err != nil || !validate.Valid {
		return nil, errors.New("invalid token")
	}

	customClaim, _ := validate.Claims.(*utils.Claims)

	subject, err := customClaim.RegisteredClaims.GetSubject()
	if err != nil {
		return nil, err
	}

	subjectJSON := utils.Subject{}
	err = json.Unmarshal([]byte(subject), &subjectJSON)
	if err != nil {
		return nil, err
	}

	if subjectJSON.Type != 0 {
		return nil, errors.New("invalid token")
	}

	exec, err := s.repo.GetUserByID(ctx.Context(), subjectJSON.ID)

	if err != nil {
		return nil, err
	}

	AT, err := utils.GenerateAccessToken(*exec)
	if err != nil {
		return nil, err
	}
	RT, err := utils.GenerateRefreshToken(*exec)
	if err != nil {
		return nil, err
	}
	ctx.Cookie(&fiber.Cookie{
		Name:     "rt",
		Value:    "Bearer " + RT,
		HTTPOnly: true,
		Expires:  time.Now().Add(24 * time.Hour * 7),
		Path:     "/",
	})
	ctx.Cookie(&fiber.Cookie{
		Name:     "at",
		Value:    "Bearer " + AT,
		HTTPOnly: true,
		Expires:  time.Now().Add(1 * time.Hour),
		Path:     "/",
	})
	return &contracts.AuthContract{
		User: *exec,
	}, nil

}

func (s *AuthService) SignInADMIN(ctx *fiber.Ctx, email string, password string) (*contracts.AuthContract, error) {
	exec, err := s.repo.SignInADMIN(ctx.Context(), email, password)

	if err != nil {
		return nil, err
	}

	AT, err := utils.GenerateAccessToken(*exec)

	if err != nil {
		return nil, err
	}

	RT, err := utils.GenerateRefreshToken(*exec)

	if err != nil {
		return nil, err
	}
	ctx.Cookie(&fiber.Cookie{
		Name:     "rt",
		Value:    "Bearer " + RT,
		HTTPOnly: true,
		Expires:  time.Now().Add(24 * time.Hour * 7),
		Path:     "/",
	})
	ctx.Cookie(&fiber.Cookie{
		Name:     "at",
		Value:    "Bearer " + AT,
		HTTPOnly: true,
		Expires:  time.Now().Add(1 * time.Hour),
		Path:     "/",
	})

	return &contracts.AuthContract{
		User: *exec,
	}, nil
}

func (s *AuthService) FirstAdmin(ctx context.Context) (bool, error) {
	return s.repo.FirstAdmin(ctx)
}

func (s *AuthService) RefreshTokenADMIN(ctx *fiber.Ctx, rt string) (*contracts.AuthContract, error) {

	bearer := "Bearer "
	rt = rt[len(bearer):]
	validate, err := utils.ValidateRT(rt)
	if err != nil {
		return nil, errors.New("invalid token")
	}

	if err != nil || !validate.Valid {
		return nil, errors.New("invalid token")
	}

	customClaim, _ := validate.Claims.(*utils.Claims)

	subject, err := customClaim.RegisteredClaims.GetSubject()
	if err != nil {
		return nil, err
	}

	subjectJSON := utils.Subject{}
	err = json.Unmarshal([]byte(subject), &subjectJSON)
	if err != nil {
		return nil, err
	}

	if subjectJSON.Type != 1 {
		return nil, errors.New("invalid token")
	}

	exec, err := s.repo.GetUserByID(ctx.Context(), subjectJSON.ID)

	if err != nil {
		return nil, err
	}

	AT, err := utils.GenerateAccessToken(*exec)
	if err != nil {
		return nil, err
	}
	RT, err := utils.GenerateRefreshToken(*exec)
	if err != nil {
		return nil, err
	}

	ctx.Cookie(&fiber.Cookie{
		Name:     "rt",
		Value:    "Bearer " + RT,
		HTTPOnly: true,
		Expires:  time.Now().Add(24 * time.Hour * 7),
		Path:     "/",
	})
	ctx.Cookie(&fiber.Cookie{
		Name:     "at",
		Value:    "Bearer " + AT,
		HTTPOnly: true,
		Expires:  time.Now().Add(1 * time.Hour),
		Path:     "/",
	})
	return &contracts.AuthContract{
		User: *exec,
	}, nil

}

func (s *AuthService) SignInWORKER(ctx *fiber.Ctx, email string, password string) (*contracts.AuthContract, error) {
	exec, err := s.repo.SignInCLIENT(ctx.Context(), email, password)

	if err != nil {
		return nil, err
	}

	AT, err := utils.GenerateAccessToken(*exec)

	if err != nil {
		return nil, err
	}

	RT, err := utils.GenerateRefreshToken(*exec)

	if err != nil {
		return nil, err
	}
	ctx.Cookie(&fiber.Cookie{
		Name:     "rt",
		Value:    "Bearer " + RT,
		HTTPOnly: true,
		Expires:  time.Now().Add(24 * time.Hour * 7),
		Path:     "/",
	})
	ctx.Cookie(&fiber.Cookie{
		Name:     "at",
		Value:    "Bearer " + AT,
		HTTPOnly: true,
		Expires:  time.Now().Add(1 * time.Hour),
		Path:     "/",
	})

	return &contracts.AuthContract{
		User: *exec,
	}, nil
}

func (s *AuthService) RefreshTokenWORKER(ctx *fiber.Ctx, rt string) (*contracts.AuthContract, error) {
	bearer := "Bearer "
	rt = rt[len(bearer):]

	validate, err := utils.ValidateRT(rt)
	if err != nil {
		return nil, err
	}

	if err != nil || !validate.Valid {
		return nil, errors.New("invalid token")
	}

	customClaim, _ := validate.Claims.(*utils.Claims)

	subject, err := customClaim.RegisteredClaims.GetSubject()
	if err != nil {
		return nil, err
	}

	subjectJSON := utils.Subject{}
	err = json.Unmarshal([]byte(subject), &subjectJSON)
	if err != nil {
		return nil, err
	}

	if subjectJSON.Type != 2 {
		return nil, errors.New("invalid token")
	}

	exec, err := s.repo.GetUserByID(ctx.Context(), subjectJSON.ID)

	if err != nil {
		return nil, err
	}

	AT, err := utils.GenerateAccessToken(*exec)
	if err != nil {
		return nil, err
	}
	RT, err := utils.GenerateRefreshToken(*exec)
	if err != nil {
		return nil, err
	}

	ctx.Cookie(&fiber.Cookie{
		Name:     "rt",
		Value:    "Bearer " + RT,
		HTTPOnly: true,
		Expires:  time.Now().Add(24 * time.Hour * 7),
		Path:     "/",
	})
	ctx.Cookie(&fiber.Cookie{
		Name:     "at",
		Value:    "Bearer " + AT,
		HTTPOnly: true,
		Expires:  time.Now().Add(1 * time.Hour),
		Path:     "/",
	})
	return &contracts.AuthContract{
		User: *exec,
	}, nil
}

func (s *AuthService) GetUserByID(ctx context.Context, id string) (*contracts.UserContract, error) {
	return s.repo.GetUserByID(ctx, id)
}
