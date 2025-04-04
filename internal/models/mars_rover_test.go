package models_test

import (
	"github.com/cameronraw/mars_rover_kata/internal/models"
	"testing"
)

func TestMarsRoverShould(t *testing.T) {
	t.Run("return it's position", func(t *testing.T) {
		rover := models.NewMarsRover()

		position, _ := rover.Execute("")

		expected := "0:0:N"

		if position != expected {
			t.Errorf("expected position to be '%s', got '%s'", expected, position)
		}
	})
	t.Run("return 0:0:W position after receiving 'L'", func(t *testing.T) {
		rover := models.NewMarsRover()

		position, _ := rover.Execute("L")

		expected := "0:0:W"

		if position != expected {
			t.Errorf("expected position to be '%s', got '%s'", expected, position)
		}
	})
	t.Run("return 0:0:S position after receiving 'LL'", func(t *testing.T) {
		rover := models.NewMarsRover()

		position, _ := rover.Execute("LL")

		expected := "0:0:S"

		if position != expected {
			t.Errorf("expected position to be '%s', got '%s'", expected, position)
		}
	})
	t.Run("return 0:0:E position after receiving 'LLL'", func(t *testing.T) {
		rover := models.NewMarsRover()

		position, _ := rover.Execute("LLL")

		expected := "0:0:E"

		if position != expected {
			t.Errorf("expected position to be '%s', got '%s'", expected, position)
		}
	})
}
