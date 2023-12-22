package entities

import (
	"github.com/go-faker/faker/v4"
	"testing"
	"time"
)

func TestNewAuth(t *testing.T) {
	t.Run("should create a new auth by new user input", func(t *testing.T) {
		input := new(NewUserInput)
		input.Name = faker.Name()
		input.Email = faker.Email()
		input.Password = faker.Password()
		input.Type = 0
		input.ConfirmPassword = input.Password
		input.BirthDate = time.Now()
		input.Doc = "12345678901"
		input.Phone = "12345678901"

		user, err := NewUser(input)
		if err != nil {
			t.Error(err)
		}

		auth := NewAuth(*user)
		if auth.User != *user {
			t.Error("User not equal")
		}

	})

	t.Run("should create a new auth by associate input", func(t *testing.T) {
		input := new(AssociateUserInput)
		input.Name = faker.Name()
		input.Email = faker.Email()
		input.Password = faker.Password()
		input.Type = 0
		input.BirthDate = time.Now()
		input.Doc = "12345678901"
		input.Phone = "12345678901"
		input.CreatedAt = time.Now()
		user := AssociateUser(input)

		auth := NewAuth(*user)
		if auth.User != *user {
			t.Error("User not equal")
		}

	})

}
