// Package main runs the input for Day 13
package main

import (
	"slices"
)

// reduceRowsVertically takes each column's characters and creates a number from it.
//
// It does so by interpreting each column as a sequence of 0's ('.') and 1's ('#')
// and interpreting that as an integer.
//
// The array of each column's integer is then returned.
//
// This function serves as a simplified hashing function for columns.
func reduceRowsVertically(rows [][]rune) []int {
	ints := make([]int, len(rows[0]))
	for x := range rows[0] {
		n := 0
		for y := range rows {
			if rows[y][x] == '#' {
				// add 2^y to n
				n |= 1 << y
			}
		}
		ints[x] = n
	}
	return ints
}

func gridRotateCCW[T any](grid [][]T) [][]T {
	w, h := len(grid[0]), len(grid)
	wt, ht := h, w

	gridNew := make([][]T, ht)
	for y := range gridNew {
		rowNew := make([]T, wt)
		for x := range rowNew {
			rowNew[x] = grid[x][ht-1-y]
		}
		gridNew[y] = rowNew
	}
	return gridNew
}

func makeGrid(rows []string) [][]rune {
	grid := make([][]rune, len(rows))
	for y, row := range rows {
		rowNew := make([]rune, len(row))
		for x, c := range row {
			rowNew[x] = c
		}
		grid[y] = rowNew
	}
	return grid
}

func findVerticalMirror(rows [][]rune) int {
	// first, we calculate the "hash" of each column,
	// meaning we only have to find the mirror location for effectively one row.
	nums := reduceRowsVertically(rows)
	// now, we keep putting these numbers on a stack until we encounter the same number twice
	stack := make([]int, 0, len(nums))
	stack = append(stack, nums[0])
	for i := 1; i < len(nums); i++ {
		if nums[i] == stack[len(stack)-1] {
			// same number encountered as before
			if isMirror(stack, nums[i:]) {
				return i
			}
		}
		stack = append(stack, nums[i])
	}

	return 0
}

func isMirror(stack []int, remainder []int) bool {
	st := slices.Clone(stack) // cant modify original stack
	for _, n := range remainder {
		if len(st) == 0 {
			// full mirror found (additional characters still to the right, but don't invalidate mirror)
			return true
		}
		if n == st[len(st)-1] {
			st = st[:len(st)-1]
		} else {
			// c not matching next character on stack, not full mirror
			return false
		}
	}
	// full mirror found (additional characters still to the left, but don't invalidate mirror)
	return true
}

// FindMirrors attempts to find the first mirror both for horizontal and vertical orientation.
//
// Returns the number of columns left of the first vertical mirror and number of rows top of the first horizontal mirror.
//
// Returns -1 per orientation if no mirror found.
func FindMirrors(rows []string) (left, top int) {
	grid := makeGrid(rows)
	left = findVerticalMirror(grid)
	top = findVerticalMirror(gridRotateCCW(grid))
	return
}
