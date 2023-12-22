package services

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func (s *AuthServiceTestSuite) TestFirstAdminShouldReturnSuccess() {
	// Arrange
	s.AuthRepository.EXPECT().FirstAdmin(gomock.Any()).Return(true, nil)
	// Act
	resp, err := s.AuthService.FirstAdmin(context.Background())

	// Assert
	assert.Nil(s.T(), err)
	assert.True(s.T(), resp)
}

func (s *AuthServiceTestSuite) TestFirstAdminShouldReturnError() {
	// Arrange
	expError := errors.New("error")
	s.AuthRepository.EXPECT().FirstAdmin(gomock.Any()).Return(false, expError).Times(1)
	// Act
	resp, err := s.AuthService.FirstAdmin(context.Background())

	// Assert
	assert.Errorf(s.T(), err, expError.Error())
	assert.False(s.T(), resp)
}
