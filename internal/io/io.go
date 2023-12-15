// Package io provides utility functions for I/O operations
package io

import (
	"fmt"
	"os"
	"strings"
)

// ReadLines fully reads the file at path and returns a list of its lines.
func ReadLines(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	s := strings.TrimSpace(string(data))
	fmt.Println("read", len(data), "bytes")
	// unify line terminators
	s = strings.ReplaceAll(s, "\r\n", "\n")
	lines := strings.Split(s, "\n")
	fmt.Println("read", len(lines), "lines")
	return lines, nil
}
