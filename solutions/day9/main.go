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

	sum := SumPredictions(lines, false)
	fmt.Println("(1) sum of predictions:", sum)
	sum = SumPredictions(lines, true)
	fmt.Println("(1) sum of predictions reversed:", sum)
}
