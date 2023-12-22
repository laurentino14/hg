package services

import (
	mockinfra "bitbucket.org/elevatt/sirius/internal/infra/mocks"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/suite"
	"github.com/valyala/fasthttp"
	"go.uber.org/mock/gomock"
	"testing"
)

type AuthServiceTestSuite struct {
	suite.Suite
	ctrl           *gomock.Controller
	ctx            *fiber.Ctx
	AuthRepository *mockinfra.MockAuthRepositoryContract

	AuthService *AuthService
}

func TestAuthServiceTestSuite(t *testing.T) {
	suite.Run(t, new(AuthServiceTestSuite))
}

func (s *AuthServiceTestSuite) SetupTest() {
	app := fiber.New()
	c := app.AcquireCtx(&fasthttp.RequestCtx{})
	s.ctx = c
	s.ctrl = gomock.NewController(s.T())

	// Mocks
	s.AuthRepository = mockinfra.NewMockAuthRepositoryContract(s.ctrl)
	s.AuthService = GenerateAuthService(s.AuthRepository)
}

//func (s *AuthServiceTestSuite) TestAuthServiceSuiteDown() {
//	defer s.ctrl.Finish()
//}
