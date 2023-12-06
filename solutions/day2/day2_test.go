// Package main runs the input for Day 2
package main

import (
	"reflect"
	"testing"
)

func TestNewGame(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want Game
	}{
		{"short", "Game 1: 2 blue", Game{ID: 1, Sets: []Set{{B: 2}}}},
		{"one full set", "Game 3: 1 blue, 2 red, 3 green", Game{ID: 3, Sets: []Set{{R: 2, G: 3, B: 1}}}},
		{"multiple sets", "Game 6: 1 blue; 2 red, 3 green; 1 red, 2 blue, 3 green", Game{ID: 6, Sets: []Set{{B: 1}, {R: 2, G: 3}, {R: 1, G: 3, B: 2}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGame(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGame() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseSets(t *testing.T) {
	tests := []struct {
		name       string
		setsString string
		want       []Set
	}{
		{"one set, one color", "3 red", []Set{{R: 3}}},
		{"one set, all colors", "6 red, 8 green, 8 blue", []Set{{R: 6, G: 8, B: 8}}},
		{"multiple sets", "5 green; 8 green, 2 red; 9 green, 5 blue, 4 red", []Set{{G: 5}, {R: 2, G: 8}, {R: 4, G: 9, B: 5}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseSets(tt.setsString); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseSets() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseColors(t *testing.T) {
	tests := []struct {
		name         string
		colorsString string
		want         Set
	}{
		{"one color", "7 blue", Set{B: 7}},
		{"two colors", "9 red, 2 green", Set{R: 9, G: 2}},
		{"three colors", "5 green, 2 blue, 9 red", Set{R: 9, G: 5, B: 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseColors(tt.colorsString); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseColors() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseColor(t *testing.T) {
	tests := []struct {
		name  string
		str   string
		wantN int
		wantC Color
	}{
		{"red", "2 red", 2, Red},
		{"blue", "6 blue", 6, Blue},
		{"green", "1 green", 1, Green},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotN, gotC := parseColor(tt.str)
			if gotN != tt.wantN {
				t.Errorf("parseColor() gotN = %v, want %v", gotN, tt.wantN)
			}
			if !reflect.DeepEqual(gotC, tt.wantC) {
				t.Errorf("parseColor() gotC = %v, want %v", gotC, tt.wantC)
			}
		})
	}
}

func TestSetHasSubset(t *testing.T) {
	tests := []struct {
		name string
		s    *Set
		s2   Set
		want bool
	}{
		{"equal", &Set{R: 1, G: 2, B: 3}, Set{R: 1, G: 2, B: 3}, true},
		{"one too low", &Set{R: 1, G: 2, B: 3}, Set{R: 2, G: 2, B: 3}, false},
		{"all higher", &Set{R: 11, G: 22, B: 33}, Set{R: 1, G: 2, B: 3}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.HasSubset(tt.s2); got != tt.want {
				t.Errorf("Set.HasSubset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGameHasSubset(t *testing.T) {
	tests := []struct {
		name string
		g    Game
		set  Set
		want bool
	}{
		{"equal", Game{Sets: []Set{{R: 1, G: 2, B: 3}}}, Set{R: 1, G: 2, B: 3}, true},
		{"is subset", Game{Sets: []Set{{R: 2, G: 4, B: 6}}}, Set{R: 3, G: 10, B: 11}, true},
		{"is not subset", Game{Sets: []Set{{R: 2, G: 4, B: 6}}}, Set{R: 3, G: 10, B: 5}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.HasSubset(tt.set); got != tt.want {
				t.Errorf("Game.HasSubset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolve(t *testing.T) {
	type args struct {
		scenario Set
		games    []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"one possible game",
			args{
				scenario: Set{R: 10, G: 11, B: 12},
				games: []string{
					"Game 1: 9 red, 6 green, 5 blue",
				},
			},
			1,
		},
		{
			"one impossible game",
			args{
				scenario: Set{R: 10, G: 11, B: 12},
				games: []string{
					"Game 1: 12 green, 6 red, 5 blue",
				},
			},
			0,
		},
		{
			"multiple games",
			args{
				scenario: Set{R: 10, G: 11, B: 12},
				games: []string{
					"Game 1: 10 green, 6 red, 5 blue",
					"Game 2: 20 blue, 2 red, 8 green",
					"Game 3: 9 green, 2 blue, 9 red; 10 blue, 11 green, 1 red",
				},
			},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.scenario, tt.args.games...); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
