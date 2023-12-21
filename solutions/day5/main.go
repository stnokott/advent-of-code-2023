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
	almanach := NewAlmanach(lines)

	lowestLocation := almanach.LowestLocation()
	fmt.Println("(1) lowest seeded location:", lowestLocation)
}
