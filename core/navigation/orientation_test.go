package navigation_test

import (
	"github.com/cameronraw/mars_rover_kata/core/navigation"
	"testing"
)

func TestNorthLeft(t *testing.T) {
	orientation := navigation.North{}
	expected := navigation.West{}
	actual := orientation.Left()

	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}

func TestNorthRight(t *testing.T) {
	orientation := navigation.North{}
	expected := navigation.East{}
	actual := orientation.Right()

	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}

func TestEastLeft(t *testing.T) {
	orientation := navigation.East{}
	expected := navigation.North{}
	actual := orientation.Left()

	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}

func TestEastRight(t *testing.T) {
	orientation := navigation.East{}
	expected := navigation.South{}
	actual := orientation.Right()

	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}

func TestSouthLeft(t *testing.T) {
	orientation := navigation.South{}
	expected := navigation.East{}
	actual := orientation.Left()

	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}

func TestSouthRight(t *testing.T) {
	orientation := navigation.South{}
	expected := navigation.West{}
	actual := orientation.Right()

	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}

func TestWestLeft(t *testing.T) {
	orientation := navigation.West{}
	expected := navigation.South{}
	actual := orientation.Left()

	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}

func TestWestRight(t *testing.T) {
	orientation := navigation.West{}
	expected := navigation.North{}
	actual := orientation.Right()

	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}
