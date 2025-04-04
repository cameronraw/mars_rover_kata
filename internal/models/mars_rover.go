package models

type MarsRover struct{}

func (m *MarsRover) Execute(s string) string {
	return "0:0:N"
}

func NewMarsRover(x int, y int, direction string) *MarsRover {
	return &MarsRover{}
}
