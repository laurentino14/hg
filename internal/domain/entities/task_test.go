package entities

import (
	"github.com/go-faker/faker/v4"
	"testing"
)

func TestNewTask(t *testing.T) {
	t.Run("Should return a new task", func(t *testing.T) {
		input := &NewTaskInput{
			Name:       faker.Name(),
			OwnerID:    faker.UUIDDigit(),
			ProjectCOD: faker.UUIDDigit(),
		}

		task, err := NewTask(input)

		if err != nil {
			t.Error("Error creating task")
		}

		if task.Name != input.Name {
			t.Errorf("Expected %s, got %s", input.Name, task.Name)
		}

		if task.OwnerID != input.OwnerID {
			t.Errorf("Expected %s, got %s", input.OwnerID, task.OwnerID)
		}

		if task.ProjectCOD != input.ProjectCOD {
			t.Errorf("Expected %s, got %s", input.ProjectCOD, task.ProjectCOD)
		}

		if task.CreatedAt.IsZero() {
			t.Error("CreatedAt is zero")
		}

		if task.UpdatedAt != nil {
			t.Error("UpdatedAt is not empty")
		}

		if task.DeletedAt != nil {
			t.Error("DeletedAt is not empty")
		}
	})
}
