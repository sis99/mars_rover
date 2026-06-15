package main

import "testing"

func checkPlateau(p *Plateau, err error) *Plateau {
	if err != nil {
		panic(err)
	}
	return p
}

func TestPlateau_Contains_Invalid(t *testing.T) {
	tests := []struct {
		name    string
		x, y    int
		plateau *Plateau
	}{
		{"negative coordinate", -1, 0, checkPlateau(NewPlateau(1, 1))},
		{"positive but out of bounds", 2, 0, checkPlateau(NewPlateau(1, 1))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.plateau.Contains(tt.x, tt.y) {
				t.Errorf("expected Contains(%d, %d) to be false", tt.x, tt.y)
			}
		})
	}
}

func TestPlateau_Contains_Valid(t *testing.T) {
	tests := []struct {
		name    string
		x, y    int
		plateau *Plateau
	}{
		{"negative if plateau allows", -1, -1, checkPlateau(NewPlateauWithBounds(2, 2, -1, -1))},
		{"same as plateau boundary", 2, 2, checkPlateau(NewPlateau(2, 2))},
		{"inside plateau", 1, 1, checkPlateau(NewPlateau(2, 2))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.plateau.Contains(tt.x, tt.y) {
				t.Errorf("expected Contains(%d, %d) to be true", tt.x, tt.y)
			}
		})
	}
}
