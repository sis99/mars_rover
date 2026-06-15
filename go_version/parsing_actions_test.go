package main

import (
	"errors"
	"slices"
	"testing"
)

var invalidActionLines = []string{
	"l",
	"SSSS",
	"LMRT",
	"LM M",
}

func TestParseActions_Invalid(t *testing.T) {
	for _, line := range invalidActionLines {
		t.Run(line, func(t *testing.T) {
			_, err := parseActions(line)
			if !errors.Is(err, ErrInvalidAction) {
				t.Errorf("expected ErrInvalidAction, got %v", err)
			}
		})
	}
}

var validActionLines = []struct {
	line     string
	expected []Action
}{
	{"R", []Action{Right}},
	{"RL", []Action{Right, Left}},
	{"RRRR", []Action{Right, Right, Right, Right}},
}

func TestParseActions_Valid(t *testing.T) {
	for _, tt := range validActionLines {
		t.Run(tt.line, func(t *testing.T) {
			got, err := parseActions(tt.line)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !slices.Equal(got, tt.expected) {
				t.Errorf("got %v, expected %v", got, tt.expected)
			}
		})
	}
}
