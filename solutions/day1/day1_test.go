package main

import (
	"reflect"
	"testing"
)

func TestNumbersInString(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []string
	}{
		{"simple", "123", []string{"1", "2", "3"}},
		{"words", "one2three", []string{"one", "2", "three"}},
		{"words overlapping", "1eightwo3", []string{"1", "eight", "two", "3"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numbersInString(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("numbersInString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseNumber(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"character1", "1", 1},
		{"character2", "9", 9},
		{"word1", "two", 2},
		{"word2", "eight", 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseNumber(tt.s); got != tt.want {
				t.Errorf("parseNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExtractNumber(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"characters", "123", 13},
		{"words", "twothreefour", 24},
		{"combined", "six7eight", 68},
		{"single", "nine", 99},
		{"overlapping", "eightwo", 82},
		{"official1", "1abc2", 12},
		{"official2", "two1nine", 29},
		{"official3", "eightwothree", 83},
		{"official4", "xtwone3four", 24},
		{"official5", "zoneight234", 14},
		{"official6", "7pqrstsixteen", 76},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extractNumber(tt.s); got != tt.want {
				t.Errorf("extractNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolve(t *testing.T) {
	tests := []struct {
		name         string
		calibrations []string
		want         int
	}{
		{
			"official1",
			[]string{
				"1abc2",
				"pqr3stu8vwx",
				"a1b2c3d4e5f",
				"treb7uchet",
			},
			142,
		},
		{
			"official2",
			[]string{
				"two1nine",
				"eightwothree",
				"abcone2threexyz",
				"xtwone3four",
				"4nineeightseven2",
				"zoneight234",
				"7pqrstsixteen",
			},
			281,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.calibrations...); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
