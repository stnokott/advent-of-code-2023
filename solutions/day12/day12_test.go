// Package main runs the input for Day 12
package main

import (
	"reflect"
	"testing"
)

func TestValid(t *testing.T) {
	type args struct {
		s      string
		target []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"valid trailing #",
			args{
				".###..##..#",
				[]int{3, 2, 1},
			},
			true,
		},
		{
			"valid trailing .",
			args{
				".#...#....###.",
				[]int{1, 1, 3},
			},
			true,
		},
		{
			"impossible first",
			args{
				".##..##..#",
				[]int{3, 2, 1},
			},
			false,
		},
		{
			"impossible last",
			args{
				".###..##..##",
				[]int{3, 2, 1},
			},
			false,
		},
		{
			"too few sequences",
			args{
				".###.#",
				[]int{3, 2, 1},
			},
			false,
		},
		{
			"too many sequences",
			args{
				"..##..####.#.#",
				[]int{2, 4, 1},
			},
			false,
		},
		{
			"only empty",
			args{
				"...",
				[]int{1},
			},
			false,
		},
		{
			"empty input",
			args{
				"",
				[]int{3, 2, 1},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := valid(tt.args.s, tt.args.target); got != tt.want {
				t.Errorf("valid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkValid(b *testing.B) {
	for n := 0; n < b.N; n++ {
		valid("...#.##.##.##.#.#....######..###..###.", []int{1, 2, 2, 2, 1, 1, 6, 3, 3})
	}
}

func TestPermutations(t *testing.T) {
	type args struct {
		s      string
		i      int
		target []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"???.###",
			args{
				"???.###",
				0,
				[]int{1, 1, 3},
			},
			1,
		},
		{
			".??..??...?##.",
			args{
				".??..??...?##.",
				0,
				[]int{1, 1, 3},
			},
			4,
		},
		{
			"?###????????",
			args{
				"?###????????",
				0,
				[]int{3, 2, 1},
			},
			10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permutations(tt.args.s, tt.args.i, tt.args.target); got != tt.want {
				t.Errorf("permutations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseLine(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name       string
		args       args
		wantString string
		wantTarget []int
	}{
		{
			"short",
			args{
				"#.#.### 1,1,3",
			},
			"#.#.###",
			[]int{1, 1, 3},
		},
		{
			"complex",
			args{
				"##..##..##..##..#. 12,99,420",
			},
			"##..##..##..##..#.",
			[]int{12, 99, 420},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := parseLine(tt.args.s)
			if got != tt.wantString {
				t.Errorf("parseLine() got = %v, want %v", got, tt.wantString)
			}
			if !reflect.DeepEqual(got1, tt.wantTarget) {
				t.Errorf("parseLine() got1 = %v, want %v", got1, tt.wantTarget)
			}
		})
	}
}

func TestSumPermutations(t *testing.T) {
	lines := []string{
		"???.### 1,1,3",
		".??..??...?##. 1,1,3",
		"?#?#?#?#?#?#?#? 1,3,1,6",
		"????.#...#... 4,1,1",
		"????.######..#####. 1,6,5",
		"?###???????? 3,2,1",
	}
	want := 21

	if got := SumPermutations(lines); got != want {
		t.Errorf("SumPermutations() = %v, want %v", got, want)
	}
}
