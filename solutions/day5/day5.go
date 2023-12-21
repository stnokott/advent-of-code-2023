// Package main runs the input for Day 5
package main

import (
	"math"

	"github.com/advent-of-code-2023/internal/mathx"
)

// LowestLocation returns the lowest location number corresponding to any of the seeds.
func (a *Almanach) LowestLocation() int {
	lowest := math.MaxInt

	for _, seedRange := range a.SeedRanges {
		// note: could be parallelized for faster results when using pairs.
		// would involve channels and early cancellation of goroutines, but it's not 21st and I'm just now solving day 5,
		// so I will postpone that to another date.
		for seedStart, seedOffset := seedRange[0], 0; seedOffset < seedRange[1]; seedOffset++ {
			dst := seedStart + seedOffset
			// traverse maps
			for _, mapping := range a.Maps {
				dst = mapping.Go(dst)
			}
			// last map traversed, so dst is now the location number
			lowest = mathx.MinInt(lowest, dst)
		}
	}
	return lowest
}

// Go traverses the map for the source value.
//
// It does so by traversing all ranges and returning the destination if a source range matches.
// If no matching ranges are found, the source value is returned as destination value.
func (m Map) Go(src int) int {
	// check if explicitly mapped
	for _, r := range m.Ranges {
		if src >= r.SrcStart && src <= r.SrcEnd {
			return src + r.DstOffset
		}
	}
	// if not mapped explicitly, src maps to the same value in dst
	return src
}
