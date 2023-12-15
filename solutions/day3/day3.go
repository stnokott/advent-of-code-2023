// Package main runs the input for Day 3
package main

import (
	"slices"

	stringsx "github.com/advent-of-code-2023/internal/strings"
)

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isDot(c byte) bool {
	return c == '.'
}

// Element is one horizontal sequence of symbols of the same type.
// This could be "123", but also "$%&"
type Element struct {
	s        string
	x0       int
	isNumber bool
}

func elementsInLine(line string) []*Element {
	els := []*Element{}
	// not using for..range because we want to modify x inside the loop
	for x := 0; x < len(line); x++ {
		if !isDot(line[x]) {
			// are we parsing a number (or a symbol)?
			isNumber := isDigit(line[x])
			var s string
			// iterate line from x until any of the following occurs:
			// - end of line reached
			// - different character type (digit/symbol) encountered
			// - dot encountered
			for i := x; i < len(line) && isDigit(line[i]) == isNumber && !isDot(line[i]); i++ {
				// add each encountered character to element string
				s += string(line[i])
			}
			el := &Element{
				s:        s,
				x0:       x,
				isNumber: isNumber,
			}
			els = append(els, el)
			x += len(el.s) - 1
		}
	}
	return els
}

// Schematic contains the input data as a 2D slice of Elements.
// The elements are in the same row as they appear in the input string.
// The elements in each row appear in the same order as in the input string.
type Schematic [][]*Element

// NewSchematic creates a new Schematic from the provided input string
func NewSchematic(lines []string) Schematic {
	rows := [][]*Element{}
	for _, line := range lines {
		row := elementsInLine(line)
		rows = append(rows, row)
	}
	return rows
}

// elementAt returns the element at the specified input string coordinates or nil, if not found.
func (s Schematic) elementAt(x, y int) *Element {
	if y < 0 || y >= len(s) {
		return nil
	}
	for _, el := range s[y] {
		if x >= el.x0 && x < el.x0+len(el.s) {
			return el
		}
	}
	return nil
}

// adjacencies defines dx,dy for 8-adjacency
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

func (s Schematic) getAdjacentElements(el *Element, y int) []*Element {
	x := el.x0
	elements := []*Element{}
	// iterate characters of element string
	for dx := range el.s {
		// iterate adjacencies
		for _, adj := range adjacencies {
			charX, charY := x+dx+adj[0], y+adj[1]
			elAdj := s.elementAt(charX, charY)
			// only add if valid element found that is not the source element
			// and not already part of the result list
			if elAdj != nil &&
				elAdj != el &&
				!slices.Contains(elements, elAdj) {
				elements = append(elements, elAdj)
			}
		}
	}
	return elements
}

func (s Schematic) sumLine(y int) int {
	sum := 0
	for _, el := range s[y] {
		if !el.isNumber {
			continue
		}
		adjacentElements := s.getAdjacentElements(el, y)
		if len(adjacentElements) > 0 {
			sum += stringsx.MustAtoi(el.s)
		}
	}
	return sum
}

func solve(lines []string) int {
	s := NewSchematic(lines)
	sum := 0
	for y := range s {
		sum += s.sumLine(y)
	}
	return sum
}
