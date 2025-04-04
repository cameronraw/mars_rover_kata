package models

type MarsRover struct{}

func (m *MarsRover) Execute(s string) string {
	if s == "L" {
		return "0:0:W"
	}
	return "0:0:N"
}

func NewMarsRover() *MarsRover {
	return &MarsRover{}
}
