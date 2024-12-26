package day25

import (
	"strconv"

	"github.com/rindlow/aoc-utils"
)

const (
	Out = iota
	InKey
	InLock
)

func pairsFit(filename string) (fit int) {
	keys := [][5]int{}
	locks := [][5]int{}
	state := Out
	current := [5]int{}
	for _, line := range utils.ReadLines(filename) {
		switch {
		case line == "":
			switch state {
			case InKey:
				keys = append(keys, current)
			case InLock:
				locks = append(locks, current)
			}
			current = [5]int{}
			state = Out
			continue
		case line == "#####":
			if state == Out {
				state = InLock
				continue
			}
		case line == ".....":
			if state == Out {
				state = InKey
				current = [5]int{-1, -1, -1, -1, -1}
				continue
			}
		}
		for pin, char := range line {
			if char == '#' {
				switch state {
				case InLock:
					current[pin] += 1
				case InKey:
					current[pin] += 1
				}
			}
		}
	}
	switch state {
	case InKey:
		keys = append(keys, current)
	case InLock:
		locks = append(locks, current)
	}

	for _, key := range keys {
		for _, lock := range locks {
			overlap := false
			for pin := 0; pin < 5; pin++ {
				if key[pin]+lock[pin] > 5 {
					overlap = true
					break
				}
			}
			if !overlap {
				fit += 1
			}
		}
	}
	return
}

func Part1(filename string) string {
	return strconv.Itoa(pairsFit(filename))
}

func Part2(filename string) string {
	return "Not implemented"
}
