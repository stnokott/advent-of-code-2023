// Package main runs the input for Day 10
package main

import "math"

const (
	ground = '.'
	start  = 'S'
)

type coord [2]int

func (c coord) add(d coord) coord {
	return coord{c[0] + d[0], c[1] + d[1]}
}

var pipes = map[rune][2]coord{
	'|': {{0, -1}, {0, 1}},
	'-': {{-1, 0}, {1, 0}},
	'L': {{0, -1}, {1, 0}},
	'J': {{0, -1}, {-1, 0}},
	'7': {{-1, 0}, {0, 1}},
	'F': {{1, 0}, {0, 1}},
}

type sketch struct {
	g     [][]rune
	start coord
}

func newSketch(lines []string) sketch {
	g := make([][]rune, len(lines))
	var startCoord coord
	for y, line := range lines {
		row := make([]rune, len(line))
		for x, c := range line {
			row[x] = c
			if c == start {
				startCoord = coord{x, y}
			}
		}
		g[y] = row
	}
	return sketch{
		g:     g,
		start: startCoord,
	}
}

const outOfBounds rune = -1

func (s sketch) at(c coord) rune {
	if c[1] >= 0 && c[0] >= 0 && c[1] < len(s.g) && c[0] < len(s.g[c[1]]) {
		return s.g[c[1]][c[0]]
	}
	return outOfBounds
}

func (s sketch) move(from coord, via coord) (to coord) {
	delta := pipes[s.at(via)]

	to = via.add(delta[0])
	// if delta[0] yields the same coordinates as the ones we're coming from,
	// we need to use the other delta instead.
	if to == from {
		to = via.add(delta[1])
	}
	return
}

var directions = []coord{
	{-1, -1},
	{0, -1},
	{1, -1},
	{1, 0},
	{1, 1},
	{0, 1},
	{-1, 1},
	{-1, 0},
}

// firstPipe finds the first pipe connected to the start pipe
func (s sketch) firstPipe() coord {
	for _, dir := range directions {
		pipeCoord := s.start.add(dir)
		if pipe := s.at(pipeCoord); pipe != outOfBounds {
			pipeDirs := pipes[pipe]
			for _, pipeDir := range pipeDirs {
				if pipeDir[0]*-1 == dir[0] && pipeDir[1]*-1 == dir[1] {
					return pipeCoord
				}
			}
		}
	}
	panic("no pipe connected to the start pipe")
}

func (s sketch) loopLength() int {
	current := s.start
	next := s.firstPipe()

	length := 1
	for ; next != s.start; length++ {
		current, next = next, s.move(current, next)
	}
	return length
}

func (s sketch) MaxDistance() int {
	return int(math.Ceil(float64(s.loopLength()) / 2))
}
