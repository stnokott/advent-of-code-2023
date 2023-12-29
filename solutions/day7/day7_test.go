// Package main runs the input for Day 7
package main

import (
	"reflect"
	"testing"
)

func TestNewHandNoWildcards(t *testing.T) {
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
				if got := NewHand(arg, 0, false); !reflect.DeepEqual(got.t, tt.want) {
					t.Errorf("NewHand().t = %v, want %v", got.t, tt.want)
				}
			})
		}
	}
}

func TestNewHandWithWildcards(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want handType
	}{
		{"Five of a kind", []string{"AAAAA", "JJJJJ", "22222", "33333", "JJJ2J"}, HandFiveKind},
		{"Four of a kind", []string{"AAAA8", "22322", "34333", "5TTTT", "45JJ4", "A2J2J", "KTJJT"}, HandFourKind},
		{"Full House", []string{"TT8T8", "33222", "T4T44", "A2A2A", "A2A2J"}, HandFullHouse},
		{"Three of a kind", []string{"18828", "43222", "T4A44", "6Q626", "6QJ2J"}, HandThreeKind},
		{"Two pairs", []string{"TTQQ8", "2A2BB", "44TQT", "93932"}, HandTwoPair},
		{"One pair", []string{"TT123", "33742", "TQTA4", "A2258", "A2J58"}, HandOnePair},
		{"Highest card", []string{"12345", "A9257", "Q9A42"}, HandHighest},
	}
	for _, tt := range tests {
		for _, arg := range tt.args {
			t.Run(tt.name+" "+arg, func(t *testing.T) {
				if got := NewHand(arg, 0, true); !reflect.DeepEqual(got.t, tt.want) {
					t.Errorf("NewHand().t = %v, want %v", got.t, tt.want)
				}
			})
		}
	}
}

func BenchmarkNewHandFiveKind(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NewHand("AAAAA", 0, false)
	}
}

func BenchmarkNewHandFourKind(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NewHand("34333", 0, false)
	}
}

func BenchmarkNewHandFullHouse(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NewHand("A2A2A", 0, false)
	}
}

func BenchmarkNewHandThreeKind(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NewHand("43222", 0, false)
	}
}

func BenchmarkNewHandTwoPair(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NewHand("44TQT", 0, false)
	}
}

func BenchmarkNewHandHighest(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NewHand("AJ257", 0, false)
	}
}

func BenchmarkNewHandOnePair(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NewHand("A2258", 0, false)
	}
}

func TestCompareHandsNoWildcard(t *testing.T) {
	g := Game{jokerWildcard: false}

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
			fn := g.compareHandsFunc()
			if got := fn(tt.x, tt.y); got != tt.want {
				t.Errorf("CompareHands() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompareHandsWithWildcard(t *testing.T) {
	g := Game{jokerWildcard: true}

	tests := []struct {
		name string
		x    Hand
		y    Hand
		want int
	}{
		{"lower kind", NewHand("5JK39", 0, true), NewHand("TTAA3", 0, true), -1},
		{"higher kind", NewHand("AA3AA", 0, true), NewHand("666AJ", 0, true), 1},
		{"equal kind, all equal cards", NewHand("AAAAA", 0, true), NewHand("AAAAA", 0, true), 0},
		{"equal kind, all lower cards", NewHand("22222", 0, true), NewHand("33333", 0, true), -1},
		{"equal kind, all higher cards", NewHand("77777", 0, true), NewHand("33333", 0, true), 1},
		{"equal kind, different card order lower", NewHand("AJJ5A", 0, true), NewHand("AJ5JA", 0, true), -1},
		{"equal kind, different card order higher", NewHand("Q5666", 0, true), NewHand("Q5J65", 0, true), 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fn := g.compareHandsFunc()
			if got := fn(tt.x, tt.y); got != tt.want {
				t.Errorf("CompareHands() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewGame(t *testing.T) {
	type args struct {
		lines         []string
		jokerWildcard bool
	}
	tests := []struct {
		name string
		args args
		want *Game
	}{
		{
			"official no wildcards",
			args{[]string{"32T3K 765", "T55J5 684", "KK677 28", "KTJJT 220", "QQQJA 483"}, false},
			&Game{
				hands: []Hand{
					{"32T3K", HandOnePair, 765},
					{"KTJJT", HandTwoPair, 220},
					{"KK677", HandTwoPair, 28},
					{"T55J5", HandThreeKind, 684},
					{"QQQJA", HandThreeKind, 483},
				},
				jokerWildcard: false,
			},
		},
		{
			"official with wildcards",
			args{[]string{"32T3K 765", "T55J5 684", "KK677 28", "KTJJT 220", "QQQJA 483"}, true},
			&Game{
				hands: []Hand{
					{"32T3K", HandOnePair, 765},
					{"KK677", HandTwoPair, 28},
					{"T55J5", HandFourKind, 684},
					{"QQQJA", HandFourKind, 483},
					{"KTJJT", HandFourKind, 220},
				},
				jokerWildcard: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGame(tt.args.lines, tt.args.jokerWildcard); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGame() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTotalWinnings(t *testing.T) {
	tests := []struct {
		name string
		g    *Game
		want int
	}{
		{
			"official",
			&Game{
				hands: []Hand{
					{"32T3K", HandOnePair, 765},
					{"KTJJT", HandTwoPair, 220},
					{"KK677", HandTwoPair, 28},
					{"T55J5", HandThreeKind, 684},
					{"QQQJA", HandThreeKind, 483},
				},
				jokerWildcard: false,
			},
			6440,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.totalWinnings(); got != tt.want {
				t.Errorf("totalWinnings() = %v, want %v", got, tt.want)
			}
		})
	}
}
