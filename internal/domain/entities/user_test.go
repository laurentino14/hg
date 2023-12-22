package entities

import (
	"github.com/go-faker/faker/v4"
	"testing"
	"time"
)

func TestNewUser(t *testing.T) {
	t.Run("Should return a new user", func(t *testing.T) {
		Type, err := faker.RandomInt(1, 2)
		if err != nil {
			t.Error("Type is empty")
		}

		password := faker.Password()

		input := &NewUserInput{
			Name:            faker.Name(),
			Email:           faker.Email(),
			Password:        password,
			ConfirmPassword: password,
			BirthDate:       time.Now(),
			Doc:             "18067792771",
			Phone:           faker.Phonenumber(),
			Type:            int8(Type[0]),
		}

		user, err := NewUser(input)

		if err != nil {
			t.Error("Error creating user")
		}

		if user.ID == "" {
			t.Errorf("ID is empty")
		}
		if user.Name != input.Name {
			t.Errorf("Expected %s, got %s", input.Name, user.Name)
		}
		if user.Email != input.Email {
			t.Errorf("Expected %s, got %s", input.Email, user.Email)
		}
		if user.Password != input.Password {
			t.Errorf("Expected %s, got %s", input.Password, user.Password)
		}
		if user.BirthDate != input.BirthDate {
			t.Errorf("Expected %s, got %s", input.BirthDate, user.BirthDate)
		}
		if user.Type != int8(Type[0]) {
			t.Errorf("Expected %d, got %d", input.Type, user.Type)
		}
		if user.Doc != input.Doc {
			t.Errorf("Expected %s, got %s", input.Doc, user.Doc)
		}
		if user.Phone != input.Phone {
			t.Errorf("Expected %s, got %s", input.Phone, user.Phone)
		}
		if user.CreatedAt.IsZero() {
			t.Error("CreatedAt is zero")
		}
		if user.UpdatedAt != nil {
			t.Error("UpdatedAt is not empty")
		}
		if user.DeletedAt != nil {
			t.Error("DeletedAt is not empty")
		}
	})

	t.Run("Should return an error when password and confirm password are different", func(t *testing.T) {
		input := &NewUserInput{
			Name:            faker.Name(),
			Email:           faker.Email(),
			Password:        faker.Password(),
			ConfirmPassword: faker.Password(),
			BirthDate:       time.Now(),
			Doc:             "18067792771",
			Phone:           faker.Phonenumber(),
			Type:            1,
		}

		_, err := NewUser(input)

		if err == nil {
			t.Error("Error creating user")
		}

	})

}
