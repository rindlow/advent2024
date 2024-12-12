package day12

import (
	"maps"
	"slices"
	"strconv"

	"github.com/rindlow/aoc-utils"
)

type Coord struct {
	X, Y int
}

type Plot struct {
	Coords map[Coord]bool
	Plant  rune
}

type LengthCalc = func(Plot, *[]string) int

func NewPlot(plant rune) Plot {
	return Plot{make(map[Coord]bool), plant}
}

func (plot *Plot) Contains(pos Coord) bool {
	_, ok := plot.Coords[pos]
	return ok
}

func (plot *Plot) Add(pos Coord) {
	plot.Coords[pos] = true
}

func neighbours(pos Coord) (nbrs []Coord) {
	return []Coord{{pos.X - 1, pos.Y}, {pos.X + 1, pos.Y}, {pos.X, pos.Y - 1}, {pos.X, pos.Y + 1}}
}

func getPlant(pos Coord, plotMap *[]string) rune {
	if pos.X < 0 || pos.X >= len((*plotMap)[0]) || pos.Y < 0 || pos.Y >= len(*plotMap) {
		return '.'
	}
	return rune((*plotMap)[pos.Y][pos.X])
}

func findPlot(start Coord, plotMap *[]string) (plot Plot) {
	plant := getPlant(start, plotMap)
	plot = NewPlot(plant)
	queue := []Coord{start}
	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]
		plot.Add(pos)
		for _, neigbour := range neighbours(pos) {
			if !plot.Contains(neigbour) && getPlant(neigbour, plotMap) == plant && slices.Index(queue, neigbour) < 0 {
				queue = append(queue, neigbour)
			}
		}
	}
	return
}

func perimeter(plot Plot, plotMap *[]string) (p int) {
	plant := ' '
	for pos := range plot.Coords {
		if plant == ' ' {
			plant = getPlant(pos, plotMap)
		}
		for _, neighbour := range neighbours(pos) {
			if getPlant(neighbour, plotMap) != plant {
				p += 1
			}
		}
	}
	return
}

func findPlots(plotMap *[]string) (plots map[rune][]Plot) {
	plots = make(map[rune][]Plot)
	for y, line := range *plotMap {
		for x, plant := range line {
			pos := Coord{x, y}
			plotFound := false
			for _, plot := range plots[plant] {
				if plot.Contains(pos) {
					plotFound = true
					break
				}
			}
			if !plotFound {
				plots[plant] = append(plots[plant], findPlot(pos, plotMap))
			}
		}
	}
	return
}

func rowCmp(a, b Coord) int {
	switch {
	case a.Y < b.Y:
		return -1
	case a.Y > b.Y:
		return 1
	default:
		if a.X < b.X {
			return -1
		}
		return 1
	}
}

func colCmp(a, b Coord) int {
	switch {
	case a.X < b.X:
		return -1
	case a.X > b.X:
		return 1
	default:
		if a.Y < b.Y {
			return -1
		}
		return 1
	}
}

func byRow(plot Plot, plotMap *[]string) (nSides int) {
	lastY := -1
	lastX := -2
	lastTopEdge := false
	lastBottomEdge := false
	for _, pos := range slices.SortedFunc(maps.Keys(plot.Coords), rowCmp) {
		if pos.Y != lastY {
			lastX = -2
			lastTopEdge = false
			lastBottomEdge = false
		}
		if pos.X > lastX+1 {
			lastTopEdge = false
			lastBottomEdge = false
		}
		if getPlant(Coord{pos.X, pos.Y - 1}, plotMap) != plot.Plant {
			if !lastTopEdge {
				nSides += 1
			}
			lastTopEdge = true
		} else {
			lastTopEdge = false
		}
		if getPlant(Coord{pos.X, pos.Y + 1}, plotMap) != plot.Plant {
			if !lastBottomEdge {
				nSides += 1
			}
			lastBottomEdge = true
		} else {
			lastBottomEdge = false
		}
		lastX = pos.X
		lastY = pos.Y
	}
	return
}

func byCol(plot Plot, plotMap *[]string) (nSides int) {
	lastY := -1
	lastX := -2
	lastLeftEdge := false
	lastRightEdge := false
	for _, pos := range slices.SortedFunc(maps.Keys(plot.Coords), colCmp) {
		if pos.X != lastX {
			lastY = -2
			lastLeftEdge = false
			lastRightEdge = false
		}
		if pos.Y > lastY+1 {
			lastLeftEdge = false
			lastRightEdge = false
		}
		if getPlant(Coord{pos.X - 1, pos.Y}, plotMap) != plot.Plant {
			if !lastLeftEdge {
				nSides += 1
			}
			lastLeftEdge = true
		} else {
			lastLeftEdge = false
		}
		if getPlant(Coord{pos.X + 1, pos.Y}, plotMap) != plot.Plant {
			if !lastRightEdge {
				nSides += 1
			}
			lastRightEdge = true
		} else {
			lastRightEdge = false
		}
		lastX = pos.X
		lastY = pos.Y
	}
	return
}

func sides(plot Plot, plotMap *[]string) (p int) {
	return byRow(plot, plotMap) + byCol(plot, plotMap)
}

func fencePrice(filename string, calc LengthCalc) (sum int) {
	plotMap := utils.ReadLines(filename)
	for _, ps := range findPlots(&plotMap) {
		for _, plot := range ps {
			area := len(plot.Coords)
			peri := calc(plot, &plotMap)
			sum += area * peri
		}
	}
	return
}

func Part1(filename string) string {
	return strconv.Itoa(fencePrice(filename, perimeter))
}

func Part2(filename string) string {
	return strconv.Itoa(fencePrice(filename, sides))
}
