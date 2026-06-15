package main

import (
	"errors"
	"testing"
)

var invalidPlateauLines = []string{
	"0 -1",
	"0 0",
	"0, 1",
	"ekqnfek",
	" 1   1",
}

func TestParsePlateau_Invalid(t *testing.T) {
	for _, line := range invalidPlateauLines {
		t.Run(line, func(t *testing.T) {
			_, err := parsePlateau(line)
			if err == nil {
				t.Errorf("expected error for %q", line)
			}
		})
	}
}

var validPlateauLines = []string{
	"1 1",
	"100 200",
	"0 1",
	"1 0",
}

func TestParsePlateau_Valid(t *testing.T) {
	for _, line := range validPlateauLines {
		t.Run(line, func(t *testing.T) {
			if _, err := parsePlateau(line); err != nil {
				t.Errorf("unexpected error for %q: %v", line, err)
			}
		})
	}
}

func TestPlateau_InvalidSize(t *testing.T) {
	tests := []struct {
		name                   string
		maxX, maxY, minX, minY int
	}{
		{"min same as max", 1, 1, 1, 1},
		{"max smaller than min", -1, 0, 1, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewPlateauWithBounds(tt.maxX, tt.maxY, tt.minX, tt.minY)
			if !errors.Is(err, ErrInvalidPlateauSize) {
				t.Errorf("expected ErrInvalidPlateauSize, got %v", err)
			}
		})
	}
}
