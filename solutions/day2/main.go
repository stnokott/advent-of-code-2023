//go:build !test

package main

import (
	"fmt"

	"github.com/advent-of-code-2023/internal/io"
)

func main() {
	scenario := Set{R: 12, G: 13, B: 14}

	lines, err := io.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}
	sum := solveDay1(scenario, lines...)
	fmt.Println("(1) sum of possible game IDs:", sum)
	power := solveDay2(lines)
	fmt.Println("(2) sum of power of minimum sets for all games:", power)
}
