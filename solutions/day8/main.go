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

	numSteps := calculateSteps(lines, false)
	fmt.Println("(1) total steps taken to end:", numSteps)
	numSteps = calculateSteps(lines, true)
	fmt.Println("(2) total steps taken to end being a ghost:", numSteps)
}
