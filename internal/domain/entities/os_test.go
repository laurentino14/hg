package entities

import "testing"

func TestNewOS(t *testing.T) {
	t.Run("Should return a new OS", func(t *testing.T) {
		input := &NewOSInput{
			Price: 100.00,
		}

		os, err := NewOS(input)
		if err != nil {
			t.Error("Error creating OS")
		}

		if os.Price != input.Price {
			t.Errorf("Expected %f, got %f", input.Price, os.Price)
		}
		if os.CreatedAt.IsZero() {
			t.Error("CreatedAt is zero")
		}

		if os.UpdatedAt != nil {
			t.Error("UpdatedAt is not empty")
		}

		if os.DeletedAt != nil {
			t.Error("DeletedAt is not empty")
		}
	})
}
