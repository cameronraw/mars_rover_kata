package mars_rover

import (
	"errors"
	"fmt"
)

type MarsRover struct {
	orientation rune
}

func NewMarsRover() *MarsRover {
	return &MarsRover{
		orientation: 'N',
	}
}

func (m *MarsRover) Execute(commandString string) (string, error) {
	command_runes := []rune(commandString)
	for i := range command_runes {
		if command_runes[i] == 'L' {
			err := m.turnLeft()
			if err != nil {
				return "", err
			}
		} else if command_runes[i] == 'R' {
			err := m.turnRight()
			if err != nil {
				return "", err
			}
		} else {
			return "", errors.New("invalid command")
		}
	}
	return fmt.Sprintf("0:0:%s", string(m.orientation)), nil
}

func (m *MarsRover) turnLeft() error {
	switch m.orientation {
	case 'N':
		m.orientation = 'W'
		return nil
	case 'W':
		m.orientation = 'S'
		return nil
	case 'S':
		m.orientation = 'E'
		return nil
	case 'E':
		m.orientation = 'N'
		return nil
	default:
		return errors.New("invalid orientation")
	}
}

func (m *MarsRover) turnRight() error {
	switch m.orientation {
	case 'N':
		m.orientation = 'E'
		return nil
	case 'E':
		m.orientation = 'S'
		return nil
	case 'S':
		m.orientation = 'W'
		return nil
	case 'W':
		m.orientation = 'N'
		return nil
	default:
		return errors.New("invalid orientation")
	}
}
