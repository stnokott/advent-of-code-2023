// Package main runs the input for Day 10
package main

import (
	"math"
)

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

type grid[T any] [][]T

func (g grid[T]) inBounds(c coord) bool {
	return c[1] >= 0 && c[0] >= 0 && c[1] < len(g) && c[0] < len(g[c[1]])
}

func (g grid[T]) at(c coord) T {
	return g[c[1]][c[0]]
}

type sketch struct {
	g     grid[rune]
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

func (s sketch) move(from coord, via coord) (to coord) {
	delta := pipes[s.g.at(via)]

	to = via.add(delta[0])
	// if delta[0] yields the same coordinates as the ones we're coming from,
	// we need to use the other delta instead.
	if to == from {
		to = via.add(delta[1])
	}
	return
}

// counter-clockwise for shoelace formula
var directions = []coord{
	{-1, 0},
	{-1, 1},
	{0, 1},
	{1, 1},
	{1, 0},
	{1, -1},
	{0, -1},
	{-1, -1},
}

// firstPipe finds the first pipe connected to the start pipe
func (s sketch) firstPipe() coord {
	for _, dir := range directions {
		pipeCoord := s.start.add(dir)
		if !s.g.inBounds(pipeCoord) {
			continue
		}
		pipe := s.g.at(pipeCoord)
		pipeDirs := pipes[pipe]
		for _, pipeDir := range pipeDirs {
			if pipeDir[0]*-1 == dir[0] && pipeDir[1]*-1 == dir[1] {
				return pipeCoord
			}
		}
	}
	panic("no pipe connected to the start pipe")
}

func (s sketch) loop() []coord {
	current := s.start
	next := s.firstPipe()

	coords := []coord{current}

	for next != s.start {
		coords = append(coords, next)
		current, next = next, s.move(current, next)
	}
	return coords
}

func (s sketch) MaxDistance() int {
	loop := s.loop()
	return int(math.Ceil(float64(len(loop)) / 2))
}

func (s sketch) NumInnerCells() int {
	loop := s.loop()

	var sum int
	for i := range loop {
		sum += loop[i][0] * loop[(i+1)%len(loop)][1]
		sum -= loop[(i+1)%len(loop)][0] * loop[i][1]
	}

	shoelaceSurface := math.Abs(float64(sum))/2 - float64(len(loop))/2 + 1
	return int(math.Ceil(shoelaceSurface))
}
