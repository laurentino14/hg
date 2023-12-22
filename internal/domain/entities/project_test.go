package entities

import (
	"github.com/go-faker/faker/v4"
	"testing"
)

func TestNewProject(t *testing.T) {
	t.Run("Should return a new project", func(t *testing.T) {
		input := &NewProjectInput{
			Name:        faker.Name(),
			Description: faker.Paragraph(),
			Type:        1,
			OwnerID:     faker.UUIDDigit(),
		}

		project, err := NewProject(input)
		if err != nil {
			t.Error("Error creating project")
		}

		if project.Name != input.Name {
			t.Errorf("Expected %s, got %s", input.Name, project.Name)
		}
		if project.Description != input.Description {
			t.Errorf("Expected %s, got %s", input.Description, project.Description)
		}
		if project.Type != input.Type {
			t.Errorf("Expected %d, got %d", input.Type, project.Type)
		}
		if project.OwnerID != input.OwnerID {
			t.Errorf("Expected %s, got %s", input.OwnerID, project.OwnerID)
		}
		if project.CreatedAt.IsZero() {
			t.Error("CreatedAt is zero")
		}
		if project.UpdatedAt != nil {
			t.Error("UpdatedAt is not empty")
		}
		if project.DeletedAt != nil {
			t.Error("DeletedAt is not empty")
		}
	})
}
