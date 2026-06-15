package main

import (
	"errors"
	"testing"
)

func TestParseFile_EmptyFile(t *testing.T) {
	_, _, err := parseFile("test_files/empty_file")
	if !errors.Is(err, ErrEmptyFile) {
		t.Errorf("expected ErrEmptyFile, got %v", err)
	}
}

func TestParseFile_ExtraSpaces(t *testing.T) {
	plateau, roversWithActions, err := parseFile("test_files/extra_spaces")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if plateau.MaxX != 5 || plateau.MaxY != 5 {
		t.Errorf("plateau: got (%d,%d), expected (5,5)", plateau.MaxX, plateau.MaxY)
	}
	if len(roversWithActions) != 1 {
		t.Fatalf("expected 1 rover, got %d", len(roversWithActions))
	}
	if got := roversWithActions[0].Rover.Coordinates(); got != "1 2 N" {
		t.Errorf("rover coords: got %q, expected %q", got, "1 2 N")
	}
	expected := []Action{Left, Forward}
	for i := range expected {
		if roversWithActions[0].Actions[i] != expected[i] {
			t.Errorf("action[%d]: got %s, expected %s", i, roversWithActions[0].Actions[i], expected[i])
		}
	}
}

func TestParseFile_MissingPlateau(t *testing.T) {
	_, _, err := parseFile("test_files/missing_plateau")
	if !errors.Is(err, ErrInvalidPlateauSize) {
		t.Errorf("expected ErrInvalidPlateauSize, got %v", err)
	}
}

func TestParseFile_MissingRover(t *testing.T) {
	_, _, err := parseFile("test_files/missing_rover")
	if !errors.Is(err, ErrInvalidRoverParams) {
		t.Errorf("expected ErrInvalidRoverParams, got %v", err)
	}
}

func TestParseFile_MissingActions(t *testing.T) {
	_, _, err := parseFile("test_files/missing_actions")
	var missingErr MissingActionsError
	if !errors.As(err, &missingErr) {
		t.Errorf("expected MissingActionsError, got %v", err)
	}
}
