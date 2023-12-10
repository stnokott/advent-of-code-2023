// Package main runs the input for Day 3
package main

import (
	"strconv"
)

func isDigit(c rune) bool {
	return c >= '0' && c <= '9'
}

func isSymbol(c rune) bool {
	return c != '.' && !isDigit(c)
}

type Schematic struct {
	numRows, numCols int
	lines            []string
}

func NewSchematic(lines []string) *Schematic {
	return &Schematic{
		numRows: len(lines),
		numCols: len(lines[0]),
		lines:   lines,
	}
}

var adjacencies = [][]int{
	{-1, 0},
	{-1, -1},
	{0, -1},
	{1, -1},
	{1, 0},
	{1, 1},
	{0, 1},
	{-1, 1},
}

func (s *Schematic) hasAdjacentSymbol(x, y int) bool {
	for _, adj := range adjacencies {
		dx, dy := x+adj[0], y+adj[1]
		if dx >= 0 && dx < s.numCols &&
			dy >= 0 && dy < s.numRows &&
			isSymbol(rune(s.lines[dy][dx])) {
			return true
		}
	}
	return false
}

func (s *Schematic) solve() int {
	sum := 0
	for y, line := range s.lines {
		currentNumber := ""
		currentNumberValid := false
		for x, c := range line {
			digit := isDigit(c)
			if digit {
				if !currentNumberValid && s.hasAdjacentSymbol(x, y) {
					currentNumberValid = true
				}
				currentNumber += string(c)
			}
			if currentNumber != "" && (!digit || x == s.numCols-1) {
				if currentNumberValid {
					num, err := strconv.Atoi(currentNumber)
					if err != nil {
						panic(err)
					}
					sum += num
				}
				currentNumber = ""
				currentNumberValid = false
			}
		}
	}
	return sum
}
