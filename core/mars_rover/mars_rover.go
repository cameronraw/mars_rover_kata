package mars_rover

import (
	"errors"
	"fmt"

	nav "github.com/cameronraw/mars_rover_kata/core/navigation"
)

type MarsRover struct {
	orientation nav.Orientation
	nav.Coordinates
}

type Position struct {
	X int
	Y int
}

type IsRover interface {
	TurnLeft()
	TurnRight()
	MoveForward()
}

func NewMarsRover() *MarsRover {
	return &MarsRover{
		orientation: &nav.North{},
		Coordinates: nav.Coordinates{X: 0, Y: 0},
	}
}

func (m *MarsRover) Execute(commandString string) (string, error) {
	command_runes := []rune(commandString)
	for i := range command_runes {
		switch command_runes[i] {
		case 'L':
			m.TurnLeft()
		case 'R':
			m.TurnRight()
		case 'M':
			m.MoveForward()
		default:
			return "", errors.New("invalid command")
		}
	}
	return fmt.Sprintf("%d:%d:%s", m.X, m.Y, m.orientation), nil
}

func (m *MarsRover) TurnLeft() {
	m.orientation = m.orientation.Left()
}

func (m *MarsRover) TurnRight() {
	m.orientation = m.orientation.Right()
}

func (m *MarsRover) MoveForward() {
	m.orientation.MoveForward(&m.Coordinates)
}
