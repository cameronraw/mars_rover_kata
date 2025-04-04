package models_test

import (
	"github.com/cameronraw/mars_rover_kata/internal/models"
	"testing"
)

func TestMarsRover(t *testing.T) {
	t.Run("should return it's position", func(t *testing.T) {
		// Arrange
		rover := models.NewMarsRover(0, 0, "N")

		// Act
		position := rover.Execute("")

		// Assert
		if position != "0:0:N" {
			t.Errorf("expected position to be '0:0:N', got '%s'", position)
		}

	})

}
