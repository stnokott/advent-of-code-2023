// Package main runs the input for Day 11
package main

import (
	"reflect"
	"testing"
)

func TestNewSky(t *testing.T) {
	lines := []string{
		"...#......",
		".......#..",
		"#.........",
		"..........",
		"......#...",
		".#........",
		".........#",
		"..........",
		".......#..",
		"#...#.....",
	}
	want := Sky{
		galaxies: []coord{
			{3, 0}, {7, 1}, {0, 2}, {6, 4}, {1, 5}, {9, 6}, {7, 8}, {0, 9}, {4, 9},
		},
		rowsEmpty: []bool{false, false, false, true, false, false, false, true, false, false},
		colsEmpty: []bool{false, false, true, false, false, true, false, false, true, false},
	}

	if got := NewSky(lines); !reflect.DeepEqual(got, want) {
		t.Errorf("NewSky() = %v, want %v", got, want)
	}
}

func TestSkyDist(t *testing.T) {
	sky := NewSky([]string{
		"...#......",
		".......#..",
		"#.........",
		"..........",
		"......#...",
		".#........",
		".........#",
		"..........",
		".......#..",
		"#...#.....",
	})
	/*
	 ...1......
	 .......2..
	 3.........
	 ..........
	 ......4...
	 .5........
	 .........6
	 ..........
	 .......7..
	 8...9.....
	*/

	type args struct {
		a               coord
		b               coord
		emptyMultiplier int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"diagonal", args{coord{1, 5}, coord{4, 9}, 2}, 9},                // 5 to 9
		{"diagonal reverse", args{coord{4, 9}, coord{1, 5}, 2}, 9},        // 9 to 5
		{"diagonal long distance", args{coord{0, 2}, coord{9, 6}, 2}, 17}, // 3 to 6
		{"vertical", args{coord{0, 2}, coord{1, 5}, 2}, 5},                // 3 to 5
		{"vertical reverse", args{coord{1, 5}, coord{0, 2}, 2}, 5},        // 5 to 3
		{"vertical straight", args{coord{0, 2}, coord{0, 9}, 2}, 9},       // 3 to 8
		{"horizontal", args{coord{1, 5}, coord{9, 6}, 2}, 12},             // 5 to 6
		{"horizontal reverse", args{coord{9, 6}, coord{1, 5}, 2}, 12},     // 6 to 5
		{"horizontal straight", args{coord{0, 9}, coord{4, 9}, 2}, 5},     // 8 to 9
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sky.dist(tt.args.a, tt.args.b, tt.args.emptyMultiplier); got != tt.want {
				t.Errorf("dist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSkyMakePairs(t *testing.T) {
	sky := NewSky([]string{
		"...#......",
		".......#..",
		"#.........",
		"..........",
		"......#...",
		".#........",
		".........#",
		"..........",
		".......#..",
		"#...#.....",
	})
	// galaxies: {3, 0}, {7, 1}, {0, 2}, {6, 4}, {1, 5}, {9, 6}, {7, 8}, {0, 9}, {4, 9}
	want := [][2]coord{
		{{3, 0}, {7, 1}},
		{{3, 0}, {0, 2}},
		{{3, 0}, {6, 4}},
		{{3, 0}, {1, 5}},
		{{3, 0}, {9, 6}},
		{{3, 0}, {7, 8}},
		{{3, 0}, {0, 9}},
		{{3, 0}, {4, 9}},
		{{7, 1}, {0, 2}},
		{{7, 1}, {6, 4}},
		{{7, 1}, {1, 5}},
		{{7, 1}, {9, 6}},
		{{7, 1}, {7, 8}},
		{{7, 1}, {0, 9}},
		{{7, 1}, {4, 9}},
		{{0, 2}, {6, 4}},
		{{0, 2}, {1, 5}},
		{{0, 2}, {9, 6}},
		{{0, 2}, {7, 8}},
		{{0, 2}, {0, 9}},
		{{0, 2}, {4, 9}},
		{{6, 4}, {1, 5}},
		{{6, 4}, {9, 6}},
		{{6, 4}, {7, 8}},
		{{6, 4}, {0, 9}},
		{{6, 4}, {4, 9}},
		{{1, 5}, {9, 6}},
		{{1, 5}, {7, 8}},
		{{1, 5}, {0, 9}},
		{{1, 5}, {4, 9}},
		{{9, 6}, {7, 8}},
		{{9, 6}, {0, 9}},
		{{9, 6}, {4, 9}},
		{{7, 8}, {0, 9}},
		{{7, 8}, {4, 9}},
		{{0, 9}, {4, 9}},
	}
	if got := sky.makePairs(); !reflect.DeepEqual(got, want) {
		t.Errorf("sky.makePairs() = %v, want %v", got, want)
	}
}

func TestSkySumDistances(t *testing.T) {
	type args struct {
		emptyMultiplier int
	}
	tests := []struct {
		name string
		s    Sky
		args args
		want int
	}{
		{
			"official",
			NewSky([]string{
				"...#......",
				".......#..",
				"#.........",
				"..........",
				"......#...",
				".#........",
				".........#",
				"..........",
				".......#..",
				"#...#.....",
			}),
			args{emptyMultiplier: 2},
			374,
		},
		{
			"official part 2 mult 10",
			NewSky([]string{
				"...#......",
				".......#..",
				"#.........",
				"..........",
				"......#...",
				".#........",
				".........#",
				"..........",
				".......#..",
				"#...#.....",
			}),
			args{emptyMultiplier: 10},
			1030,
		},
		{
			"official part 2 mult 100",
			NewSky([]string{
				"...#......",
				".......#..",
				"#.........",
				"..........",
				"......#...",
				".#........",
				".........#",
				"..........",
				".......#..",
				"#...#.....",
			}),
			args{emptyMultiplier: 100},
			8410,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SumDistances(tt.args.emptyMultiplier); got != tt.want {
				t.Errorf("sky.SumDistances() = %v, want %v", got, tt.want)
			}
		})
	}
}
