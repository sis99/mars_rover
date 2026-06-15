package main

import "fmt"

type Direction string

const (
	North Direction = "N"
	East  Direction = "E"
	South Direction = "S"
	West  Direction = "W"
)

var allDirections = [4]Direction{North, East, South, West}

type Action string

const (
	Left    Action = "L"
	Right   Action = "R"
	Forward Action = "M"
)

type Plateau struct {
	MaxX, MaxY int
	MinX, MinY int
}

func NewPlateau(maxX, maxY int) (*Plateau, error) {
	return NewPlateauWithBounds(maxX, maxY, 0, 0)
}

func NewPlateauWithBounds(maxX, maxY, minX, minY int) (*Plateau, error) {
	if maxX <= minX && maxY <= minY {
		return nil, ErrInvalidPlateauSize
	}
	return &Plateau{MaxX: maxX, MaxY: maxY, MinX: minX, MinY: minY}, nil
}

func (p Plateau) Contains(x, y int) bool {
	return p.MinX <= x && x <= p.MaxX && p.MinY <= y && y <= p.MaxY
}

type Rover struct {
	X, Y      int
	Direction Direction
}

func (r Rover) Coordinates() string {
	return fmt.Sprintf("%d %d %s", r.X, r.Y, r.Direction)
}

func directionIndex(d Direction) int {
	for i, dir := range allDirections {
		if dir == d {
			return i
		}
	}
	panic("unknown direction: " + string(d))
}

func (r *Rover) TurnRight() {
	idx := directionIndex(r.Direction)
	r.Direction = allDirections[(idx+1)%len(allDirections)]
}

func (r *Rover) TurnLeft() {
	idx := directionIndex(r.Direction)
	r.Direction = allDirections[(idx-1+len(allDirections))%len(allDirections)]
}

func (r *Rover) nextCoordinates() (int, int) {
	switch r.Direction {
	case North:
		return r.X, r.Y + 1
	case East:
		return r.X + 1, r.Y
	case South:
		return r.X, r.Y - 1
	case West:
		return r.X - 1, r.Y
	}
	panic("unknown direction: " + string(r.Direction))
}

func (r *Rover) Move(plateau *Plateau) error {
	newX, newY := r.nextCoordinates()
	if !plateau.Contains(newX, newY) {
		return ErrInvalidMove
	}
	r.X = newX
	r.Y = newY
	return nil
}

func (r *Rover) ExecuteAction(action Action, plateau *Plateau) error {
	switch action {
	case Left:
		r.TurnLeft()
	case Right:
		r.TurnRight()
	case Forward:
		return r.Move(plateau)
	default:
		return ErrInvalidAction
	}
	return nil
}
