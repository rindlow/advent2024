package main

import (
	"slices"
	"strconv"

	"github.com/rindlow/aoc-utils"
)

func readLists(filename string) (lists [2][]int) {
	lines := utils.ReadLines(filename)
	for _, line := range lines {
		nums := utils.StringToNumbers(line)
		for i, num := range nums {
			lists[i] = append(lists[i], num)
		}
	}
	return
}

func sumDiffs(filename string) (sum int) {
	lists := readLists(filename)
	for _, list := range lists {
		slices.Sort(list)
	}
	for i, a := range lists[0] {
		sum += utils.Abs(a - lists[1][i])
	}
	return
}

func similarityScore(filename string) (sum int) {
	count := make(map[int]int)
	lists := readLists(filename)
	for _, number := range lists[1] {
		count[number] = count[number] + 1
	}
	for _, a := range lists[0] {
		sum += a * count[a]
	}
	return
}

func part1(filename string) string {
	return strconv.Itoa(sumDiffs(filename))
}

func part2(filename string) string {
	return strconv.Itoa(similarityScore(filename))
}
