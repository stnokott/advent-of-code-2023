// Package slicesx provides high-level utility functions for slice/array operations
package slicesx

import (
	"cmp"
	"slices"
)

// Intersect returns a slice of elements present in a and b.
//
// This function is considerably faster than IntersectBrute when dealing with larger datasets,
// but more memory-efficient with smaller datasets.
//
// See benchmarks for details.
func Intersect[S ~[]T, T cmp.Ordered](a S, b S) []T {
	if len(a) == 0 || len(b) == 0 {
		return S{}
	}
	// find input slice with min length
	var x, y S
	if len(a) <= len(b) {
		x, y = a, b
	} else {
		x, y = b, a
	}
	return intersectSized(x, y)
}

// intersectSized returns intersect of two slices.
// len(a) <= len(b) needs to be true.
//
// Example input:  [41 48 83 86 17], [83 86 6 31 17 9 48 53]
// Example output: [17 48 83 86]
func intersectSized[S ~[]T, T cmp.Ordered](a S, b S) S {
	// [41 48 83 86 17], [83 86 6 31 17 9 48 53]
	slices.Sort(a)
	slices.Sort(b)
	// [17 41 48 83 86], [6 9 17 31 48 53 83 86]
	intersect := make(S, 0, len(b)) // pre-allocate output to size of larger slice

	// iterate smaller slice
	ib := 0
	for _, n := range a {
		for n > b[ib] {
			ib++
			if ib > len(b)-1 {
				// end of larger slice reached, can return now
				return intersect
			}
		}
		if n == b[ib] {
			intersect = append(intersect, n)
			//ib++
		}
	}
	// end of smaller slice reached
	return intersect
}

// IntersectBrute returns a slice of elements present in a and b.
//
// This function is considerably more memory-efficient than Intersect when dealing with larger datasets,
// but faster with smaller datasets.
//
// See benchmarks for details.
func IntersectBrute[S ~[]T, T cmp.Ordered](a S, b S) []T {
	if len(a) == 0 || len(b) == 0 {
		return S{}
	}
	return intersectSizedBrute(a, b)
}

func intersectSizedBrute[S ~[]T, T cmp.Ordered](a S, b S) S {
	intersect := make(S, 0) // pre-allocate output to size of larger slice
	for _, x := range a {
		for _, y := range b {
			if x == y && !slices.Contains(intersect, x) {
				intersect = append(intersect, x)
			}
		}
	}
	return intersect
}

// All returns true if qualififerFunc returns true for all items in x.
func All[S ~[]T, T any](x S, qualifierFunc func(T) bool) bool {
	for _, el := range x {
		if !qualifierFunc(el) {
			return false
		}
	}
	return true
}

func Reduce[S ~[]T, T any, V any](x S, reduceFunc func(acc V, val T) V, initial V) V {
	result := initial
	for _, el := range x {
		result = reduceFunc(result, el)
	}
	return result
}
