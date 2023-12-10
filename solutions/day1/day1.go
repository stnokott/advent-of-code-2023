// Package main runs the input for Day 1
package main

import (
	"strconv"
	"strings"
)

var replacements = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

// numbersInString returns all *overlapping* occurences of valid numbers in s.
func numbersInString(s string) []string {
	numbers := make([]string, 0)
	// iterate string
	for c := range s {
		// iterate all possible number strings
		for k, v := range replacements {
			if strings.HasPrefix(s[c:], k) {
				// number word found
				numbers = append(numbers, s[c:c+len(k)])
			} else if s[c] == v[0] {
				// number character found
				numbers = append(numbers, string(s[c]))
			}
		}
	}
	return numbers
}

func parseNumber(s string) int {
	// replace words with numbers
	for k, v := range replacements {
		if s == k {
			s = v
			break
		}
	}
	// convert to integer
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func extractNumber(s string) int {
	numbers := numbersInString(s)
	a := parseNumber(numbers[0])
	b := parseNumber(numbers[len(numbers)-1])
	return a*10 + b
}

func solve(calibrations []string) int {
	sum := 0
	for _, s := range calibrations {
		sum += extractNumber(s)
	}
	return sum
}
