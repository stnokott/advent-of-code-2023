// Package main runs the input for Day 7
package main

import (
	"reflect"
	"testing"
)

func TestNewHand(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want handType
	}{
		{"Five of a kind", []string{"AAAAA", "JJJJJ", "22222", "33333"}, HandFiveKind},
		{"Four of a kind", []string{"AAAA8", "JJJ2J", "22322", "34333", "5TTTT"}, HandFourKind},
		{"Full House", []string{"TT8T8", "33222", "T4T44", "A2A2A"}, HandFullHouse},
		{"Three of a kind", []string{"18828", "43222", "T4A44", "6Q626"}, HandThreeKind},
		{"Two pairs", []string{"TTQQ8", "2A2BB", "44TQT", "93932"}, HandTwoPair},
		{"One pair", []string{"TT123", "33742", "TQTA4", "A2258"}, HandOnePair},
		{"Highest card", []string{"12345", "AJ257", "Q9AJ2"}, HandHighest},
	}
	for _, tt := range tests {
		for _, arg := range tt.args {
			t.Run(tt.name+" "+arg, func(t *testing.T) {
				if got := NewHand(arg, 0); !reflect.DeepEqual(got.t, tt.want) {
					t.Errorf("NewHand().t = %v, want %v", got.t, tt.want)
				}
			})
		}
	}
}

func BenchmarkHandTypeFiveKind(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NewHand("AAAAA", 0)
	}
}

func BenchmarkHandTypeFourKind(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NewHand("34333", 0)
	}
}

func BenchmarkHandTypeFullHouse(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NewHand("A2A2A", 0)
	}
}

func BenchmarkHandTypeThreeKind(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NewHand("43222", 0)
	}
}

func BenchmarkHandTypeTwoPair(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NewHand("44TQT", 0)
	}
}

func BenchmarkHandTypeHighest(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NewHand("AJ257", 0)
	}
}

func BenchmarkHandTypeOnePair(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NewHand("A2258", 0)
	}
}

func TestCompareHands(t *testing.T) {
	tests := []struct {
		name string
		x    Hand
		y    Hand
		want int
	}{
		{"lower kind", Hand{"5JK39", HandHighest, 0}, Hand{"TTAA3", HandTwoPair, 0}, -1},
		{"higher kind", Hand{"AA3AA", HandFourKind, 0}, Hand{"666AJ", HandThreeKind, 0}, 1},
		{"equal kind, all equal cards", Hand{"AAAAA", HandFiveKind, 0}, Hand{"AAAAA", HandFiveKind, 0}, 0},
		{"equal kind, all lower cards", Hand{"22222", HandFiveKind, 0}, Hand{"33333", HandFiveKind, 0}, -1},
		{"equal kind, all higher cards", Hand{"77777", HandFiveKind, 0}, Hand{"33333", HandFiveKind, 0}, 1},
		{"equal kind, different card order lower", Hand{"Q5566", HandTwoPair, 0}, Hand{"Q5665", HandTwoPair, 0}, -1},
		{"equal kind, different card order higher", Hand{"AJJ5A", HandTwoPair, 0}, Hand{"AJ5JA", HandTwoPair, 0}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compareHands(tt.x, tt.y); got != tt.want {
				t.Errorf("CompareHands() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMakeSortedHands(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		args args
		want []Hand
	}{
		{
			"official",
			args{[]string{"32T3K 765", "T55J5 684", "KK677 28", "KTJJT 220", "QQQJA 483"}},
			[]Hand{
				{"32T3K", HandOnePair, 765},
				{"KTJJT", HandTwoPair, 220},
				{"KK677", HandTwoPair, 28},
				{"T55J5", HandThreeKind, 684},
				{"QQQJA", HandThreeKind, 483},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeSortedHands(tt.args.lines); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeSortedHands() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTotalWinnings(t *testing.T) {
	type args struct {
		sortedHands []Hand
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"official",
			args{[]Hand{
				{"32T3K", HandOnePair, 765},
				{"KTJJT", HandTwoPair, 220},
				{"KK677", HandTwoPair, 28},
				{"T55J5", HandThreeKind, 684},
				{"QQQJA", HandThreeKind, 483},
			}},
			6440,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := totalWinnings(tt.args.sortedHands); got != tt.want {
				t.Errorf("totalWinnings() = %v, want %v", got, tt.want)
			}
		})
	}
}
