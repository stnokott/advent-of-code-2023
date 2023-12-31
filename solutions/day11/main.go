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

	sumDistances := s.SumDistances()
	fmt.Println("(1) sum of minimum distances between galaxy pairs:", sumDistances)
}
