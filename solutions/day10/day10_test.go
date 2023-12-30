// Package main runs the input for Day 10
package main

import (
	"reflect"
	"testing"
)

func TestNewSketch(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		args args
		want sketch
	}{
		{
			"simple",
			args{[]string{
				".....",
				".S-7.",
				".|.|.",
				".L-J.",
				".....",
			}},
			sketch{
				g: [][]rune{
					{'.', '.', '.', '.', '.'},
					{'.', 'S', '-', '7', '.'},
					{'.', '|', '.', '|', '.'},
					{'.', 'L', '-', 'J', '.'},
					{'.', '.', '.', '.', '.'},
				},
				start: coord{1, 1},
			},
		},
		{
			"distractions",
			args{[]string{
				"-L|F7",
				"7S-7|",
				"L|7||",
				"-L-J|",
				"L|-JF",
			}},
			sketch{
				g: [][]rune{
					{'-', 'L', '|', 'F', '7'},
					{'7', 'S', '-', '7', '|'},
					{'L', '|', '7', '|', '|'},
					{'-', 'L', '-', 'J', '|'},
					{'L', '|', '-', 'J', 'F'},
				},
				start: coord{1, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newSketch(tt.args.lines); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newSketch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSketchAt(t *testing.T) {
	s := newSketch([]string{
		".F-7.",
		".|.|.",
		".L-J.",
	})

	type args struct {
		c coord
	}
	tests := []struct {
		name string
		args args
		want rune
	}{
		{"simple", args{coord{1, 2}}, 'L'},
		{"out of bounds top left", args{coord{0, -1}}, outOfBounds},
		{"out of bounds bottom right", args{coord{5, 3}}, outOfBounds},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := s.at(tt.args.c); got != tt.want {
				t.Errorf("sketch.at() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSketchMove(t *testing.T) {
	s := newSketch([]string{
		".F-7.",
		".|.|.",
		".L-J.",
	})

	type args struct {
		from coord
		via  coord
	}
	tests := []struct {
		name   string
		args   args
		wantTo coord
	}{
		{"|", args{from: coord{3, 0}, via: coord{3, 1}}, coord{3, 2}},
		{"-", args{from: coord{3, 0}, via: coord{2, 0}}, coord{1, 0}},
		{"L", args{from: coord{1, 1}, via: coord{1, 2}}, coord{2, 2}},
		{"J", args{from: coord{2, 2}, via: coord{3, 2}}, coord{3, 1}},
		{"7", args{from: coord{2, 0}, via: coord{3, 0}}, coord{3, 1}},
		{"F", args{from: coord{2, 0}, via: coord{1, 0}}, coord{1, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTo := s.move(tt.args.from, tt.args.via); !reflect.DeepEqual(gotTo, tt.wantTo) {
				t.Errorf("sketch.move() = %v, want %v", gotTo, tt.wantTo)
			}
		})
	}
}

func TestSketchFirstPipe(t *testing.T) {
	tests := []struct {
		name      string
		s         sketch
		want      coord
		wantPanic bool
	}{
		{
			"default",
			newSketch([]string{
				".....",
				".S-7.",
				".|.|.",
				".L-J.",
				".....",
			}),
			coord{2, 1},
			false,
		},
		{
			"no connection",
			newSketch([]string{
				".....",
				".S|7.",
				".7.|.",
				".L-J.",
				".....",
			}),
			coord{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if err := recover(); (err != nil) != tt.wantPanic {
					t.Errorf("sketch.firstPipe() panic = %t, wantPanic = %t", err != nil, tt.wantPanic)
				}
			}()
			if got := tt.s.firstPipe(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sketch.firstPipe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSketchLoopLength(t *testing.T) {
	tests := []struct {
		name string
		s    sketch
		want int
	}{
		{
			"circular",
			newSketch([]string{
				"-L|F7",
				"7S-7|",
				"L|7||",
				"-L-J|",
				"L|-JF",
			}),
			8,
		},
		{
			"complex",
			newSketch([]string{
				"7-F7-",
				".FJ|7",
				"SJLL7",
				"|F--J",
				"LJ.LJ",
			}),
			16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.loopLength(); got != tt.want {
				t.Errorf("sketch.loopLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSketchMaxDistance(t *testing.T) {
	tests := []struct {
		name string
		s    sketch
		want int
	}{
		{
			"circular",
			newSketch([]string{
				"-L|F7",
				"7S-7|",
				"L|7||",
				"-L-J|",
				"L|-JF",
			}),
			4,
		},
		{
			"complex",
			newSketch([]string{
				"7-F7-",
				".FJ|7",
				"SJLL7",
				"|F--J",
				"LJ.LJ",
			}),
			8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.MaxDistance(); got != tt.want {
				t.Errorf("sketch.MaxDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
