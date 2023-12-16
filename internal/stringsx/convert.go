package stringsx

import "strconv"

// MustAtoi wraps strconv.Atoi and panics if it errors
func MustAtoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}
