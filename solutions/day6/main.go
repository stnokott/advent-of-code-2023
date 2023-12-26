//go:build !test

package main

import (
	"fmt"

	"github.com/advent-of-code-2023/internal/iox"
)

func main() {
	lines, err := iox.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}

	races := parseRaces(lines)
	result := raceMultiplications(races)
	fmt.Println("(1) multiplications of ways the records can be beat:", result)
	mergedRace := parseRace(lines)
	result = mergedRace.numSolutions()
	fmt.Println("(2) multiplications of ways the record can be beat:", result)
}
