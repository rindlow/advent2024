package day17

import (
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/rindlow/aoc-utils"
)

type State struct {
	A, B, C, IP int
}

func atoi(a string) int {
	i, err := strconv.Atoi(a)
	if err != nil {
		log.Fatalf("Atoi %s: %q", a, err)
	}
	return i
}

func readInput(filename string) (state State, program []int) {
	for _, line := range utils.ReadLines(filename) {
		switch {
		case strings.HasPrefix(line, "Register A: "):
			state.A = atoi(line[12:])
		case strings.HasPrefix(line, "Register B: "):
			state.B = atoi(line[12:])
		case strings.HasPrefix(line, "Register C: "):
			state.C = atoi(line[12:])
		case strings.HasPrefix(line, "Program: "):
			program = utils.StringToNumbersWithDelimiter(line[9:], ",")
		}
	}
	return
}

func combo(operand int, state *State) int {
	switch operand {
	case 0, 1, 2, 3:
		return operand
	case 4:
		return state.A
	case 5:
		return state.B
	case 6:
		return state.C
	}
	log.Fatalf("Illegal combo operand")
	return -1
}

func execute(operator, operand int, state *State) (int, bool) {
	state.IP += 2
	switch operator {
	case 0: // adv
		state.A /= 1 << combo(operand, state)
	case 1: // bxl
		state.B ^= operand
	case 2: // bst
		state.B = combo(operand, state) & 7
	case 3: // jnz
		if state.A != 0 {
			state.IP = operand
		}
	case 4: // bxc
		state.B ^= state.C
	case 5: // out
		return combo(operand, state) & 7, true
	case 6: // bdv
		state.B = state.A / (1 << combo(operand, state))
	case 7: // bdv
		state.C = state.A / (1 << combo(operand, state))
	}
	return -1, false
}

func runProgram(state State, program []int) string {
	output := []string{}
	for state.IP < len(program)-1 {
		out, isOut := execute(program[state.IP], program[state.IP+1], &state)
		if isOut {
			output = append(output, strconv.Itoa(out))
		}
	}
	return strings.Join(output, ",")
}

func findSelfOutput(state State, program []int) int {
	a := 0
	rev := slices.Clone(program)
	slices.Reverse(rev)
	for _, digit := range rev {
		a <<= 3
	loop:
		for j := 0; j < 8; j++ {
			stateCopy := state
			stateCopy.A = a | j
		run:
			for stateCopy.IP < len(program)-1 {
				out, isOut := execute(program[stateCopy.IP], program[stateCopy.IP+1], &stateCopy)
				if isOut {
					if out == digit {
						a |= j
						break loop
					} else {
						break run
					}
				}
			}
		}
	}
	return a
}

func Part1(filename string) string {
	state, program := readInput(filename)
	return runProgram(state, program)
}

func Part2(filename string) string {
	state, program := readInput(filename)
	return strconv.Itoa(findSelfOutput(state, program))
}
