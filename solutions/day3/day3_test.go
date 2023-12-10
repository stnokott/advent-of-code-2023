package main

import (
	"reflect"
	"testing"
)

func TestNewSchematic(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		args args
		want *Schematic
	}{
		{"simple", args{[]string{"123", "456", "789"}}, &Schematic{numRows: 3, numCols: 3, lines: []string{"123", "456", "789"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSchematic(tt.args.lines); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSchematic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchematicHasAdjacentSymbol(t *testing.T) {
	s := NewSchematic(
		[]string{
			"467..114..",
			"...*......",
			"..35..633.",
			"......#...",
			"617*......",
			".....+.58.",
			"..592.....",
			"......755.",
			"...$.*....",
			".664.598..",
		},
	)

	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"right", args{2, 4}, true},
		{"top", args{3, 2}, true},
		{"bottom", args{6, 2}, true},
		{"diagonal", args{2, 0}, true},
		{"none", args{5, 0}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := s.hasAdjacentSymbol(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Schematic.hasAdjacentSymbol() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchematicSolve(t *testing.T) {
	tests := []struct {
		name  string
		lines []string
		want  int
	}{
		{"simple", []string{"..123..#456.7"}, 456},
		{"at start", []string{"123#..."}, 123},
		{"at end", []string{"...#123"}, 123},
		{"multirow", []string{".123.", ".#.5$."}, 128},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSchematic(tt.lines)
			if got := s.solve(); got != tt.want {
				t.Errorf("Schematic.solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
