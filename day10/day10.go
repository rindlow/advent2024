package day10

import (
	"strconv"

	"github.com/rindlow/aoc-utils"
)

type Coord struct {
	X, Y int
}

type Map = []string

func height(m Map, pos Coord) int {
	if pos.X < 0 || pos.X >= len(m[0]) || pos.Y < 0 || pos.Y >= len(m) {
		return -1
	}
	return int(m[pos.Y][pos.X] - '0')
}

func findPaths(m Map, pos Coord, nextHeight int) (paths []Coord) {
	if nextHeight > 9 {
		return []Coord{pos}
	}
	for _, n := range []Coord{{pos.X - 1, pos.Y}, {pos.X + 1, pos.Y}, {pos.X, pos.Y - 1}, {pos.X, pos.Y + 1}} {
		if height(m, n) == nextHeight {
			paths = append(paths, findPaths(m, n, nextHeight+1)...)
		}
	}
	return
}

func sumScores(filename string) (sum int) {
	trailMap := utils.ReadLines(filename)
	for y, line := range trailMap {
		for x, h := range line {
			if h == '0' {
				trailEnds := make(map[Coord]bool)
				for _, end := range findPaths(trailMap, Coord{x, y}, 1) {
					trailEnds[end] = true
				}
				sum += len(trailEnds)
			}
		}
	}
	return
}

func sumRatings(filename string) (sum int) {
	trailMap := utils.ReadLines(filename)
	for y, line := range trailMap {
		for x, h := range line {
			if h == '0' {
				sum += len(findPaths(trailMap, Coord{x, y}, 1))
			}
		}
	}
	return
}
func Part1(filename string) string {
	return strconv.Itoa(sumScores(filename))
}

func Part2(filename string) string {
	return strconv.Itoa(sumRatings(filename))
}
