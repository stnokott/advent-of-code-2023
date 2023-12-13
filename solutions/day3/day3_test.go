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
		{"only_number", args{"123"}, []*Element{{s: "123", xStart: 0, isNumber: true}}},
		{"number_and_dots", args{"..456.."}, []*Element{{s: "456", xStart: 2, isNumber: true}}},
		{"number_and_symbols", args{".42$."}, []*Element{{s: "42", xStart: 1, isNumber: true}, {s: "$", xStart: 3, isNumber: false}}},
		{"only_symbols", args{"$$%"}, []*Element{{s: "$$%", xStart: 0, isNumber: false}}},
		{"symbols_and_dots", args{"....%&.."}, []*Element{{s: "%&", xStart: 4, isNumber: false}}},
		{"symbols_and_numbers", args{".$%/999"}, []*Element{{s: "$%/", xStart: 1, isNumber: false}, {s: "999", xStart: 4, isNumber: true}}},
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
					{s: "123", xStart: 0, isNumber: true},
				},
				{
					{s: "456", xStart: 0, isNumber: true},
				},
				{
					{s: "57", xStart: 2, isNumber: true},
					{s: "$", xStart: 6, isNumber: false},
					{s: "8", xStart: 7, isNumber: true},
				},
				{
					{s: "789", xStart: 0, isNumber: true},
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

func TestSchematicSumLine(t *testing.T) {
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
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"number", args{6}, 592},
		{"multiple_numbers", args{0}, 467},
		{"symbol", args{1}, 0},
		{"multiple_symbols", args{8}, 0},
		{"numbers_and_symbols_invalid", args{5}, 0},
		{"numbers_and_symbols_valid", args{4}, 617},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := s.sumLine(tt.args.y); got != tt.want {
				t.Errorf("Schematic.sumLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolve(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"simple", args{[]string{"..123..#456.7"}}, 456},
		{"at_start", args{[]string{"123#..."}}, 123},
		{"at_end", args{[]string{"...#123"}}, 123},
		{"multirow", args{[]string{".123.", ".#.5$."}}, 128},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.lines); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
