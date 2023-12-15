//go:build !test

package main

import (
	"fmt"

	iox "github.com/advent-of-code-2023/internal/io"
)

func main() {
	lines, err := iox.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}

	s := NewSchematic(lines)
	sumParts := solveParts(s)
	fmt.Println("(1) sum of qualifying parts:", sumParts)
	sumGears := solveGears(s)
	fmt.Println("(2) sum of gear ratios:", sumGears)
}
