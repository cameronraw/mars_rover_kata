package navigation

import "fmt"

type Orientation interface {
	Left() Orientation
	Right() Orientation
	fmt.Stringer
}

type North struct{}
type West struct{}
type East struct{}
type South struct{}

func (n North) Left() Orientation {
	return West{}
}

func (n North) Right() Orientation {
	return East{}
}

func (n North) String() string {
	return "N"
}

func (w West) Left() Orientation {
	return South{}
}

func (w West) Right() Orientation {
	return North{}
}

func (w West) String() string {
	return "W"
}

func (e East) Left() Orientation {
	return North{}
}

func (e East) Right() Orientation {
	return South{}
}

func (e East) String() string {
	return "E"
}

func (s South) Left() Orientation {
	return East{}
}

func (s South) Right() Orientation {
	return West{}
}

func (s South) String() string {
	return "S"
}
