package entities

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        string     `json:"user_id"`
	Name      string     `json:"name"`
	Email     string     `json:"email""`
	Password  string     `json:"-"`
	BirthDate time.Time  `json:"birth_date"`
	Type      int8       `json:"type"`
	Doc       string     `json:"doc"`
	Phone     string     `json:"phone"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type NewUserInput struct {
	Name            string    `json:"name"`
	Email           string    `json:"email"`
	Password        string    `json:"password"`
	ConfirmPassword string    `json:"confirm_password"`
	BirthDate       time.Time `json:"birth_date"`
	Type            int8      `json:"type"`
	Doc             string    `json:"doc"`
	Phone           string    `json:"phone"`
}

type UpdateUserInput struct {
	ID        string    `json:"user_id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	BirthDate time.Time `json:"birth_date"`
	Doc       string    `json:"doc"`
	Phone     string    `json:"phone"`
}

type UpdateUserPasswordInput struct {
	ID              string `json:"user_id"`
	OldPassword     string `json:"old_password"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type AssociateUserInput struct {
	ID        string     `json:"user_id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	BirthDate time.Time  `json:"birth_date"`
	Type      int8       `json:"type"`
	Doc       string     `json:"doc"`
	Phone     string     `json:"phone"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func NewUser(input *NewUserInput) (*User, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, errors.New("error generating UUID")
	}
	user := &User{}

	if input.Password != input.ConfirmPassword {
		return nil, errors.New("passwords do not match")
	}

	user.ID = id.String()
	user.Name = input.Name
	user.Email = input.Email
	user.Password = input.Password
	user.BirthDate = input.BirthDate
	user.Type = input.Type
	user.Doc = input.Doc
	user.Phone = input.Phone
	user.CreatedAt = time.Now()
	user.UpdatedAt = nil
	user.DeletedAt = nil

	return user, nil
}

func AssociateUser(input *AssociateUserInput) *User {
	user := &User{}

	user.ID = input.ID
	user.Name = input.Name
	user.Email = input.Email
	user.Password = input.Password
	user.BirthDate = input.BirthDate
	user.Type = input.Type
	user.Doc = input.Doc
	user.Phone = input.Phone
	user.CreatedAt = input.CreatedAt
	user.UpdatedAt = input.UpdatedAt
	user.DeletedAt = input.DeletedAt

	return user
}
