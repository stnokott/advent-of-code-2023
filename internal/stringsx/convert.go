package stringsx

import (
	"strconv"
	"strings"
)

// MustAtoi converts a trimmed string to a number.
// It panics if an errors occurs.
func MustAtoi(s string) int {
	n, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		panic(err)
	}
	return n
}
