package services

import (
	"bitbucket.org/elevatt/sirius/internal/data/contracts"
	"bitbucket.org/elevatt/sirius/internal/data/utils"
	"errors"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"time"
)

func (s *AuthServiceTestSuite) TestRefreshTokenADMINShouldReturnSuccess() {
	// Arrange
	password := faker.Password()
	user, _ := contracts.NewUserContract(&contracts.NewUserInputContract{
		Name:            faker.Name(),
		Email:           faker.Email(),
		Password:        password,
		ConfirmPassword: password,
		BirthDate:       time.Now(),
		Type:            1,
		Doc:             "12345678900",
		Phone:           "21900000000",
	})
	tokenGen, _ := utils.GenerateRefreshToken(*user)
	token := "Bearer " + tokenGen
	s.AuthRepository.EXPECT().GetUserByID(gomock.Any(), user.ID).Return(user, nil).Times(1)

	// Act
	resp, err := s.AuthService.RefreshTokenADMIN(s.ctx, token)

	// Assert
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), resp.User.ID, user.ID)
}

func (s *AuthServiceTestSuite) TestRefreshTokenADMINShouldReturnError() {
	// Arrange
	password := faker.Password()
	user, _ := contracts.NewUserContract(&contracts.NewUserInputContract{
		Name:            faker.Name(),
		Email:           faker.Email(),
		Password:        password,
		ConfirmPassword: password,
		BirthDate:       time.Now(),
		Type:            1,
		Doc:             "12345678900",
		Phone:           "21900000000",
	})
	tokenGen, _ := utils.GenerateRefreshToken(*user)
	token := "aisnduyasbuydbasud" + tokenGen + "a"
	errExp := errors.New("invalid token")

	// Act
	resp, err := s.AuthService.RefreshTokenADMIN(s.ctx, token)

	// Assert
	assert.Equal(s.T(), err, errExp)
	assert.Nil(s.T(), resp)
}
