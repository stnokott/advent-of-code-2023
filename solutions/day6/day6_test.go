// Package main runs the input for Day 6
package main

import (
	"reflect"
	"testing"
)

func TestParseRaces(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		args args
		want []Race
	}{
		{"one race", args{[]string{"Time: 45", "Distance: 700891"}}, []Race{{time: 45, distance: 700891}}},
		{"multiple races", args{[]string{"Time: 7   112  762", "Distance: 823  33   6"}}, []Race{{7, 823}, {112, 33}, {762, 6}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseRaces(tt.args.lines); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseRaces() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseRace(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		args args
		want Race
	}{
		{"one number", args{[]string{"Time: 45", "Distance: 700891"}}, Race{time: 45, distance: 700891}},
		{"multiple numbers", args{[]string{"Time: 7   112  762", "Distance: 823  33   6"}}, Race{time: 7112762, distance: 823336}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseRace(tt.args.lines); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseRace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistance(t *testing.T) {
	type args struct {
		total int
		push  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"no movement", args{total: 1, push: 1}, 0},
		{"no press", args{total: 42, push: 0}, 0},
		{"double", args{total: 2, push: 1}, 1},
		//
		{"7;0", args{total: 7, push: 0}, 0},
		{"7;1", args{total: 7, push: 1}, 6},
		{"7;2", args{total: 7, push: 2}, 10},
		{"7;3", args{total: 7, push: 3}, 12},
		{"7;4", args{total: 7, push: 4}, 12},
		{"7;5", args{total: 7, push: 5}, 10},
		{"7;6", args{total: 7, push: 6}, 6},
		{"7;7", args{total: 7, push: 7}, 0},
		//
		{"6;0", args{total: 6, push: 0}, 0},
		{"6;1", args{total: 6, push: 1}, 5},
		{"6;2", args{total: 6, push: 2}, 8},
		{"6;3", args{total: 6, push: 3}, 9},
		{"6;4", args{total: 6, push: 4}, 8},
		{"6;5", args{total: 6, push: 5}, 5},
		{"6;6", args{total: 6, push: 6}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distance(tt.args.total, tt.args.push); got != tt.want {
				t.Errorf("distance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRaceNumSolutions(t *testing.T) {
	tests := []struct {
		name string
		r    Race
		want int
	}{
		{"short", Race{time: 7, distance: 9}, 4},
		{"medium", Race{time: 15, distance: 40}, 8},
		{"long", Race{time: 30, distance: 200}, 9},
		{"very long", Race{time: 71530, distance: 940200}, 71503},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.numSolutions(); got != tt.want {
				t.Errorf("Race.numSolutions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRaceMultiplications(t *testing.T) {
	type args struct {
		races []Race
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"official sample part 1", args{[]Race{{7, 9}, {15, 40}, {30, 200}}}, 288},
		{"official sample part 2", args{[]Race{{71530, 940200}}}, 71503},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := raceMultiplications(tt.args.races); got != tt.want {
				t.Errorf("raceMultiplications() = %v, want %v", got, tt.want)
			}
		})
	}
}
