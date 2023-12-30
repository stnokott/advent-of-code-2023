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

	s := newSketch(lines)

	maxDistance := s.MaxDistance()
	fmt.Println("(1) max loop distance from starting point:", maxDistance)
	countInner := s.NumInnerCells()
	fmt.Println("(2) number of inner cells:", countInner)
}
