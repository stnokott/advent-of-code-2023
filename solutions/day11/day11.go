// Package main runs the input for Day 11
package main

import (
	"github.com/advent-of-code-2023/internal/mathx"
	"github.com/advent-of-code-2023/internal/slicesx"
)

const galaxy = '#'

type coord [2]int

// Sky represents multiple galaxies in a 2D grid.
type Sky struct {
	galaxies             []coord
	rowsEmpty, colsEmpty []bool
}

// NewSky creates a new sky instance from the input lines.
func NewSky(lines []string) Sky {
	rowsEmpty := slicesx.Repeat(true, len(lines))
	colsEmpty := slicesx.Repeat(true, len(lines[0]))

	coords := []coord{}
	for y, line := range lines {
		for x, c := range line {
			if c == galaxy {
				coords = append(coords, coord{x, y})
				rowsEmpty[y] = false
				colsEmpty[x] = false
			}
		}
	}

	return Sky{
		galaxies:  coords,
		rowsEmpty: rowsEmpty,
		colsEmpty: colsEmpty,
	}
}

func (s Sky) dist(a coord, b coord, emptyMultiplier int) int {
	expX := 0
	for x := mathx.MinInt(a[0], b[0]); x <= mathx.MaxInt(a[0], b[0]); x++ {
		if s.colsEmpty[x] {
			expX++
		}
	}
	expY := 0
	for y := mathx.MinInt(a[1], b[1]); y <= mathx.MaxInt(a[1], b[1]); y++ {
		if s.rowsEmpty[y] {
			expY++
		}
	}

	// no diagonal movement, so we move as if on stair steps
	dx := mathx.AbsInt(a[0]-b[0]) + expX*(emptyMultiplier-1)
	dy := mathx.AbsInt(a[1]-b[1]) + expY*(emptyMultiplier-1)
	return dx + dy
}

func (s Sky) makePairs() [][2]coord {
	pairs := [][2]coord{}
	for i, g1 := range s.galaxies {
		for _, g2 := range s.galaxies[i+1:] {
			pairs = append(pairs, [2]coord{g1, g2})
		}
	}
	return pairs
}

// SumDistances calculates the sum of minimum distances between all distinct, non-repeating galaxy pairs.
func (s Sky) SumDistances(emptyMultiplier int) int {
	pairs := s.makePairs()
	sum := 0
	for _, pair := range pairs {
		sum += s.dist(pair[0], pair[1], emptyMultiplier)
	}
	return sum
}
