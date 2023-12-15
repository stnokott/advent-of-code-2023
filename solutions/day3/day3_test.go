package main

import (
	"reflect"
	"testing"
)

func TestElementsInLine(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want []*Element
	}{
		{"only_number", args{"123"}, []*Element{{s: "123", x0: 0, isNumber: true}}},
		{"number_and_dots", args{"..456.."}, []*Element{{s: "456", x0: 2, isNumber: true}}},
		{"number_and_symbols", args{".42$."}, []*Element{{s: "42", x0: 1, isNumber: true}, {s: "$", x0: 3, isNumber: false}}},
		{"only_symbols", args{"$$%"}, []*Element{{s: "$$%", x0: 0, isNumber: false}}},
		{"symbols_and_dots", args{"....%&.."}, []*Element{{s: "%&", x0: 4, isNumber: false}}},
		{"symbols_and_numbers", args{".$%/999"}, []*Element{{s: "$%/", x0: 1, isNumber: false}, {s: "999", x0: 4, isNumber: true}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := elementsInLine(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("elementsInLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSchematic(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		args args
		want Schematic
	}{
		{
			"combined",
			args{[]string{"123", "456", "..57..$8", "789"}},
			[][]*Element{
				{
					{s: "123", x0: 0, isNumber: true},
				},
				{
					{s: "456", x0: 0, isNumber: true},
				},
				{
					{s: "57", x0: 2, isNumber: true},
					{s: "$", x0: 6, isNumber: false},
					{s: "8", x0: 7, isNumber: true},
				},
				{
					{s: "789", x0: 0, isNumber: true},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSchematic(tt.args.lines); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSchematic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchematicGetAdjacentElements(t *testing.T) {
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
		el *Element
		y  int
	}
	tests := []struct {
		name string
		args args
		want []*Element
	}{
		{"right", args{s[4][0], 4}, []*Element{s[4][1]}},
		{"top", args{s[2][0], 2}, []*Element{s[1][0]}},
		{"bottom", args{s[2][1], 2}, []*Element{s[3][0]}},
		{"diagonal", args{s[0][0], 0}, []*Element{s[1][0]}},
		{"none", args{s[0][1], 0}, []*Element{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := s.getAdjacentElements(tt.args.el, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Schematic.getAdjacentElements() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolveParts(t *testing.T) {
	type args struct {
		s Schematic
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"simple", args{NewSchematic([]string{"..123..#456.7"})}, 456},
		{"at_start", args{NewSchematic([]string{"123#..."})}, 123},
		{"at_end", args{NewSchematic([]string{"...#123"})}, 123},
		{"multirow", args{NewSchematic([]string{".123.", ".#.5$."})}, 128},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveParts(tt.args.s); got != tt.want {
				t.Errorf("solveParts() = %v, want %v", got, tt.want)
			}
		})
	}
}
