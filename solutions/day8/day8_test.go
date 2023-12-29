// Package main runs the input for Day 8
package main

import (
	"reflect"
	"testing"
)

func TestNewNetworkRegular(t *testing.T) {
	// testing official samples
	lines := []string{
		"AAA = (BBB, CCC)",
		"BBB = (DDD, EEE)",
		"CCC = (ZZZ, GGG)",
		"DDD = (DDD, DDD)",
		"EEE = (EEE, EEE)",
		"GGG = (GGG, GGG)",
		"ZZZ = (ZZZ, ZZZ)",
	}

	want := Network{
		net: map[string][2]string{
			"AAA": {"BBB", "CCC"},
			"BBB": {"DDD", "EEE"},
			"CCC": {"ZZZ", "GGG"},
			"DDD": {"DDD", "DDD"},
			"EEE": {"EEE", "EEE"},
			"GGG": {"GGG", "GGG"},
			"ZZZ": {"ZZZ", "ZZZ"},
		},
		startNodes: []string{"AAA"},
	}

	if got := NewNetwork(lines, startNodeQualifierRegular); !reflect.DeepEqual(got, want) {
		t.Errorf("NewNetwork() = %v, want %v", got, want)
	}
}

func TestNewNetworkGhost(t *testing.T) {
	// testing official samples
	lines := []string{
		"11A = (11B, XXX)",
		"11B = (XXX, 11Z)",
		"11Z = (11B, XXX)",
		"22A = (22B, XXX)",
		"22B = (22C, 22C)",
		"22C = (22Z, 22Z)",
		"22Z = (22B, 22B)",
		"XXX = (XXX, XXX)",
	}

	want := Network{
		net: map[string][2]string{
			"11A": {"11B", "XXX"},
			"11B": {"XXX", "11Z"},
			"11Z": {"11B", "XXX"},
			"22A": {"22B", "XXX"},
			"22B": {"22C", "22C"},
			"22C": {"22Z", "22Z"},
			"22Z": {"22B", "22B"},
			"XXX": {"XXX", "XXX"},
		},
		startNodes: []string{"11A", "22A"},
	}

	if got := NewNetwork(lines, startNodeQualifierGhost); !reflect.DeepEqual(got, want) {
		t.Errorf("NewNetwork() = %v, want %v", got, want)
	}
}

func TestNetworkWalk(t *testing.T) {
	type args struct {
		start        []string
		instructions *instructions
		isEndFunc    func(string) bool
	}
	tests := []struct {
		name string
		net  Network
		args args
		want int
	}{
		{
			"official regular 1",
			NewNetwork([]string{
				"AAA = (BBB, CCC)",
				"BBB = (DDD, EEE)",
				"CCC = (ZZZ, GGG)",
				"DDD = (DDD, DDD)",
				"EEE = (EEE, EEE)",
				"GGG = (GGG, GGG)",
				"ZZZ = (ZZZ, ZZZ)",
			}, startNodeQualifierRegular),
			args{
				[]string{"AAA"},
				&instructions{S: "RL"},
				endNodeQualifierRegular,
			},
			2,
		},
		{
			"official regular 2",
			NewNetwork([]string{
				"AAA = (BBB, BBB)",
				"BBB = (AAA, ZZZ)",
				"ZZZ = (ZZZ, ZZZ)",
			}, startNodeQualifierRegular),
			args{
				[]string{"AAA"},
				&instructions{S: "LLR"},
				endNodeQualifierRegular,
			},
			6,
		},
		{
			"official ghost",
			NewNetwork([]string{
				"11A = (11B, XXX)",
				"11B = (XXX, 11Z)",
				"11Z = (11B, XXX)",
				"22A = (22B, XXX)",
				"22B = (22C, 22C)",
				"22C = (22Z, 22Z)",
				"22Z = (22B, 22B)",
				"XXX = (XXX, XXX)",
			}, startNodeQualifierGhost),
			args{
				[]string{"11A", "22A"},
				&instructions{S: "LR"},
				endNodeQualifierGhost,
			},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.net.Walk(tt.args.instructions, tt.args.isEndFunc); got != tt.want {
				t.Errorf("Node.walk() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateSteps(t *testing.T) {
	type args struct {
		lines []string
		ghost bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"official regular 1",
			args{
				[]string{
					"RL",
					"",
					"AAA = (BBB, CCC)",
					"BBB = (DDD, EEE)",
					"CCC = (ZZZ, GGG)",
					"DDD = (DDD, DDD)",
					"EEE = (EEE, EEE)",
					"GGG = (GGG, GGG)",
					"ZZZ = (ZZZ, ZZZ)",
				},
				false,
			},
			2,
		},
		{
			"official regular 2",
			args{
				[]string{
					"LLR",
					"",
					"AAA = (BBB, BBB)",
					"BBB = (AAA, ZZZ)",
					"ZZZ = (ZZZ, ZZZ)",
				},
				false,
			},
			6,
		},
		{
			"official ghost",
			args{
				[]string{
					"LR",
					"",
					"11A = (11B, XXX)",
					"11B = (XXX, 11Z)",
					"11Z = (11B, XXX)",
					"22A = (22B, XXX)",
					"22B = (22C, 22C)",
					"22C = (22Z, 22Z)",
					"22Z = (22B, 22B)",
					"XXX = (XXX, XXX)",
				},
				true,
			},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateSteps(tt.args.lines, tt.args.ghost); got != tt.want {
				t.Errorf("calculateSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}
