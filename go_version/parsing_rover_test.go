package main

import (
	"errors"
	"testing"
)

var invalidRoverLines = []string{
	"0 -1 N",
	"0 0 P",
	"0, 1; N",
	"ekqnfek",
	" 0  -1  N  ",
}

func TestExtractRover_Invalid(t *testing.T) {
	plateau := checkPlateau(NewPlateau(1, 1))
	for _, line := range invalidRoverLines {
		t.Run(line, func(t *testing.T) {
			_, err := extractRover(plateau, line)
			if !errors.Is(err, ErrInvalidRoverParams) {
				t.Errorf("expected ErrInvalidRoverParams, got %v", err)
			}
		})
	}
}

var validRoverLines = []struct {
	line                 string
	expectedX, expectedY int
	expectedDir          Direction
}{
	{"0 -1 N", 0, -1, North},
	{"11 10 S", 11, 10, South},
	{"1 1 W", 1, 1, West},
}

func TestExtractRover_Valid(t *testing.T) {
	plateau := checkPlateau(NewPlateauWithBounds(20, 20, -1, -1))
	for _, tt := range validRoverLines {
		t.Run(tt.line, func(t *testing.T) {
			rover, err := extractRover(plateau, tt.line)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if rover.X != tt.expectedX || rover.Y != tt.expectedY || rover.Direction != tt.expectedDir {
				t.Errorf("got (%d, %d, %s), expected (%d, %d, %s)", rover.X, rover.Y, rover.Direction, tt.expectedX, tt.expectedY, tt.expectedDir)
			}
			if rover.Coordinates() != tt.line {
				t.Errorf("Coordinates(): got %q, expected %q", rover.Coordinates(), tt.line)
			}
		})
	}
}
