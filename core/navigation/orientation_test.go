package navigation

import (
	"testing"
)

func TestNorthLeft(t *testing.T) {
	orientation := North{}
	expected := &West{}
	actual := orientation.Left()

	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}

func TestNorthRight(t *testing.T) {
	orientation := North{}
	expected := &East{}
	actual := orientation.Right()

	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}

func TestEastLeft(t *testing.T) {
	orientation := East{}
	expected := &North{}
	actual := orientation.Left()

	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}

func TestEastRight(t *testing.T) {
	orientation := East{}
	expected := &South{}
	actual := orientation.Right()

	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}

func TestSouthLeft(t *testing.T) {
	orientation := South{}
	expected := &East{}
	actual := orientation.Left()

	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}

func TestSouthRight(t *testing.T) {
	orientation := South{}
	expected := &West{}
	actual := orientation.Right()

	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}

func TestWestLeft(t *testing.T) {
	orientation := West{}
	expected := &South{}
	actual := orientation.Left()

	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}

func TestWestRight(t *testing.T) {
	orientation := West{}
	expected := &North{}
	actual := orientation.Right()

	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}
