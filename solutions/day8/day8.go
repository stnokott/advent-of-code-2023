// Package main runs the input for Day 8
package main

import (
	"regexp"
)

// Node is a navigation step in the binary tree.
type Node struct {
	Left  *Node
	Right *Node
	isEnd bool
}

// NewTree constructs a new binary tree from the input lines and returns the root node ("AAA").
func NewTree(lines []string) *Node {
	nodes := make(map[string]*Node, len(lines))
	nodeGetOrDefault := func(name string) *Node {
		n := nodes[name]
		if n == nil {
			n = &Node{}
			nodes[name] = n
		}
		return n
	}

	for _, line := range lines {
		name, left, right := parseLine(line)

		node := nodeGetOrDefault(name)
		l := nodeGetOrDefault(left)
		r := nodeGetOrDefault(right)
		node.Left, node.Right = l, r
	}
	nodes["ZZZ"].isEnd = true
	return nodes["AAA"]
}

var regexIdentifier = regexp.MustCompile(`[A-Z]{3}`)

func parseLine(s string) (name, left, right string) {
	matches := regexIdentifier.FindAllString(s, 3)
	return matches[0], matches[1], matches[2]
}

type instructions struct {
	S string
	i int
}

func (in *instructions) Next() byte {
	b := in.S[in.i%len(in.S)]
	in.i++
	return b
}

func (n *Node) walk(instructions *instructions, steps int) int {
	if n.isEnd {
		return steps
	}
	var next *Node
	if instructions.Next() == 'L' {
		next = n.Left
	} else {
		next = n.Right
	}
	return next.walk(instructions, steps+1)
}

func calculateSteps(lines []string) int {
	instructions := &instructions{S: lines[0]}
	tree := NewTree(lines[2:])
	return tree.walk(instructions, 0)
}
