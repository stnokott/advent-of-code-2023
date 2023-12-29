// Package main runs the input for Day 9
package main

import (
	"reflect"
	"testing"
)

func TestDiffs(t *testing.T) {
	type args struct {
		vals []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"all equal", args{[]int{3, 3, 3, 3, 3}}, []int{0, 0, 0, 0}},
		{"+1", args{[]int{6, 7, 8, 9, 10}}, []int{1, 1, 1, 1}},
		{"+3+4+5", args{[]int{4, 7, 11, 16, 22}}, []int{3, 4, 5, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := diffs(tt.args.vals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("diffs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNext(t *testing.T) {
	type args struct {
		vals []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"depth 1", args{[]int{0, 3, 6, 9, 12, 15, 18}}, 21}, // +3 +3 +3 ...
		{"depth 2", args{[]int{1, 3, 6, 10, 15, 21}}, 28},    // +2 +3 +4 ...
		{"depth 3", args{[]int{10, 13, 16, 21, 30, 45}}, 68}, // +3 +3 +5 +9 +15 ...
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := next(tt.args.vals); got != tt.want {
				t.Errorf("next() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParse(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"simple", args{"4 19 20 89"}, []int{4, 19, 20, 89}},
		{"one number", args{"19"}, []int{19}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parse(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumPredictions(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"official",
			args{[]string{
				"0 3 6 9 12 15",
				"1 3 6 10 15 21",
				"10 13 16 21 30 45",
			}},
			114,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumPredictions(tt.args.lines); got != tt.want {
				t.Errorf("SumPredictions() = %v, want %v", got, tt.want)
			}
		})
	}
}
