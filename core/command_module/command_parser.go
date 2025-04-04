package commandmodule

import (
	m "github.com/cameronraw/mars_rover_kata/core/mars_rover"
)

type CanParseCommand interface {
	ParseCommand(command string, rover m.IsRover) Command
}

type Command interface {
	Execute()
}

type CommandParser struct {
	rover *m.MarsRover
}

func NewCommandParser() *CommandParser {
	return &CommandParser{}
}

func (cp *CommandParser) ParseCommand(command string, rover m.IsRover) Command {
	switch command {
	case "L":
		return NewLeftCommand(rover)
	case "R":
		return NewRightCommand(rover)
	default:
		return nil
	}
}

type LeftCommand struct {
	rover m.IsRover
}
type RightCommand struct {
	rover m.IsRover
}

func NewLeftCommand(rover m.IsRover) LeftCommand {
	return LeftCommand{
		rover: rover,
	}
}

func NewRightCommand(rover m.IsRover) RightCommand {
	return RightCommand{
		rover: rover,
	}
}

func (c LeftCommand) Execute() {
	c.rover.TurnLeft()
}

func (c RightCommand) Execute() {
	c.rover.TurnRight()
}
