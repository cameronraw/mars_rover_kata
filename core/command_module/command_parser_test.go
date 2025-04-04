package commandmodule_test

import (
	commandmodule "github.com/cameronraw/mars_rover_kata/core/command_module"
	"testing"
)

type MockRover struct {
	commandParser commandmodule.CanParseCommand
	CalledFuncs   []string
}

func NewMockRover(commandParser *commandmodule.CommandParser) *MockRover {
	return &MockRover{
		commandParser: commandParser,
	}
}

func (m *MockRover) TurnLeft() {
	m.CalledFuncs = append(m.CalledFuncs, "TurnLeft")
}

func (m *MockRover) TurnRight() {
	m.CalledFuncs = append(m.CalledFuncs, "TurnRight")
}

func (m *MockRover) MoveForward() {
	m.CalledFuncs = append(m.CalledFuncs, "MoveForward")
}

func TestCommandParserShould(t *testing.T) {

	t.Run("return LeftCommand when parsing 'L'", func(t *testing.T) {

		commandParser := commandmodule.NewCommandParser()
		mockRover := NewMockRover(commandParser)

		command := commandParser.ParseCommand("L", mockRover)
		command.Execute()

		funcCall := mockRover.CalledFuncs[0]

		if funcCall != "TurnLeft" {
			t.Errorf("expected TurnLeft to be called, got '%s'", funcCall)
		}
	})

	t.Run("return RightCommand when parsing 'R'", func(t *testing.T) {

		commandParser := commandmodule.NewCommandParser()
		mockRover := NewMockRover(commandParser)

		command := commandParser.ParseCommand("R", mockRover)
		command.Execute()

		funcCall := mockRover.CalledFuncs[0]

		if funcCall != "TurnRight" {
			t.Errorf("expected TurnRight to be called, got '%s'", funcCall)
		}
	})
}
