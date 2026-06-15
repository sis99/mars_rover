package main

import "errors"

var (
	ErrInvalidPlateauSize = errors.New("invalid plateau size")
	ErrInvalidMove        = errors.New("invalid move")
	ErrInvalidAction      = errors.New("invalid action")
	ErrInvalidRoverParams = errors.New("invalid rover params")
	ErrEmptyFile          = errors.New("empty file")
)

type MissingActionsError struct {
	Rover string
}

func (e MissingActionsError) Error() string {
	return "missing actions for rover: " + e.Rover
}
