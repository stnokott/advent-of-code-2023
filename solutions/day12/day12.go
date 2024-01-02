// Package main runs the input for Day 12
package main

import (
	"strings"

	"github.com/advent-of-code-2023/internal/strconvx"
)

// valid checks if s fulfills the provided target.
//
// The string may only consist of '.' and '#' characters.
//
//	.###.##..#  3,2,1  -> true
//	.##..##..#  3,2,1  -> false
//
// Note: it it possible to implement this function by creating a regular expression from target.
// Although such a solution might be easier to read, it will most certainly be slower, so I went with this approach instead.
func valid(s string, target []int) bool {
	parts := strings.FieldsFunc(s, func(r rune) bool {
		return r == '.'
	})
	if len(parts) != len(target) {
		return false
	}
	for i, t := range target {
		if t != len(parts[i]) {
			return false
		}
	}
	return true
}

func permutations(s string, i int, target []int) int {
	// move forward until at unknown position
	for i < len(s) && s[i] != '?' {
		i++
	}
	if i == len(s) {
		// end reached, no '?' characters left
		valid := valid(s, target)
		if valid {
			return 1
		}
		return 0
	}
	// end not reached yet, so we are at a '?' character and must split up here

	p0 := permutations(s[:i]+"#"+s[i+1:], i+1, target)
	p1 := permutations(s[:i]+"."+s[i+1:], i+1, target)
	return p0 + p1
}

func parseLine(s string) (string, []int) {
	parts := strings.Split(s, " ")
	targetParts := strings.Split(parts[1], ",")
	target := make([]int, len(targetParts))
	for i, tp := range targetParts {
		target[i] = strconvx.MustAtoi(tp)
	}
	return parts[0], target
}

// SumPermutations returns the sum of all possible permutations for each line which fulfill the
// requirements of #-sequences.
func SumPermutations(lines []string) int {
	sum := 0
	for _, line := range lines {
		s, target := parseLine(line)
		sum += permutations(s, 0, target)
	}
	return sum
}
