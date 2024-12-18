package day18

import (
	"fmt"
	"maps"
	"strconv"

	"github.com/rindlow/aoc-utils"
)

type Coord struct {
	X, Y int
}

func readIncoming(filename string) (incoming []Coord) {
	for _, line := range utils.ReadLines(filename) {
		nums := utils.StringToNumbersWithDelimiter(line, ",")
		incoming = append(incoming, Coord{nums[0], nums[1]})
	}
	return
}

func equals(a, b Coord) bool {
	return a.X == b.X && a.Y == b.Y
}

func any[T comparable](set map[T]bool) bool {
	for v := range maps.Values(set) {
		if v {
			return true
		}
	}
	return false
}

func neighbours(pos Coord, space map[Coord]bool, max int) (neigbours []Coord) {
	for _, n := range []Coord{{pos.X - 1, pos.Y}, {pos.X + 1, pos.Y},
		{pos.X, pos.Y - 1}, {pos.X, pos.Y + 1}} {
		_, wall := space[n]
		if !wall && n.X >= 0 && n.X <= max && n.Y >= 0 && n.Y <= max {
			neigbours = append(neigbours, n)
		}
	}
	return
}

func h(pos, goal Coord) int {
	return goal.X - pos.X + goal.Y - pos.Y
}

func shortestPath(space map[Coord]bool, max int) int {
	start := Coord{0, 0}
	goal := Coord{max, max}

	openSet := make(map[Coord]bool)
	openSet[start] = true
	cameFrom := make(map[Coord]Coord)
	gScore := make(map[Coord]int)
	gScore[start] = 0
	fScore := make(map[Coord]int)
	fScore[start] = h(start, goal)

	for any(openSet) {
		// fmt.Printf("openSet = %v\n", openSet)
		minF := 100000
		var current Coord
		for Coord, exists := range openSet {
			if exists && fScore[Coord] < minF {
				current = Coord
				minF = fScore[Coord]
			}
		}
		// fmt.Printf(" current = %v, minF = %d\n", current, minF)
		if equals(current, goal) {
			return gScore[current]
		}
		openSet[current] = false
		for _, n := range neighbours(current, space, max) {
			tentative := gScore[current] + 1
			neighbourG, ok := gScore[n]
			// fmt.Printf("ok = %v, tentative = %d, neighbourG = %d\n", ok, tentative, neighbourG)
			if !ok || tentative < neighbourG {
				cameFrom[n] = current
				gScore[n] = tentative
				fScore[n] = tentative + h(n, goal)
			}
			_, seen := openSet[n]
			if !seen {
				openSet[n] = true
			}
		}
	}
	return -1
}

func runShortestPath(filename string, nBytes, max int) int {
	incoming := readIncoming(filename)
	space := make(map[Coord]bool)
	for i := 0; i < nBytes; i++ {
		space[incoming[i]] = true
	}
	return shortestPath(space, max)
}

func firstBlocking(filename string, nBytes, max int) string {
	incoming := readIncoming(filename)
	space := make(map[Coord]bool)
	i := 0
	for {
		space[incoming[i]] = true
		if i >= nBytes {
			pathLen := shortestPath(space, max)
			if pathLen < 0 {
				return fmt.Sprintf("%d,%d", incoming[i].X, incoming[i].Y)
			}
		}
		i += 1
	}
}

func Part1(filename string) string {
	return strconv.Itoa(runShortestPath(filename, 1024, 70))
}

func Part2(filename string) string {
	return firstBlocking(filename, 1024, 70)
}
