// Package main runs the input for Day 4
package main

import (
	"reflect"
	"testing"
)

func TestParseNumbersString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"empty", args{""}, []int{}},
		{"one_element", args{"42"}, []int{42}},
		{"multiple", args{"42 123 9 44"}, []int{42, 123, 9, 44}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseNumbersString(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseNumbersString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExtractNumbers(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name     string
		args     args
		wantWins []int
		wantHave []int
	}{
		{"one_element", args{"Card 1: 41 | 44"}, []int{41}, []int{44}},
		{"multiple", args{"Card 99: 1 22 33 42 8 | 98 889 12 2"}, []int{1, 22, 33, 42, 8}, []int{98, 889, 12, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotWins, gotHave := extractNumbers(tt.args.s)
			if !reflect.DeepEqual(gotWins, tt.wantWins) {
				t.Errorf("extractNumbers() gotWins = %v, want %v", gotWins, tt.wantWins)
			}
			if !reflect.DeepEqual(gotHave, tt.wantHave) {
				t.Errorf("extractNumbers() gotHave = %v, want %v", gotHave, tt.wantHave)
			}
		})
	}
}

func TestCardPointAccu(t *testing.T) {
	type args struct {
		numWins int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"zero", args{0}, 0},
		{"one", args{1}, 1},
		{"two", args{2}, 2},
		{"four", args{4}, 8},
		{"high", args{8}, 128},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cardPointAccu(tt.args.numWins); got != tt.want {
				t.Errorf("cardPointAccu() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumCardPoints(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"simple_no_match", args{[]string{"Card 1: 12 | 34"}}, 0},
		{"simple_match", args{[]string{"Card 1: 12 | 12"}}, 1},
		{"multiple_matches", args{[]string{"Card 12: 7 12 56 1 | 98 2 1 12 13"}}, 2},
		{"multiple_rows", args{[]string{"Card 71: 5 32 18 56 2 | 65 32 56 2", "Card 73: 6 99 3 | 3 7 44 7 58 33"}}, 4 + 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumCardPoints(tt.args.lines); got != tt.want {
				t.Errorf("sumCardPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
