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

	almanach := NewAlmanach(lines, false)
	locationNoPairs := almanach.LowestLocation()
	fmt.Println("(1) lowest seeded location (no pairs):", locationNoPairs)

	almanach = NewAlmanach(lines, true)
	locationWithPairs := almanach.LowestLocation()
	fmt.Println("(2) lowest seeded location (with pairs):", locationWithPairs)
}
