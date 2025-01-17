package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"time"

	"rindlow.se/advent2024/day1"
	"rindlow.se/advent2024/day2"
    "rindlow.se/advent2024/day3"
    "rindlow.se/advent2024/day4"
    "rindlow.se/advent2024/day5"
    "rindlow.se/advent2024/day6"
    "rindlow.se/advent2024/day7"
    "rindlow.se/advent2024/day8"
    "rindlow.se/advent2024/day9"
    "rindlow.se/advent2024/day10"
    "rindlow.se/advent2024/day11"
    "rindlow.se/advent2024/day12"
    "rindlow.se/advent2024/day13"
    "rindlow.se/advent2024/day14"
    "rindlow.se/advent2024/day15"
    "rindlow.se/advent2024/day16"
    "rindlow.se/advent2024/day17"
    "rindlow.se/advent2024/day18"
    "rindlow.se/advent2024/day19"
    "rindlow.se/advent2024/day20"
    "rindlow.se/advent2024/day21"
    "rindlow.se/advent2024/day22"
    "rindlow.se/advent2024/day23"
    "rindlow.se/advent2024/day24"
    "rindlow.se/advent2024/day25"
	// Insert more modules above this line
)

func runDay(day int, part1 func(string) string, part2 func(string) string) time.Duration {
	filename := fmt.Sprintf("input/day%d.txt", day)

	fmt.Printf("Day %d:\n", day)

	// Run part1
	start1 := time.Now()
	fmt.Printf("  part 1: %s\n", part1(filename))
	elapsed1 := time.Since(start1)
	fmt.Printf("    %s\n", elapsed1)

	// Run part2
	start2 := time.Now()
	fmt.Printf("  part 2: %s\n", part2(filename))
	elapsed2 := time.Since(start2)
	fmt.Printf("    %s\n\n", elapsed2)
	return elapsed1 + elapsed2
}

func main() {
	var elapsed time.Duration
	days := map[int][]func(string) string{
		1: {day1.Part1, day1.Part2},
		2: {day2.Part1, day2.Part2},
		3: {day3.Part1, day3.Part2},
		4: {day4.Part1, day4.Part2},
		5: {day5.Part1, day5.Part2},
		6: {day6.Part1, day6.Part2},
		7: {day7.Part1, day7.Part2},
		8: {day8.Part1, day8.Part2},
		9: {day9.Part1, day9.Part2},
		10: {day10.Part1, day10.Part2},
		11: {day11.Part1, day11.Part2},
		12: {day12.Part1, day12.Part2},
		13: {day13.Part1, day13.Part2},
		14: {day14.Part1, day14.Part2},
		15: {day15.Part1, day15.Part2},
		16: {day16.Part1, day16.Part2},
		17: {day17.Part1, day17.Part2},
		18: {day18.Part1, day18.Part2},
		19: {day19.Part1, day19.Part2},
		20: {day20.Part1, day20.Part2},
		21: {day21.Part1, day21.Part2},
		22: {day22.Part1, day22.Part2},
		23: {day23.Part1, day23.Part2},
		24: {day24.Part1, day24.Part2},
		25: {day25.Part1, day25.Part2},
		// Insert more days above this line
	}
	dayNos := make([]int, len(days))
	if len(os.Args) > 1 {
		day, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatalf("%q", err)
		}
		parts, ok := days[day]
		if !ok {
			log.Fatalf("Day %d not found", day)
		}
		elapsed += runDay(day, parts[0], parts[1])
	} else {
		i := 0
		for day := range days {
			dayNos[i] = day
			i += 1
		}
		slices.Sort(dayNos)
		for _, day := range dayNos {
			elapsed += runDay(day, days[day][0], days[day][1])
		}
	}
	fmt.Printf("\nTotal time: %s\n\n", elapsed)
}
