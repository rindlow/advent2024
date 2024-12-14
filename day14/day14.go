package day14

import (
	"fmt"
	"log"
	"regexp"
	"slices"
	"strconv"

	"github.com/rindlow/aoc-utils"
)

type Coord struct {
	X, Y int
}

type Robot struct {
	Pos, Vel Coord
}

func Atoi(a string) int {
	i, err := strconv.Atoi(a)
	if err != nil {
		log.Fatalf("Atoi %s: %q", a, err)
	}
	return i
}

func readRobots(filename string) (robots []Robot) {
	re := regexp.MustCompile(`p=(\d+),(\d+) v=(-?\d+),(-?\d+)`)
	for _, line := range utils.ReadLines(filename) {
		match := re.FindStringSubmatch(line)
		robots = append(robots, Robot{Coord{Atoi(match[1]), Atoi(match[2])},
			Coord{Atoi(match[3]), Atoi(match[4])}})
	}
	return
}

func CmpX(a, b Coord) int {
	switch {
	case a.X < b.X:
		return -1
	case a.X > b.X:
		return 1
	default:
		return 0
	}
}

func safetyFactor(filename string, lenX, lenY int) int {
	robots := readRobots(filename)
	quads := make([]int, 4)
	halfX := lenX / 2
	halfY := lenY / 2
	for _, robot := range robots {
		x := (robot.Pos.X + 100*robot.Vel.X) % lenX
		if x < 0 {
			x += lenX
		}
		y := (robot.Pos.Y + 100*robot.Vel.Y) % lenY
		if y < 0 {
			y += lenY
		}

		switch {
		case x < halfX && y < halfY:
			quads[0] += 1
		case x > halfX && y < halfY:
			quads[1] += 1
		case x < halfX && y > halfY:
			quads[2] += 1
		case x > halfX && y > halfY:
			quads[3] += 1
		}
	}
	return quads[0] * quads[1] * quads[2] * quads[3]
}

func printGrid(pos map[Coord]bool, lenX, lenY int) {
	for y := 0; y < lenY; y++ {
		for x := 0; x < lenX; x++ {
			if pos[Coord{x, y}] {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func easterEgg(filename string, lenX, lenY int) {
	robots := readRobots(filename)
	t := 1
	for {
		grid := make(map[Coord]bool)
		for i, robot := range robots {
			x := ((robot.Pos.X+robot.Vel.X)%lenX + lenX) % lenX
			y := ((robot.Pos.Y+robot.Vel.Y)%lenY + lenY) % lenY
			pos := Coord{x, y}
			grid[pos] = true
			robots[i].Pos = pos
		}
		if (t-95)%101 == 0 {
			fmt.Printf("t = %d\n", t)
			printGrid(grid, lenX, lenY)
		}
		t += 1
	}
}

func findEasterEgg(filename string, lenX, lenY int) int {
	robots := readRobots(filename)
	halfY := lenY / 2
	t := 1
	for {
		row := []Coord{}
		for i, robot := range robots {
			robots[i].Pos.X = ((robot.Pos.X+robot.Vel.X)%lenX + lenX) % lenX
			robots[i].Pos.Y = ((robot.Pos.Y+robot.Vel.Y)%lenY + lenY) % lenY
			if robots[i].Pos.Y == halfY {
				row = append(row, robots[i].Pos)
			}
		}
		slices.SortFunc(row, CmpX)
		run := 0
		last := -2
		for _, pos := range row {
			if pos.X == last+1 {
				run += 1
				if run > 7 {
					return t
				}
			} else {
				run = 0
			}
			last = pos.X
		}
		t += 1
	}
}

func Part1(filename string) string {
	return strconv.Itoa(safetyFactor(filename, 101, 103))
}

func Part2(filename string) string {
	return strconv.Itoa(findEasterEgg(filename, 101, 103))
}
