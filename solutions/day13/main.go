//go:build !test

package main

import (
	"fmt"

	"github.com/advent-of-code-2023/internal/iox"
	"github.com/advent-of-code-2023/internal/slicesx"
)

func main() {
	lines, err := iox.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}

	notes := slicesx.Split(lines, "")

	sum := 0
	for _, note := range notes {
		left, top := FindMirrors(note)
		sum += left + top*100
	}
	fmt.Println("(1) sum of notes:", sum)
}
