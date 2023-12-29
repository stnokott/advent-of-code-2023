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

	g := NewGame(lines, false)
	winnings := g.totalWinnings()
	fmt.Println("(1) total winnings:", winnings)
	g = NewGame(lines, true)
	winnings = g.totalWinnings()
	fmt.Println("(2) total winnings with joker wildcards:", winnings)
}
