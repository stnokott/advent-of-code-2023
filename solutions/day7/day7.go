// Package main runs the input for Day 7
package main

import (
	"cmp"
	"slices"
	"strings"

	"github.com/advent-of-code-2023/internal/strconvx"
)

// Game contains the setup for this game
type Game struct {
	hands         []Hand
	jokerWildcard bool
}

// NewGame constructs a new Game instance from the parameters.
func NewGame(lines []string, jokerWildcard bool) *Game {
	hands := make([]Hand, len(lines))
	for i, line := range lines {
		split := strings.SplitN(line, " ", 2)
		hands[i] = NewHand(split[0], strconvx.MustAtoi(split[1]), jokerWildcard)
	}
	g := &Game{
		hands:         hands,
		jokerWildcard: jokerWildcard,
	}
	slices.SortFunc(g.hands, g.compareHandsFunc())
	return g
}

// Hand represents a hand of five cards.
type Hand struct {
	s   string
	t   handType
	bid int
}

// handType is the type of hand represented by 5 cards.
// Weaker hands have lower integer values, starting with 0.
type handType int

const (
	HandHighest   handType = iota // HandHighest is High card, where all cards' labels are distinct: 23456
	HandOnePair                   // HandOnePair is One pair, where two cards share one label, and the other three cards have a different label from the pair and each other: A23A4
	HandTwoPair                   // HandTwoPair is Two pair, where two cards share one label, two other cards share a second label, and the remaining card has a third label: 23432
	HandThreeKind                 // HandThreeKind is Three of a kind, where three cards have the same label, and the remaining two cards are each different from any other card in the hand: TTT98
	HandFullHouse                 // HandFullHouse is Full house, where three cards have the same label, and the remaining two cards share a different label: 23332
	HandFourKind                  // HandFourKind is Four of a kind, where four cards have the same label and one card has a different label: AA8AA
	HandFiveKind                  // HandFiveKind is Five of a kind, where all five cards have the same label: AAAAA
)

// NewHand constructs a new Hand instance from the input hand string.
// The input must have length 5.
//
// Example input: A744J
func NewHand(s string, bid int, jokerWildcard bool) Hand {
	// count occurence for each character
	charMap := make(map[rune]int, len(s))
	for _, c := range s {
		charMap[c]++
	}

	var processFunc func(charMap map[rune]int) (int, int)
	if jokerWildcard {
		processFunc = processCharMapWildcards
	} else {
		processFunc = processCharMap
	}
	minDistincts, maxCount := processFunc(charMap)

	var t handType
	switch minDistincts {
	case 0, 1:
		t = HandFiveKind
	case 2:
		if maxCount == 4 {
			t = HandFourKind
		} else {
			t = HandFullHouse
		}
	case 3:
		if maxCount == 3 {
			t = HandThreeKind
		} else {
			t = HandTwoPair
		}
	case 4:
		t = HandOnePair
	default:
		t = HandHighest
	}

	return Hand{
		s:   s,
		t:   t,
		bid: bid,
	}
}

func processCharMap(charMap map[rune]int) (minDistincts, maxCount int) {
	// calculate number of distinct characters
	minDistincts = len(charMap)
	// get the maximum occurence count of a character,
	// will be used for differentiating between hand types later
	for _, v := range charMap {
		if v > maxCount {
			maxCount = v
		}
	}
	return
}

func processCharMapWildcards(charMap map[rune]int) (minDistincts, maxCount int) {
	// calculate *minimum* number of distinct characters.
	// derivation of hand type will still work the same.
	minDistincts = len(charMap)
	// subtract 1 from minimum number of distincts if at least one joker is contained
	if charMap['J'] > 0 {
		minDistincts--
	}

	maxCount = 0
	for r, v := range charMap {
		// ignore jokers here...
		if r != 'J' && v > maxCount {
			maxCount = v
		}
	}
	// ...so that we can add the total joker occurences to maxCount now
	maxCount += charMap['J']
	return
}

const (
	cardOrder          = "23456789TJQKA" // reversed to work with strings.Index()
	cardOrderWildcards = "J23456789TQKA"
)

// compareHandsFunc returns a comparison function for Hand structs.
//
// The returned functions returns:
//
//	-1 if this x has lower value than y.
//	 0 if this x has the same value as y,
//	+1 if this x has greater value than y.
func (g *Game) compareHandsFunc() func(Hand, Hand) int {
	var order string
	if g.jokerWildcard {
		order = cardOrderWildcards
	} else {
		order = cardOrder
	}

	return func(x, y Hand) int {
		if result := cmp.Compare(x.t, y.t); result != 0 {
			return result
		}
		// equal ranks, need to check characters
		for i := range x.s {
			if result := cmp.Compare(
				strings.IndexByte(order, x.s[i]),
				strings.IndexByte(order, y.s[i]),
			); result != 0 {
				return result
			}
		}
		return 0
	}
}

func (g *Game) totalWinnings() int {
	// assumes g.hands is sorted
	x := 0
	for i, hand := range g.hands {
		x += (i + 1) * hand.bid
	}
	return x
}
