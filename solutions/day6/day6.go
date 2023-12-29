// Package main runs the input for Day 6
package main

import (
	"math"
	"regexp"

	"github.com/advent-of-code-2023/internal/strconvx"
)

var regexNumbers = regexp.MustCompile(`(\d+)`)

func parseRaces(lines []string) []Race {
	times := regexNumbers.FindAllString(lines[0], -1)
	distances := regexNumbers.FindAllString(lines[1], len(times))

	races := make([]Race, len(times))
	for i := 0; i < len(races); i++ {
		races[i] = Race{
			time:     strconvx.MustAtoi(times[i]),
			distance: strconvx.MustAtoi(distances[i]),
		}
	}
	return races
}

var regexNonNumbers = regexp.MustCompile(`(\D)`)

func parseRace(lines []string) Race {
	time := regexNonNumbers.ReplaceAllString(lines[0], "")
	distance := regexNonNumbers.ReplaceAllString(lines[1], "")
	return Race{
		time:     strconvx.MustAtoi(time),
		distance: strconvx.MustAtoi(distance),
	}
}

// Race contains data about a singular race
type Race struct {
	time     int // total time the race takes
	distance int // current distance record which needs to be beaten
}

func distance(total int, push int) int {
	return (total - push) * push
}

func (r Race) numSolutions() int {
	// Results (distances) are spread across push times like a gauss/bell curve.
	// This means that our distances have two important properties:
	//   - until the "middle" of the bell curve, values only increase
	//   - values are "mirrored" at the middle of the bell curve, so only decrease after the middle
	middle := int(math.Floor(float64(r.time) / 2))

	// We start at push time 1 (since 0 cannot be a winning input).
	// We iterate until the middle of the bell curve, knowing that output values are only mirrored after that.
	for n := 1; n <= middle; n++ {
		// Check if we can beat the record
		if distance(r.time, n) > r.distance {
			// When we encounter the first record-beating distance, we can calculate the total number of record-beating inputs.
			// Since we knwo that there will only be higher values from here (n) until the middle of the bell curve, we can
			// take the number of inputs from n to middle and double it to get the total.
			result := (middle - n + 1) * 2
			if r.time%2 == 0 {
				// even total time means we have only one instead of two "middle" values,
				// so we subtract one from the result in that case.
				result--
			}
			return result
		}
	}
	return 0
}

func raceMultiplications(races []Race) int {
	x := 1
	for _, r := range races {
		x *= r.numSolutions()
	}
	return x
}
