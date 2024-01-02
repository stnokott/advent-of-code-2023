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

	sumPermutations := SumPermutations(lines)
	fmt.Println("(1) sum of possible permutations:", sumPermutations)
}
