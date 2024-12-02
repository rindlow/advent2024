package main

import (
	"strconv"

	"github.com/rindlow/aoc-utils"
)

func isSafe(levels []int) bool {
	sign := utils.Sign(levels[1] - levels[0])
	for i, level := range levels[1:] {
		diff := level - levels[i]
		absdiff := utils.Abs(diff)
		if utils.Sign(diff) != sign || absdiff < 1 || absdiff > 3 {
			return false
		}
	}
	return true
}

func removeOne(levels []int) (lists [][]int) {
	for i := range levels {
		switch i {
		case 0:
			lists = append(lists, levels[1:])
		case len(levels):
			lists = append(lists, levels[:len(levels)-1])
		default:
			newlist := make([]int, i, len(levels)-1)
			copy(newlist, levels[:i])
			newlist = append(newlist, levels[i+1:]...)
			lists = append(lists, newlist)
		}
	}
	return
}

func isSafeWithDampener(levels []int) bool {
	if isSafe(levels) {
		return true
	}
	for _, list := range removeOne(levels) {
		if isSafe(list) {
			return true
		}
	}
	return false
}

func safeLevels(filename string, checkFn func([]int) bool) (count int) {
	for _, line := range utils.ReadLines(filename) {
		levels := utils.StringToNumbers(line)
		if checkFn(levels) {
			count += 1
		}
	}
	return
}

func part1(filename string) string {
	return strconv.Itoa(safeLevels(filename, isSafe))
}

func part2(filename string) string {
	return strconv.Itoa(safeLevels(filename, isSafeWithDampener))
}
