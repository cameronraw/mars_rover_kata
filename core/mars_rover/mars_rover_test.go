package mars_rover_test

import (
	"testing"

	"github.com/cameronraw/mars_rover_kata/core/mars_rover"
	nav "github.com/cameronraw/mars_rover_kata/core/navigation"
)

func TestMarsRoverShould(t *testing.T) {
	testCases := []struct {
		name     string
		command  string
		expected string
	}{
		{"return it's position", "", "0:0:N"},
		{"return 0:0:W position after receiving 'L'", "L", "0:0:W"},
		{"return 0:0:S position after receiving 'LL'", "LL", "0:0:S"},
		{"return 0:0:E position after receiving 'LLL'", "LLL", "0:0:E"},
		{"return 0:0:E position after receiving 'R'", "R", "0:0:E"},
		{"return 0:0:S position after receiving 'RR'", "RR", "0:0:S"},
		{"return 0:0:W position after receiving 'RRR'", "RRR", "0:0:W"},
		{"return 0:1:N position after receiving 'M'", "M", "0:1:N"},
		{"return 1:0:E position after receiving 'RM'", "RM", "1:0:E"},
		{"return 9:0:W position after receiving a grid of 10,15 and command 'LM'", "LM", "9:0:W"},
		{"return 9:0:W position after receiving a grid of 10,15 and command 'LM'", "LM", "9:0:W"},
		{"return 0:14:S position after receiving a grid of 10,15 and command 'LLM'", "LLM", "0:14:S"},
		{"return 0:1:N position after receiving a grid of 10,15 and command 'MMMMMMMMMMMMMMMM'", "MMMMMMMMMMMMMMMM", "0:1:N"},
		{"return 1:0:E position after receiving a grid of 10,15 and command 'RMMMMMMMMMMM'", "RMMMMMMMMMMM", "1:0:E"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			grid := nav.NewGrid(10, 15)
			rover := mars_rover.NewMarsRover(*grid)

			position, _ := rover.Execute(tc.command)

			if position != tc.expected {
				t.Errorf("expected position to be '%s', got '%s'", tc.expected, position)
			}
		})
	}
}
