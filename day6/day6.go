package day6

import (
	"maps"
	"strconv"

	"github.com/rindlow/aoc-utils"
)

type Coord struct {
	X, Y int
}

type State struct {
	Pos    Coord
	Dx, Dy int
}

func (s State) next() State {
	return State{Coord{s.Pos.X + s.Dx, s.Pos.Y + s.Dy}, s.Dx, s.Dy}
}

func (s State) turn() State {
	switch {
	case s.Dx == 0 && s.Dy == -1:
		return State{s.Pos, 1, 0}
	case s.Dx == 1 && s.Dy == 0:
		return State{s.Pos, 0, 1}
	case s.Dx == 0 && s.Dy == 1:
		return State{s.Pos, -1, 0}
	default:
		return State{s.Pos, 0, -1}
	}
}

func readMap(filename string) (start State, max Coord, obstacles map[Coord]bool) {
	obstacles = make(map[Coord]bool)
	guardMap := utils.ReadLines(filename)
	max = Coord{len(guardMap[0]) - 1, len(guardMap) - 1}
	for row, line := range guardMap {
		for col, char := range line {
			pos := Coord{col, row}
			switch char {
			case '^':
				start = State{pos, 0, -1}
			case '#':
				obstacles[pos] = true
			}
		}
	}
	return
}

func walk(state State, max Coord, obstacles map[Coord]bool) ([]State, bool) {
	states := []State{}
	visited := make(map[State]bool)
	for {
		if state.Pos.X < 0 || state.Pos.X > max.X || state.Pos.Y < 0 || state.Pos.Y > max.Y {
			break
		}
		if visited[state] {
			return []State{}, true
		}
		states = append(states, state)
		visited[state] = true
		for obstacles[state.next().Pos] {
			state = state.turn()
		}
		state = state.next()
	}
	return states, false
}

func visitedPositions(filename string) int {
	start, max, obstacles := readMap(filename)
	states, _ := walk(start, max, obstacles)
	visited := make(map[Coord]bool)
	for _, state := range states {
		visited[state.Pos] = true
	}
	return len(visited)
}

func numberOfLoops(filename string) int {
	start, max, obstacles := readMap(filename)
	states, _ := walk(start, max, obstacles)
	loops := make(map[Coord]bool)
	checked := make(map[Coord]bool)
	for step, state := range states {
		if step == 0 || checked[state.Pos] {
			continue
		}
		newObstacles := maps.Clone(obstacles)
		newObstacles[state.Pos] = true
		_, loop := walk(states[step-1], max, newObstacles)
		if loop {
			loops[state.Pos] = true
		}
		checked[state.Pos] = true
	}
	return len(loops)
}

func Part1(filename string) string {
	return strconv.Itoa(visitedPositions(filename))
}

func Part2(filename string) string {
	return strconv.Itoa(numberOfLoops(filename))
}
