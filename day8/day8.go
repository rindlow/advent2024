package day8

import (
	"strconv"

	"github.com/rindlow/aoc-utils"
)

type Coord struct {
	X, Y int
}

type Resonator = func(a, b, max Coord) []Coord

func readMap(filename string) (antennas map[rune][]Coord, max Coord) {
	antennas = make(map[rune][]Coord)
	input := utils.ReadLines(filename)
	max = Coord{len(input[0]) - 1, len(input) - 1}
	for y, line := range input {
		for x, char := range line {
			if char != '.' {
				antennas[char] = append(antennas[char], Coord{x, y})
			}
		}
	}
	return
}

func antinodes(a, b, max Coord) []Coord {
	dx := a.X - b.X
	dy := a.Y - b.Y
	return []Coord{{a.X + dx, a.Y + dy}, {b.X - dx, b.Y - dy}}
}

func harmonics(a, b, max Coord) (nodes []Coord) {
	nodes = []Coord{}
	dx := a.X - b.X
	dy := a.Y - b.Y
	n := utils.Max(max.X/utils.Abs(dx), max.Y/utils.Abs(dy))
	for i := -n; i < 2*n; i++ {
		nodes = append(nodes, Coord{a.X + i*dx, a.Y + i*dy})
	}
	return
}

func uniqueLocations(filename string, resonator Resonator) (count int) {
	antennas, max := readMap(filename)
	locations := make(map[Coord]bool)
	for _, coords := range antennas {
		for i := 0; i < len(coords)-1; i++ {
			for j := i + 1; j < len(coords); j++ {
				for _, node := range resonator(coords[i], coords[j], max) {
					locations[node] = true
				}
			}
		}
	}
	for coord := range locations {
		if coord.X >= 0 && coord.X <= max.X && coord.Y >= 0 && coord.Y <= max.Y {
			count += 1
		}
	}
	return
}

func Part1(filename string) string {
	return strconv.Itoa(uniqueLocations(filename, antinodes))
}

func Part2(filename string) string {
	return strconv.Itoa(uniqueLocations(filename, harmonics))
}
