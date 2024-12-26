package day24

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/rindlow/aoc-utils"
)

type Gate struct {
	A, Operation, B, Output string
}

type State struct {
	Values map[string]bool
	Gates  map[string]Gate
}

func readState(filename string) (state State) {
	state.Values = make(map[string]bool)
	state.Gates = make(map[string]Gate)
	readState := true
	for _, line := range utils.ReadLines(filename) {
		switch {
		case line == "":
			readState = false
		case readState:
			state.Values[line[:3]] = line[5] == '1'
		default:
			space := strings.Split(line, " ")
			state.Gates[space[4]] = Gate{space[0], space[1], space[2], space[4]}
		}
	}
	return
}

func updateState(output string, state *State) bool {
	gate := state.Gates[output]
	a, ok := state.Values[gate.A]
	if !ok {
		a = updateState(gate.A, state)
	}
	b, ok := state.Values[gate.B]
	if !ok {
		b = updateState(gate.B, state)
	}
	switch gate.Operation {
	case "AND":
		state.Values[output] = a && b
	case "OR":
		state.Values[output] = a || b
	case "XOR":
		state.Values[output] = a != b
	}
	return state.Values[output]
}

func booleanGates(state *State) (z int) {
	zGates := []string{}
	for output := range state.Gates {
		if output[0] == 'z' {
			zGates = append(zGates, output)
		}
	}
	slices.Sort(zGates)
	slices.Reverse(zGates)
	for _, output := range zGates {
		updateState(output, state)
		z <<= 1
		if state.Values[output] {
			z |= 1
		}
	}
	return
}

func gatesInfront(root string, state *State, depth int) (gates []Gate) {
	if depth == 0 {
		return
	}
	for _, gate := range state.Gates {
		if gate.A == root || gate.B == root {
			gates = append(gates, gate)
			gates = append(gates, gatesInfront(gate.Output, state, depth-1)...)
		}
	}
	return
}

func checkAdder(pin int, carryIn string, state *State) (carryOut string, swapped []string) {
	var xor1, xor2, and1, and2, or Gate
	x := fmt.Sprintf("x%02d", pin)
	for _, gate := range gatesInfront(x, state, 2) {
		if gate.A[0] == 'x' || gate.B[0] == 'x' {
			switch gate.Operation {
			case "XOR":
				xor1 = gate
			case "AND":
				and1 = gate
			}
		}
		if gate.A == carryIn || gate.B == carryIn {
			switch gate.Operation {
			case "XOR":
				xor2 = gate
			case "AND":
				and2 = gate
			}
		}
		if gate.Operation == "OR" {
			or = gate
			carryOut = gate.Output
		}
	}
	if pin == 0 {
		carryOut = and1.Output
		return
	}
	if or.A == "" {
		or = gatesInfront(and2.Output, state, 1)[0]
	}
	if xor2.Output[0] != 'z' {
		swapped = append(swapped, xor2.Output)
		switch {
		case and1.Output[0] == 'z':
			swapped = append(swapped, and1.Output)
			and1.Output, xor2.Output = xor2.Output, and1.Output
		case and2.Output[0] == 'z':
			swapped = append(swapped, and2.Output)
			and2.Output, xor2.Output = xor2.Output, and2.Output
		case xor1.Output[0] == 'z':
			swapped = append(swapped, xor1.Output)
			xor1.Output, xor2.Output = xor2.Output, xor1.Output
		case or.Output[0] == 'z':
			swapped = append(swapped, or.Output)
			or.Output, xor2.Output = xor2.Output, or.Output
		}
		carryOut = or.Output
	}
	if xor2.A != xor1.Output && xor2.B != xor1.Output {
		input := xor2.A
		if xor2.A == carryIn {
			input = xor2.B
		}
		swapped = append(swapped, xor1.Output)
		switch {
		case and1.Output == input:
			swapped = append(swapped, and1.Output)
			and1.Output, xor1.Output = xor1.Output, and1.Output
		case and2.Output == input:
			swapped = append(swapped, and2.Output)
			and2.Output, xor1.Output = xor1.Output, and2.Output
		case xor2.Output == input:
			swapped = append(swapped, xor2.Output)
			xor1.Output, xor2.Output = xor2.Output, xor1.Output
		case or.Output == input:
			swapped = append(swapped, or.Output)
			or.Output, xor1.Output = xor1.Output, or.Output
		}
		carryOut = or.Output
	}
	return
}

func swappedWires(state *State) (swapped []string) {
	carryIn := ""
	var swap []string
	for pin := 0; pin < 45; pin++ {
		carryIn, swap = checkAdder(pin, carryIn, state)
		if len(swap) == 2 {
			state.Gates[swap[0]], state.Gates[swap[1]] = state.Gates[swap[1]], state.Gates[swap[0]]
		}
		swapped = append(swapped, swap...)
	}
	return
}

func Part1(filename string) string {
	state := readState(filename)
	return strconv.Itoa(booleanGates(&state))
}

func Part2(filename string) string {
	state := readState(filename)
	swapped := swappedWires(&state)
	slices.Sort(swapped)
	return strings.Join(swapped, ",")
}
