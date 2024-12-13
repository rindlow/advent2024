package day13

import (
	"log"
	"regexp"
	"strconv"

	"github.com/rindlow/aoc-utils"
)

type Coord struct {
	X, Y int
}
type Machine struct {
	A, B, Prize Coord
}

func CoordFromStrings(xStr, yStr string) Coord {
	x, err := strconv.Atoi(xStr)
	if err != nil {
		log.Fatalf("Atoi %s: %q", xStr, err)
	}
	y, err := strconv.Atoi(yStr)
	if err != nil {
		log.Fatalf("Atoi %s: %q", yStr, err)
	}
	return Coord{x, y}
}

func readMachines(filename string) (machines []Machine) {
	re := regexp.MustCompile(`Button ([AB]): X\+(\d+), Y\+(\d+)|Prize: X=(\d+), Y=(\d+)`)
	machine := Machine{}
	for _, line := range utils.ReadLines(filename) {
		match := re.FindStringSubmatch(line)
		switch {
		case len(match) == 0:
			continue
		case match[1] == "A":
			machine.A = CoordFromStrings(match[2], match[3])
		case match[1] == "B":
			machine.B = CoordFromStrings(match[2], match[3])
		case match[4] != "":
			machine.Prize = CoordFromStrings(match[4], match[5])
			machines = append(machines, machine)
			machine = Machine{}
		}
	}
	return
}

func solution(machine Machine) int {
	dividendA := machine.B.Y*machine.Prize.X - machine.B.X*machine.Prize.Y
	divisorA := machine.A.X*machine.B.Y - machine.A.Y*machine.B.X
	if dividendA%divisorA != 0 {
		return 0
	}
	A := dividendA / divisorA
	dividendB := machine.Prize.X - machine.A.X*A
	divisorB := machine.B.X
	if dividendB%divisorB != 0 {
		return 0
	}
	B := dividendB / divisorB
	return 3*A + B
}

func patch(machine *Machine) *Machine {
	machine.Prize.X += 10000000000000
	machine.Prize.Y += 10000000000000
	return machine
}

func fewestTokens(filename string) (sum int) {
	for _, machine := range readMachines(filename) {
		sum += solution(machine)
	}
	return
}

func fewestTokensPatched(filename string) (sum int) {
	for _, machine := range readMachines(filename) {
		sum += solution(*patch(&machine))
	}
	return
}

func Part1(filename string) string {
	return strconv.Itoa(fewestTokens(filename))
}

func Part2(filename string) string {
	return strconv.Itoa(fewestTokensPatched(filename))
}
