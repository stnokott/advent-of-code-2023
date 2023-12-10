//go:build !test

package main

import (
	"fmt"

	"github.com/advent-of-code-2023/internal/io"
)

func main() {
	lines, err := io.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}

	s := NewSchematic(lines)
	solution := s.solve()
	fmt.Println(solution)
}
