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
