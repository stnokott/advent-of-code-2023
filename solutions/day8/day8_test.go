// Package main runs the input for Day 8
package main

import (
	"reflect"
	"testing"
)

func TestNewTree(t *testing.T) {
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
	zzz := &Node{isEnd: true}
	zzz.Left, zzz.Right = zzz, zzz
	ggg := &Node{}
	ggg.Left, ggg.Right = ggg, ggg
	eee := &Node{}
	eee.Left, eee.Right = eee, eee
	ddd := &Node{}
	ddd.Left, ddd.Right = ddd, ddd
	ccc := &Node{Left: zzz, Right: ggg}
	bbb := &Node{Left: ddd, Right: eee}
	aaa := &Node{Left: bbb, Right: ccc}
	want := aaa

	if got := NewTree(lines); !reflect.DeepEqual(got, want) {
		t.Errorf("NewTree() = %v, want %v", got, want)
	}
}

func TestNodeWalk(t *testing.T) {
	type args struct {
		instructions *instructions
		steps        int
	}
	tests := []struct {
		name string
		n    *Node
		args args
		want int
	}{
		{
			"official 1",
			NewTree([]string{
				"AAA = (BBB, CCC)",
				"BBB = (DDD, EEE)",
				"CCC = (ZZZ, GGG)",
				"DDD = (DDD, DDD)",
				"EEE = (EEE, EEE)",
				"GGG = (GGG, GGG)",
				"ZZZ = (ZZZ, ZZZ)",
			}),
			args{
				&instructions{S: "RL"},
				0,
			},
			2,
		},
		{
			"official 2",
			NewTree([]string{
				"AAA = (BBB, BBB)",
				"BBB = (AAA, ZZZ)",
				"ZZZ = (ZZZ, ZZZ)",
			}),
			args{
				&instructions{S: "LLR"},
				0,
			},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.walk(tt.args.instructions, tt.args.steps); got != tt.want {
				t.Errorf("Node.walk() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateSteps(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"official 1",
			args{[]string{
				"RL",
				"",
				"AAA = (BBB, CCC)",
				"BBB = (DDD, EEE)",
				"CCC = (ZZZ, GGG)",
				"DDD = (DDD, DDD)",
				"EEE = (EEE, EEE)",
				"GGG = (GGG, GGG)",
				"ZZZ = (ZZZ, ZZZ)",
			}},
			2,
		},
		{
			"official 2",
			args{[]string{
				"LLR",
				"",
				"AAA = (BBB, BBB)",
				"BBB = (AAA, ZZZ)",
				"ZZZ = (ZZZ, ZZZ)",
			}},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateSteps(tt.args.lines); got != tt.want {
				t.Errorf("calculateSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}
