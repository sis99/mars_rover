package main

import (
	"fmt"
	"os"
	"strings"
)

func run(path string) (string, error) {
	plateau, roversWithActions, err := parseFile(path)
	if err != nil {
		return "", err
	}

	for i := range roversWithActions {
		for _, action := range roversWithActions[i].Actions {
			if err := roversWithActions[i].Rover.ExecuteAction(action, plateau); err != nil {
				return "", err
			}
		}
	}

	coords := make([]string, len(roversWithActions))
	for i := range roversWithActions {
		coords[i] = roversWithActions[i].Rover.Coordinates()
	}
	return strings.Join(coords, "\n"), nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: mars_rover <path>")
		os.Exit(1)
	}

	result, err := run(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(result)
}
