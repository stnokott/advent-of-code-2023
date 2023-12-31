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

	s := NewSky(lines)

	sumDistances := s.SumDistances(2)
	fmt.Println("(1) sum of minimum distances between galaxy pairs (mult=2):", sumDistances)
	sumDistances = s.SumDistances(1000000)
	fmt.Println("(1) sum of minimum distances between galaxy pairs (mult=1000000):", sumDistances)
}
