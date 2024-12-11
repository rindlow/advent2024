package day11

import (
	"math"
	"strconv"

	"github.com/rindlow/aoc-utils"
)

type CacheKey struct {
	num, iter int
}
type Cache map[CacheKey]int

func evolve(stone, n int, cache Cache) (length int) {
	if n == 0 {
		return 1
	}
	key := CacheKey{stone, n}
	stones, ok := cache[key]
	if ok {
		return stones
	}
	nDigits := int(math.Floor(math.Log10(float64(stone)))) + 1
	switch {
	case stone == 0:
		length = evolve(1, n-1, cache)
	case nDigits%2 == 0:
		divisor := int(math.Pow10(nDigits / 2))
		a := stone / divisor
		b := stone % divisor
		length = evolve(a, n-1, cache) + evolve(b, n-1, cache)
	default:
		length = evolve(stone*2024, n-1, cache)
	}
	cache[key] = length
	return
}

func numStones(filename string, n int) (sum int) {
	cache := make(Cache)
	for _, stone := range utils.StringToNumbers(utils.ReadLines(filename)[0]) {
		sum += evolve(stone, n, cache)
	}
	return
}

func Part1(filename string) string {
	return strconv.Itoa(numStones(filename, 25))
}

func Part2(filename string) string {
	return strconv.Itoa(numStones(filename, 75))
}
