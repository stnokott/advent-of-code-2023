// Package main runs the input for Day 9
package main

import (
	"slices"
	"strings"

	"github.com/advent-of-code-2023/internal/slicesx"
	"github.com/advent-of-code-2023/internal/strconvx"
)

func diffs(vals []int) []int {
	diffs := make([]int, len(vals)-1)

	for i := 0; i < len(vals)-1; i++ {
		diffs[i] = vals[i+1] - vals[i]
	}
	return diffs
}

func isZero(x int) bool {
	return x == 0
}

func next(vals []int) int {
	diffs := diffs(vals)

	if slicesx.All(diffs, isZero) {
		// return last value + 0
		return vals[len(vals)-1]
	}
	// return last value + recursive next
	return vals[len(vals)-1] + next(diffs)
}

func parse(line string) []int {
	matches := strings.Split(line, " ")
	ints := make([]int, len(matches))
	for i, x := range matches {
		ints[i] = strconvx.MustAtoi(x)
	}
	return ints
}

// SumPredictions returns the sum of each prediction for each line.
func SumPredictions(lines []string, reverse bool) int {
	sum := 0
	for _, line := range lines {
		nums := parse(line)
		if reverse {
			slices.Reverse(nums)
		}
		sum += next(nums)
	}
	return sum
}
