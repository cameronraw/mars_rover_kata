package mars_rover_test

import (
	"github.com/cameronraw/mars_rover_kata/core/mars_rover"
	"testing"
)

func TestMarsRoverShould(t *testing.T) {
	t.Run("return it's position", func(t *testing.T) {
		rover := mars_rover.NewMarsRover()

		position, _ := rover.Execute("")

		expected := "0:0:N"

		if position != expected {
			t.Errorf("expected position to be '%s', got '%s'", expected, position)
		}
	})
	t.Run("return 0:0:W position after receiving 'L'", func(t *testing.T) {
		rover := mars_rover.NewMarsRover()

		position, _ := rover.Execute("L")

		expected := "0:0:W"

		if position != expected {
			t.Errorf("expected position to be '%s', got '%s'", expected, position)
		}
	})
	t.Run("return 0:0:S position after receiving 'LL'", func(t *testing.T) {
		rover := mars_rover.NewMarsRover()

		position, _ := rover.Execute("LL")

		expected := "0:0:S"

		if position != expected {
			t.Errorf("expected position to be '%s', got '%s'", expected, position)
		}
	})
	t.Run("return 0:0:E position after receiving 'LLL'", func(t *testing.T) {
		rover := mars_rover.NewMarsRover()

		position, _ := rover.Execute("LLL")

		expected := "0:0:E"

		if position != expected {
			t.Errorf("expected position to be '%s', got '%s'", expected, position)
		}
	})
	t.Run("return 0:0:E position after receiving 'R'", func(t *testing.T) {
		rover := mars_rover.NewMarsRover()

		position, _ := rover.Execute("R")

		expected := "0:0:E"

		if position != expected {
			t.Errorf("expected position to be '%s', got '%s'", expected, position)
		}
	})
	t.Run("return 0:0:S position after receiving 'RR'", func(t *testing.T) {
		rover := mars_rover.NewMarsRover()

		position, _ := rover.Execute("RR")

		expected := "0:0:S"

		if position != expected {
			t.Errorf("expected position to be '%s', got '%s'", expected, position)
		}
	})
	t.Run("return 0:0:W position after receiving 'RRR'", func(t *testing.T) {
		rover := mars_rover.NewMarsRover()

		position, _ := rover.Execute("RRR")

		expected := "0:0:W"

		if position != expected {
			t.Errorf("expected position to be '%s', got '%s'", expected, position)
		}
	})
	t.Run("return error when receiving invalid command", func(t *testing.T) {
		rover := mars_rover.NewMarsRover()

		_, err := rover.Execute("X")

		if err == nil {
			t.Errorf("expected an error, got nil")
		}
	})
}
