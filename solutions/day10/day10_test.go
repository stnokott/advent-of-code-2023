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

func TestGridInBounds(t *testing.T) {
	g := grid[rune]{
		{'.', 'F', '-', '7', '.'},
		{'.', '|', '.', '|', '.'},
		{'.', 'L', '-', 'J', '.'},
	}

	type args struct {
		c coord
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"in bounds", args{coord{1, 2}}, true},
		{"out of bounds top left", args{coord{0, -1}}, false},
		{"out of bounds bottom right", args{coord{5, 3}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := g.inBounds(tt.args.c); got != tt.want {
				t.Errorf("grid.inBounds() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGridAt(t *testing.T) {
	g := grid[rune]{
		{'.', 'F', '-', '7', '.'},
		{'.', '|', '.', '|', '.'},
		{'.', 'L', '-', 'J', '.'},
	}

	type args struct {
		c coord
	}
	tests := []struct {
		name string
		args args
		want rune
	}{
		{"inner", args{coord{1, 2}}, 'L'},
		{"border", args{coord{0, 2}}, '.'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := g.at(tt.args.c); got != tt.want {
				t.Errorf("grid.at() = %v, want %v", got, tt.want)
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
			coord{1, 2},
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

func TestSketchLoop(t *testing.T) {
	tests := []struct {
		name string
		s    sketch
		want []coord
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
			[]coord{
				{1, 1}, {1, 2}, {1, 3}, {2, 3}, {3, 3}, {3, 2}, {3, 1}, {2, 1},
			},
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
			[]coord{
				{0, 2}, {0, 3}, {0, 4}, {1, 4}, {1, 3}, {2, 3}, {3, 3}, {4, 3},
				{4, 2}, {3, 2}, {3, 1}, {3, 0}, {2, 0}, {2, 1}, {1, 1}, {1, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.loop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sketch.loop() = %v, want %v", got, tt.want)
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

func TestSketchNumInnerCells(t *testing.T) {
	tests := []struct {
		name string
		s    sketch
		want int
	}{
		{
			"small circular",
			newSketch([]string{
				".....",
				".S-7.",
				".|.|.",
				".L-J.",
				".....",
			}),
			1,
		},
		{
			"small complex",
			newSketch([]string{
				"..F7.",
				".FJ|.",
				"SJLL7",
				"|F--J",
				"LJ...",
			}),
			1,
		},
		{
			"big simple",
			newSketch([]string{
				"...........",
				".S-------7.",
				".|F-----7|.",
				".||.....||.",
				".||.....||.",
				".|L-7.F-J|.",
				".|..|.|..|.",
				".L--J.L--J.",
				"...........",
			}),
			4,
		},
		{
			"big complex",
			newSketch([]string{
				".F----7F7F7F7F-7....",
				".|F--7||||||||FJ....",
				".||.FJ||||||||L7....",
				"FJL7L7LJLJ||LJ.L-7..",
				"L--J.L7...LJS7F-7L7.",
				"....F-J..F7FJ|L7L7L7",
				"....L7.F7||L7|.L7L7|",
				".....|FJLJ|FJ|F7|.LJ",
				"....FJL-7.||.||||...",
				"....L---J.LJ.LJLJ...",
			}),
			8,
		},
		{
			"big even more complex",
			newSketch([]string{
				"FF7FSF7F7F7F7F7F---7",
				"L|LJ||||||||||||F--J",
				"FL-7LJLJ||||||LJL-77",
				"F--JF--7||LJLJ7F7FJ-",
				"L---JF-JLJ.||-FJLJJ7",
				"|F|F-JF---7F7-L7L|7|",
				"|FFJF7L7F-JF7|JL---7",
				"7-L-JL7||F7|L7F-7F7|",
				"L.L7LFJ|||||FJL7||LJ",
				"L7JLJL-JLJLJL--JLJ.L",
			}),
			10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.NumInnerCells(); got != tt.want {
				t.Errorf("sketch.NumInnerCells() = %v, want %v", got, tt.want)
			}
		})
	}
}
