// Package main runs the input for Day 4
package main

import (
	"strings"

	"github.com/advent-of-code-2023/internal/slicesx"
	"github.com/advent-of-code-2023/internal/stringsx"
)

// parseNumbersString converts a string of space-separated integers to a slice of integers.
//
// Example input: "41 48 83 86 17"
// Example output: [41, 48, 83, 86, 17]
func parseNumbersString(s string) []int {
	cleaned := strings.TrimSpace(strings.ReplaceAll(s, "  ", " "))
	numbers := strings.Split(cleaned, " ")
	if s == "" {
		return []int{}
	}
	result := make([]int, len(numbers))
	for i, n := range numbers {
		result[i] = stringsx.MustAtoi(n)
	}
	return result
}

// extractNumbers returns the winning numbers and the numbers we have from a card string representation.
//
// Example input: "Card 1: 41 48 | 83  9 48 53"
// Example output: ([41, 48], [83, 9, 48, 53])
func extractNumbers(s string) (wins []int, have []int) {
	raw := strings.SplitN(s, ": ", 2)[1] // remove "Card <number>: " prefix
	parts := strings.SplitN(raw, " | ", 2)
	wins, have = parseNumbersString(parts[0]), parseNumbersString(parts[1])
	return
}

// cardPoints returns the point value of a card by parsing its string representation.
func cardPoints(s string) int {
	wins, have := extractNumbers(s)
	intersect := slicesx.Intersect(wins, have)
	return cardPointAccu(len(intersect))
}

// cardPointAccu accumulates the points for a card.
//
// Specification: "The first match makes the card worth one point and each match after the first doubles the point value of that card."
func cardPointAccu(numWins int) int {
	if numWins == 0 {
		return 0
	}
	return 1 << (numWins - 1)
}

func sumCardPoints(lines []string) int {
	sum := 0
	for _, line := range lines {
		sum += cardPoints(line)
	}
	return sum
}
