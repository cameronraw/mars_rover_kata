package models

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
