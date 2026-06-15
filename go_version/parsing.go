package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	// matches two non-negative integers separated by a space e.g. "5 5"
	plateauRE = regexp.MustCompile(`^\d+ \d+$`)
	// matches two integers and a cardinal direction separated by spaces e.g. "1 2 N"
	roverRE = regexp.MustCompile(`^-?\d+ -?\d+ [NESW]$`)
)

func parsePlateau(line string) (*Plateau, error) {
	if line == "" {
		return nil, ErrEmptyFile
	}
	if !plateauRE.MatchString(line) {
		return nil, ErrInvalidPlateauSize
	}
	xString, yString, _ := strings.Cut(line, " ") // regex guarantees valid non-negative integers
	x, _ := strconv.Atoi(xString)
	y, _ := strconv.Atoi(yString)
	return NewPlateau(x, y)
}

func parseRover(line string) (Rover, error) {
	if !roverRE.MatchString(line) {
		return Rover{}, ErrInvalidRoverParams
	}
	xString, rest, _ := strings.Cut(line, " ") // regex guarantees valid integers
	yString, direction, _ := strings.Cut(rest, " ")
	x, _ := strconv.Atoi(xString)
	y, _ := strconv.Atoi(yString)
	return Rover{X: x, Y: y, Direction: Direction(direction)}, nil
}

func parseActions(line string) ([]Action, error) {
	actions := make([]Action, 0, len(line))
	for _, ch := range line {
		a := Action(string(ch))
		switch a {
		case Left, Right, Forward:
			actions = append(actions, a)
		default:
			return nil, fmt.Errorf("%w: %c", ErrInvalidAction, ch)
		}
	}
	return actions, nil
}

func extractRover(plateau *Plateau, line string) (Rover, error) {
	rover, err := parseRover(line)
	if err != nil {
		return Rover{}, err
	}
	if !plateau.Contains(rover.X, rover.Y) {
		return Rover{}, ErrInvalidRoverParams
	}
	return rover, nil
}

type RoverWithActions struct {
	Rover   Rover
	Actions []Action
}

func parseStream(r io.Reader) (*Plateau, []RoverWithActions, error) {
	scanner := bufio.NewScanner(r)
	var lines []string
	for scanner.Scan() {
		if line := strings.TrimSpace(scanner.Text()); line != "" {
			lines = append(lines, line)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	if len(lines) == 0 {
		return nil, nil, ErrEmptyFile
	}

	plateau, err := parsePlateau(lines[0])
	if err != nil {
		return nil, nil, err
	}

	var result []RoverWithActions
	for i := 1; i < len(lines); i += 2 {
		rover, err := extractRover(plateau, lines[i])
		if err != nil {
			return nil, nil, err
		}

		if i+1 >= len(lines) {
			return nil, nil, MissingActionsError{Rover: rover.Coordinates()}
		}

		actions, err := parseActions(lines[i+1])
		if err != nil {
			return nil, nil, err
		}

		result = append(result, RoverWithActions{Rover: rover, Actions: actions})
	}

	return plateau, result, nil
}

func parseFile(path string) (*Plateau, []RoverWithActions, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer func() { _ = f.Close() }()
	return parseStream(f)
}
