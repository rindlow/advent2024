package day6

import (
	"maps"
	"slices"
	"strconv"

	"github.com/rindlow/aoc-utils"
)

type Pos struct {
	X, Y int
}

func (pos Pos) add(dir Pos) Pos {
	return Pos{pos.X + dir.X, pos.Y + dir.Y}
}

func turn(dir Pos) Pos {
	switch {
	case dir.X == 0 && dir.Y == -1:
		return Pos{1, 0}
	case dir.X == 1 && dir.Y == 0:
		return Pos{0, 1}
	case dir.X == 0 && dir.Y == 1:
		return Pos{-1, 0}
	default:
		return Pos{0, -1}
	}
}

func readMap(filename string) (start Pos, max Pos, obstacles map[Pos]bool) {
	obstacles = make(map[Pos]bool)
	guardMap := utils.ReadLines(filename)
	max = Pos{len(guardMap[0]) - 1, len(guardMap) - 1}
	for row, line := range guardMap {
		for col, char := range line {
			switch char {
			case '^':
				start = Pos{col, row}
			case '#':
				obstacles[Pos{col, row}] = true
			}
		}
	}
	return
}

func walk(pos Pos, dir Pos, max Pos, obstacles map[Pos]bool) ([]Pos, bool) {
	visited := make(map[Pos][]Pos)
	for {
		if pos.X < 0 || pos.X > max.X || pos.Y < 0 || pos.Y > max.Y {
			break
		}
		if slices.Contains(visited[pos], dir) {
			return []Pos{}, true
		}
		visited[pos] = append(visited[pos], dir)
		for obstacles[pos.add(dir)] {
			dir = turn(dir)
		}
		pos = pos.add(dir)
	}
	return slices.Collect(maps.Keys(visited)), false
}

func visitedPositions(filename string) int {
	start, max, obstacles := readMap(filename)
	visited, _ := walk(start, Pos{0, -1}, max, obstacles)
	return len(visited)
}

func numberOfLoops(filename string) (count int) {
	start, max, obstacles := readMap(filename)
	dir := Pos{0, -1}
	visited, _ := walk(start, dir, max, obstacles)
	for _, pos := range visited {
		newObstacles := maps.Clone(obstacles)
		newObstacles[pos] = true
		_, loop := walk(start, dir, max, newObstacles)
		if loop {
			count += 1
		}
	}
	return
}

func Part1(filename string) string {
	return strconv.Itoa(visitedPositions(filename))
}

func Part2(filename string) string {
	return strconv.Itoa(numberOfLoops(filename))
}
