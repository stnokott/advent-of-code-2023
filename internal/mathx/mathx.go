// Package mathx provides utility functions for mathematics
package mathx

import "math"

// MinInt wraps math.Min() by converting a & b to float64 and converting the result back from float64 to int.
func MinInt(a int, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

// MaxInt wraps math.Max() by converting a & b to float64 and converting the result back from float64 to int.
func MaxInt(a int, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

// AbsInt returns the absolute value of x as integer.
func AbsInt(x int) int {
	return int(math.Abs(float64(x)))
}

// GCD returns the greatest common divisor for a and b.
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// LCM returns the lowest common multiplier for a and b.
func LCM(a, b int) int {
	return a / GCD(a, b) * b
}
