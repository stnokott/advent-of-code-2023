// Package main runs the input for Day 13

package main

import (
	"reflect"
	"testing"
)

func TestReduceRowsVertically(t *testing.T) {
	type args struct {
		rows [][]rune
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"one digit",
			args{
				[][]rune{
					{'#', '.'},
				},
			},
			[]int{1, 0},
		},
		{
			"length three",
			args{
				[][]rune{
					{'#', '#', '.'},
					{'.', '.', '#'},
					{'#', '.', '.'},
				},
			},
			[]int{0b101, 0b001, 0b010},
		},
		{
			"empty",
			args{
				[][]rune{
					{'.', '.'},
					{'.', '.'},
				},
			},
			[]int{0, 0},
		},
		{
			"long",
			args{
				[][]rune{
					{'#', '.', '#', '.', '#', '#', '.', '.'},
					{'.', '.', '#', '#', '.', '#', '#', '.'},
					{'#', '#', '.', '.', '.', '.', '#', '.'},
					{'.', '#', '#', '.', '#', '.', '.', '.'},
					{'#', '#', '#', '#', '.', '#', '#', '.'},
					{'.', '.', '#', '.', '#', '.', '.', '#'},
				},
			},
			[]int{0b010101, 0b011100, 0b111011, 0b010010, 0b101001, 0b010011, 0b010110, 0b100000},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reduceRowsVertically(tt.args.rows); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reduceRowsVertically() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGridRotateCCW(t *testing.T) {
	type args struct {
		rows [][]rune
	}
	tests := []struct {
		name string
		args args
		want [][]rune
	}{
		{
			"simple",
			args{
				[][]rune{
					{'1', '2'},
					{'3', '4'},
				},
			},
			[][]rune{
				{'2', '4'},
				{'1', '3'},
			},
		},
		{
			"different hw",
			args{
				[][]rune{
					{'1', '2', '3'},
					{'4', '5', '6'},
				},
			},
			[][]rune{
				{'3', '6'},
				{'2', '5'},
				{'1', '4'},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gridRotateCCW(tt.args.rows); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("gridRotateCCW() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMakeGrid(t *testing.T) {
	type args struct {
		rows []string
	}
	tests := []struct {
		name string
		args args
		want [][]rune
	}{
		{
			"small",
			args{
				[]string{
					"123",
					"456",
				},
			},
			[][]rune{
				{'1', '2', '3'},
				{'4', '5', '6'},
			},
		},
		{
			"big",
			args{
				[]string{
					"abcdefgh",
					"ijklmnop",
					"qrstuvwx",
					"yz012345",
					"67890-.,",
				},
			},
			[][]rune{
				{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'},
				{'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p'},
				{'q', 'r', 's', 't', 'u', 'v', 'w', 'x'},
				{'y', 'z', '0', '1', '2', '3', '4', '5'},
				{'6', '7', '8', '9', '0', '-', '.', ','},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeGrid(tt.args.rows); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeGrid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindVerticalMirror(t *testing.T) {
	type args struct {
		rows [][]rune
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"official",
			args{
				[][]rune{
					{'#', '.', '#', '#', '.', '.', '#', '#', '.'},
					{'.', '.', '#', '.', '#', '#', '.', '#', '.'},
					{'#', '#', '.', '.', '.', '.', '.', '.', '#'},
					{'#', '#', '.', '.', '.', '.', '.', '.', '#'},
					{'.', '.', '#', '.', '#', '#', '.', '#', '.'},
					{'.', '.', '#', '#', '.', '.', '#', '#', '.'},
					{'#', '.', '#', '.', '#', '#', '.', '#', '.'},
				},
			},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findVerticalMirror(tt.args.rows); got != tt.want {
				t.Errorf("findVerticalMirror() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsMirror(t *testing.T) {
	type args struct {
		stack     []int
		remainder []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"exact mirror",
			args{
				stack:     []int{1, 2, 2, 3},
				remainder: []int{3, 2, 2, 1},
			},
			true,
		},
		{
			"additional stack",
			args{
				stack:     []int{1, 5, 7, 3},
				remainder: []int{3, 7},
			},
			true,
		},
		{
			"additional remainder",
			args{
				stack:     []int{6, 9, 3},
				remainder: []int{3, 9, 6, 4, 5},
			},
			true,
		},
		{
			"no mirror",
			args{
				stack:     []int{5, 8, 9},
				remainder: []int{9, 8, 4},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMirror(tt.args.stack, tt.args.remainder); got != tt.want {
				t.Errorf("isMirror() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindMirrors(t *testing.T) {
	type args struct {
		rows []string
	}
	tests := []struct {
		name     string
		args     args
		wantLeft int
		wantTop  int
	}{
		{
			"small",
			args{
				[]string{
					".##..#",
					"#..#..",
					".##..#",
					"#..#..",
				},
			},
			2,
			0,
		},
		{
			"official 1",
			args{
				[]string{
					"#.##..##.",
					"..#.##.#.",
					"##......#",
					"##......#",
					"..#.##.#.",
					"..##..##.",
					"#.#.##.#.",
				},
			},
			5,
			0,
		},
		{
			"official 2",
			args{
				[]string{
					"#...##..#",
					"#....#..#",
					"..##..###",
					"#####.##.",
					"#####.##.",
					"..##..###",
					"#....#..#",
				},
			},
			0,
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLeft, gotTop := FindMirrors(tt.args.rows)
			if gotLeft != tt.wantLeft {
				t.Errorf("FindMirrors() gotLeft = %v, want %v", gotLeft, tt.wantLeft)
			}
			if gotTop != tt.wantTop {
				t.Errorf("FindMirrors() gotTop = %v, want %v", gotTop, tt.wantTop)
			}
		})
	}
}
