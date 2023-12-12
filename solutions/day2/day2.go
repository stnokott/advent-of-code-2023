// Package main runs the input for Day 2
package main

import (
	"fmt"
	"strings"

	"github.com/advent-of-code-2023/internal/str"
)

// Set contains a configuration for one set
type Set struct {
	R, G, B int
}

// HasSubset returns true if d is a complete subset of this set.
func (s Set) HasSubset(s2 Set) bool {
	return s.R >= s2.R && s.G >= s2.G && s.B >= s2.B
}

// Power it the number of each color multiplied together
func (s Set) Power() int {
	return s.R * s.G * s.B
}

// Game contains its ID and the list of sets in this game.
// Use NewGame() to construct an instance from a string.
type Game struct {
	ID   int
	Sets []Set
}

// HasSubset returns true if the passed set is a subset of all sets in this game.
func (g Game) HasSubset(set Set) bool {
	for _, s := range g.Sets {
		if !set.HasSubset(s) {
			return false
		}
	}
	return true
}

// MinPossibleSet returns the smallest set (by cube count) for which this game would be possible.
func (g Game) MinPossibleSet() Set {
	s := Set{} // init with 0,0,0
	for _, gameSet := range g.Sets {
		s.R = max(s.R, gameSet.R)
		s.G = max(s.G, gameSet.G)
		s.B = max(s.B, gameSet.B)
	}
	return s
}

// NewGame creates a new Game instance from an input string.
// Example input: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
func NewGame(s string) Game {
	colonIndex := strings.IndexRune(s, ':')
	// get id (string between first ' ' and ':')
	id := str.MustAtoi(s[strings.IndexRune(s, ' ')+1 : colonIndex])
	// get sets string (everything after colon+space)
	setsString := s[colonIndex+2:]
	sets := parseSets(setsString)
	return Game{
		ID:   id,
		Sets: sets,
	}
}

// parseSets reads a string of sets and returns the list of parsed set objects
// Example input: "3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
func parseSets(setsString string) []Set {
	colorsStrings := strings.Split(setsString, "; ")
	// colors[i] = "3 blue"
	configs := make([]Set, len(colorsStrings))
	for i, colorsString := range colorsStrings {
		configs[i] = parseColors(colorsString)
	}
	return configs
}

// parseColors parses a string for one set, containing 1-3 colors into a Set object.
// Example input: "3 blue, 4 red"
func parseColors(colorsString string) Set {
	colorNumbers := strings.Split(colorsString, ", ")
	// colorNumbers[i] = "3 blue"
	cfg := Set{}
	for _, colorNumber := range colorNumbers {
		n, color := parseColor(colorNumber)
		switch color {
		case Red:
			cfg.R = n
		case Green:
			cfg.G = n
		case Blue:
			cfg.B = n
		default:
			panic(fmt.Sprintf("unknown color type %T returned for '%s'", color, colorNumber))
		}
	}
	return cfg
}

// Color has constants for R, G & B
type Color int

const (
	Red   Color = iota // Red is the color red
	Green              // Green is the color green
	Blue               // Blue is the color blue
)

// SetColorNumberFrom sets the attribute of this set matching the provided string.
// Example input: "3 blue"
func parseColor(s string) (n int, c Color) {
	var err error
	defer func() {
		if err != nil {
			panic(fmt.Sprintf("error getting number from color string '%s': %v", s, err))
		}
	}()
	parts := strings.Split(s, " ")
	numberStr, color := parts[0], parts[1]
	n = str.MustAtoi(numberStr)
	switch color {
	case "red":
		c = Red
	case "green":
		c = Green
	case "blue":
		c = Blue
	default:
		err = fmt.Errorf("invalid color '%s'", color)
	}
	return
}

func solveDay1(scenario Set, games ...string) int {
	x := 0
	for _, gameStr := range games {
		game := NewGame(gameStr)
		if game.HasSubset(scenario) {
			x += game.ID
		}
	}
	return x
}

func solveDay2(games []string) int {
	x := 0
	for _, gameStr := range games {
		game := NewGame(gameStr)
		minSet := game.MinPossibleSet()
		x += minSet.Power()
	}
	return x
}
