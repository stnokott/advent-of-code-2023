package main

import (
	"strings"

	"github.com/advent-of-code-2023/internal/stringsx"
)

// Almanach contains data for the whole almanach.
type Almanach struct {
	SeedRanges []SeedRange
	Maps       []Map // Maps contains all maps of this almanach, sorted by appearance in the input string.
}

// SeedRange maps the start of a seed range and its length to the first and second value of an array respectively.
type SeedRange [2]int

// NewAlmanach constructs a new Almanach instance from the input lines.
//
// useSeedPairs indicates whether two subsequent seed numbers in the first line should
// be treated as seed range start and seed range length respectively.
// If false, every seed value will simply be used as a regular seed number.
func NewAlmanach(lines []string, useSeedPairs bool) *Almanach {
	var seedRanges []SeedRange
	if useSeedPairs {
		seedRanges = parseSeedsAsPairs(lines[0])
	} else {
		seedRanges = parseSeedsAsValues(lines[0])
	}
	maps := make([]Map, 0)

	// start iterating from third line (first map line)
	titleLineIndex := -1
	for i := 2; i < len(lines); i++ {
		if titleLineIndex == -1 && isMapStartLine(lines[i]) {
			titleLineIndex = i
		} else if i == len(lines)-1 || lines[i+1] == "" {
			// end of end of total lines or map section reached, save current mappings
			m := NewMap(lines[titleLineIndex], lines[titleLineIndex+1:i+1])
			maps = append(maps, m)
			titleLineIndex = -1
		}
	}

	return &Almanach{
		SeedRanges: seedRanges,
		Maps:       maps,
	}
}

func isMapStartLine(s string) bool {
	return strings.HasSuffix(s, "map:")
}

func parseSeeds(s string) []int {
	seedsStr := strings.SplitN(s, " ", 2)[1]
	seedsSplit := strings.Split(seedsStr, " ")
	seeds := make([]int, len(seedsSplit))
	for i, ss := range seedsSplit {
		seeds[i] = stringsx.MustAtoi(ss)
	}
	return seeds
}

func parseSeedsAsValues(s string) []SeedRange {
	seeds := parseSeeds(s)
	pairs := make([]SeedRange, len(seeds))
	for i := range pairs {
		pairs[i] = SeedRange{
			seeds[i],
			1,
		}
	}
	return pairs
}

func parseSeedsAsPairs(s string) []SeedRange {
	seeds := parseSeeds(s)
	pairs := make([]SeedRange, int(float64(len(seeds))/2))
	for i := range pairs {
		pairs[i] = SeedRange{
			seeds[i*2],
			seeds[i*2+1],
		}
	}
	return pairs
}

// Map contains data about all range for one type of map.
type Map struct {
	Name   string
	Ranges []Range
}

// NewMap constructs a Map instance (not map) from the input lines.
//
// Example input: "soil-to-fertilizer map:" ["0 15 37" "37 52 2" "39 0 15"]
// Example output: *Map{Name: "soil-to-fertilizer", Ranges: [...]}
func NewMap(titleLine string, mapLines []string) Map {
	name := strings.SplitN(titleLine, " ", 2)[0]
	ranges := make([]Range, len(mapLines))
	for i := range ranges {
		ranges[i] = NewRange(mapLines[i])
	}
	return Map{
		Name:   name,
		Ranges: ranges,
	}
}

// Range contains data about one mapping range.
type Range struct {
	SrcStart  int
	SrcEnd    int
	DstOffset int
}

// NewRange constructs a Range instance from an input string.
//
// Example input:  0 15 37
// Example output: *Range{SrcStart: 15, SrcEnd: 52, DstOffset: -15}
func NewRange(s string) Range {
	parts := strings.Split(s, " ") // len(parts) == 3
	dstStart := stringsx.MustAtoi(parts[0])
	srcStart := stringsx.MustAtoi(parts[1])
	length := stringsx.MustAtoi(parts[2])
	return Range{
		SrcStart:  srcStart,
		SrcEnd:    srcStart + length,
		DstOffset: dstStart - srcStart,
	}
}
