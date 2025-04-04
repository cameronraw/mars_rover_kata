package models

type MarsRover struct{}

func (m *MarsRover) Execute(s string) string {
	if s == "L" {
		return "0:0:W"
	}
	if s == "LL" {
		return "0:0:S"
	}
	if s == "LLL" {
		return "0:0:E"
	}
	return "0:0:N"
}

func NewMarsRover() *MarsRover {
	return &MarsRover{}
}
