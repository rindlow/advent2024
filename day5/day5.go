package day5

import (
	"slices"
	"strconv"

	"github.com/rindlow/aoc-utils"
)

func parseInput(filename string) (func(int, int) int, [][]int) {
	rules := make(map[int][]int)
	updates := [][]int{}
	doRules := true
	for _, line := range utils.ReadLines(filename) {
		if line == "" {
			doRules = false
			continue
		}
		if doRules {
			pages := utils.StringToNumbersWithDelimiter(line, "|")
			rules[pages[0]] = append(rules[pages[0]], pages[1])
		} else {
			updates = append(updates, utils.StringToNumbersWithDelimiter(line, ","))
		}
	}
	cmp := func(a, b int) int {
		if slices.Contains(rules[a], b) {
			return -1
		}
		if slices.Contains(rules[b], a) {
			return 1
		}
		return 0
	}
	return cmp, updates
}

func sumMiddlePage(filename string) (sum int) {
	cmp, updates := parseInput(filename)
	for _, update := range updates {
		if slices.IsSortedFunc(update, cmp) {
			sum += update[len(update)/2]
		}
	}
	return
}

func sumMiddlePageCorrected(filename string) (sum int) {
	cmp, updates := parseInput(filename)
	for _, update := range updates {
		if !slices.IsSortedFunc(update, cmp) {
			slices.SortFunc(update, cmp)
			sum += update[len(update)/2]
		}
	}
	return
}

func Part1(filename string) string {
	return strconv.Itoa(sumMiddlePage(filename))
}

func Part2(filename string) string {
	return strconv.Itoa(sumMiddlePageCorrected(filename))
}
