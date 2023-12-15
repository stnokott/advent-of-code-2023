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
	solution := solveParts(s)
	fmt.Println(solution)
}
