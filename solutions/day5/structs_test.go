// Package main runs the input for Day 5
package main

import (
	"reflect"
	"testing"
)

func TestNewAlmanach(t *testing.T) {
	type args struct {
		lines        []string
		useSeedPairs bool
	}
	tests := []struct {
		name string
		args args
		want *Almanach
	}{
		{
			"no ranges",
			args{
				[]string{
					"seeds: 12 3 4",
					"",
					"seed-to-soil map:",
					"50 98 2",
					"52 50 48",
					"",
					"soil-to-fertilizer map:",
					"12 5 78",
				},
				false,
			},
			&Almanach{
				SeedRanges: []SeedRange{{12, 1}, {3, 1}, {4, 1}},
				Maps: []Map{
					{
						Name: "seed-to-soil",
						Ranges: []Range{
							NewRange("50 98 2"),
							NewRange("52 50 48"),
						},
					},
					{
						Name: "soil-to-fertilizer",
						Ranges: []Range{
							NewRange("12 5 78"),
						},
					},
				},
			},
		},
		{
			"with ranges",
			args{
				[]string{
					"seeds: 12 3 4 5",
					"",
					"seed-to-soil map:",
					"50 98 2",
					"52 50 48",
					"",
					"soil-to-fertilizer map:",
					"12 5 78",
				},
				true,
			},
			&Almanach{
				SeedRanges: []SeedRange{{12, 3}, {4, 5}},
				Maps: []Map{
					{
						Name: "seed-to-soil",
						Ranges: []Range{
							NewRange("50 98 2"),
							NewRange("52 50 48"),
						},
					},
					{
						Name: "soil-to-fertilizer",
						Ranges: []Range{
							NewRange("12 5 78"),
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAlmanach(tt.args.lines, tt.args.useSeedPairs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAlmanach() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsMapStartLine(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"valid", args{"soil-to-fertilizer map:"}, true},
		{"numbers", args{"0 15 37"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMapStartLine(tt.args.s); got != tt.want {
				t.Errorf("isMapStartLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseSeeds(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"short", args{"seeds: 56"}, []int{56}},
		{"long", args{"seeds: 73 12 4 98 1076 2"}, []int{73, 12, 4, 98, 1076, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseSeeds(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseSeeds() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseSeedsAsValues(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []SeedRange
	}{
		{"short", args{"seeds: 56 12"}, []SeedRange{{56, 1}, {12, 1}}},
		{"long", args{"seeds: 73 12 4 98 1076 2"}, []SeedRange{{73, 1}, {12, 1}, {4, 1}, {98, 1}, {1076, 1}, {2, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseSeedsAsValues(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseSeedsAsValues() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseSeedsAsPairs(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []SeedRange
	}{
		{"short", args{"seeds: 56 12"}, []SeedRange{{56, 12}}},
		{"long", args{"seeds: 73 12 4 98 1076 2"}, []SeedRange{{73, 12}, {4, 98}, {1076, 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseSeedsAsPairs(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseSeedsAsPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewMap(t *testing.T) {
	type args struct {
		titleLine string
		mapLines  []string
	}
	tests := []struct {
		name string
		args args
		want Map
	}{
		{"one_line", args{"soil-to-fertilizer map:", []string{"0 15 18"}}, Map{Name: "soil-to-fertilizer", Ranges: []Range{NewRange("0 15 18")}}},
		{"multiple_lines", args{"water-to-light map:", []string{"83 109 2", "8 66 2", "7 4 1"}}, Map{Name: "water-to-light", Ranges: []Range{NewRange("83 109 2"), NewRange("8 66 2"), NewRange("7 4 1")}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMap(tt.args.titleLine, tt.args.mapLines); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewRange(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want Range
	}{
		{"equal", args{"67 67 10"}, Range{SrcStart: 67, SrcEnd: 77, DstOffset: 0}},
		{"positive_offset", args{"58 13 7"}, Range{SrcStart: 13, SrcEnd: 20, DstOffset: 45}},
		{"negative_offset", args{"14 45 60"}, Range{SrcStart: 45, SrcEnd: 105, DstOffset: -31}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRange(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
