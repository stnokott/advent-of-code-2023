// Package main runs the input for Day 8
package main

import (
	"slices"
	"strings"

	"github.com/advent-of-code-2023/internal/mathx"
	"github.com/advent-of-code-2023/internal/slicesx"
)

// Network contains the network itself and its starting nodes.
type Network struct {
	net        map[string][2]string
	startNodes []string
}

type nodeQualifier func(name string) bool

// NewNetwork creates a new network from the input lines.
func NewNetwork(lines []string, isStartNodeFunc nodeQualifier) Network {
	net := make(map[string][2]string, len(lines))
	startNodes := []string{}
	for _, line := range lines {
		from, left, right := line[0:3], line[7:10], line[12:15]
		net[from] = [2]string{left, right}
		if isStartNodeFunc(from) {
			startNodes = append(startNodes, from)
		}
	}
	return Network{
		net:        net,
		startNodes: startNodes,
	}
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

// Walk applies the instructions (wrapping) to the starting nodes of the network and
// returns the number of steps taken.
func (net Network) Walk(instructions *instructions, isEndFunc nodeQualifier) int {
	nodes := slices.Clone(net.startNodes)

	steps := make([]int, 0, len(nodes))
	// continue until all current nodes qualify as end node
	for step := 0; len(nodes) > 0; {
		instruction := instructions.Next()
		// process instruction
		for i, node := range nodes {
			if instruction == 'L' {
				nodes[i] = net.net[node][0]
			} else {
				nodes[i] = net.net[node][1]
			}
		}
		step++

		i := 0
		for _, n := range nodes {
			// remove nodes which qualify as end nodes
			if !isEndFunc(n) {
				nodes[i] = n
				i++
			} else {
				// save current step for each end node
				steps = append(steps, step)
			}
		}
		nodes = nodes[:i]
	}

	result := slicesx.Reduce(steps[1:], func(acc int, val int) int {
		return mathx.LCM(acc, val)
	}, steps[0])
	return result
}

func startNodeQualifierRegular(name string) bool {
	return name == "AAA"
}
func endNodeQualifierRegular(name string) bool {
	return name == "ZZZ"
}
func startNodeQualifierGhost(name string) bool {
	return strings.HasSuffix(name, "A")
}
func endNodeQualifierGhost(name string) bool {
	return strings.HasSuffix(name, "Z")
}

func calculateSteps(lines []string, ghost bool) int {
	var isStartFunc, isEndFunc nodeQualifier
	if ghost {
		isStartFunc = startNodeQualifierGhost
		isEndFunc = endNodeQualifierGhost
	} else {
		isStartFunc = startNodeQualifierRegular
		isEndFunc = endNodeQualifierRegular
	}

	instructions := &instructions{S: lines[0]}
	net := NewNetwork(lines[2:], isStartFunc)

	return net.Walk(instructions, isEndFunc)
}
