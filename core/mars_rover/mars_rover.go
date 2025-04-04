package mars_rover

import (
	"errors"
	"fmt"
	nav "github.com/cameronraw/mars_rover_kata/core/navigation"
)

type MarsRover struct {
	orientation nav.Orientation
}

func NewMarsRover() *MarsRover {
	return &MarsRover{
		orientation: nav.North{},
	}
}

func (m *MarsRover) Execute(commandString string) (string, error) {
	command_runes := []rune(commandString)
	for i := range command_runes {
		if command_runes[i] == 'L' {
			m.turnLeft()
		} else if command_runes[i] == 'R' {
			m.turnRight()
		} else {
			return "", errors.New("invalid command")
		}
	}
	return fmt.Sprintf("0:0:%s", m.orientation), nil
}

func (m *MarsRover) turnLeft() {
	m.orientation = m.orientation.Left()
}

func (m *MarsRover) turnRight() {
	m.orientation = m.orientation.Right()
}
