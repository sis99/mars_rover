package main

import (
	"errors"
	"testing"
)

var rightTurn = map[Direction]Direction{
	North: East,
	East:  South,
	South: West,
	West:  North,
}

var leftTurn = map[Direction]Direction{
	North: West,
	West:  South,
	South: East,
	East:  North,
}

func TestRover_TurnRight(t *testing.T) {
	for _, dir := range allDirections {
		t.Run(string(dir), func(t *testing.T) {
			rover := &Rover{X: 1, Y: 1, Direction: dir}
			rover.TurnRight()
			if rover.Direction != rightTurn[dir] {
				t.Errorf("TurnRight from %s: got %s, expected %s", dir, rover.Direction, rightTurn[dir])
			}
		})
	}
}

func TestRover_TurnLeft(t *testing.T) {
	for _, dir := range allDirections {
		t.Run(string(dir), func(t *testing.T) {
			rover := &Rover{X: 1, Y: 1, Direction: dir}
			rover.TurnLeft()
			if rover.Direction != leftTurn[dir] {
				t.Errorf("TurnLeft from %s: got %s, expected %s", dir, rover.Direction, leftTurn[dir])
			}
		})
	}
}

func TestRover_Move(t *testing.T) {
	plateau := checkPlateau(NewPlateau(2, 2))
	tests := []struct {
		direction            Direction
		expectedX, expectedY int
	}{
		{North, 1, 2},
		{East, 2, 1},
		{South, 1, 0},
		{West, 0, 1},
	}
	for _, tt := range tests {
		t.Run(string(tt.direction), func(t *testing.T) {
			rover := &Rover{X: 1, Y: 1, Direction: tt.direction}
			if err := rover.Move(plateau); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if rover.X != tt.expectedX || rover.Y != tt.expectedY {
				t.Errorf("got (%d, %d), expected (%d, %d)", rover.X, rover.Y, tt.expectedX, tt.expectedY)
			}
		})
	}
}

func TestRover_Move_OutOfBounds(t *testing.T) {
	plateau := checkPlateau(NewPlateau(1, 1))
	rover := &Rover{X: 1, Y: 1, Direction: North}
	if err := rover.Move(plateau); !errors.Is(err, ErrInvalidMove) {
		t.Errorf("expected ErrInvalidMove, got %v", err)
	}
}
