package day22

import (
	"strconv"

	"github.com/rindlow/aoc-utils"
)

func nextRandom(secret int) int {
	secret = ((secret << 6) ^ secret) % 16777216
	secret = ((secret >> 5) ^ secret) % 16777216
	secret = ((secret << 11) ^ secret) % 16777216
	return secret
}

func sum2000(filename string) (sum int) {
	for _, num := range utils.ReadLinesAsInt(filename) {
		for i := 0; i < 2000; i++ {
			num = nextRandom(num)
		}
		sum += num
	}
	return
}

func maxBananas(filename string) (max int) {
	sequences := make(map[[4]int]int)
	for _, num := range utils.ReadLinesAsInt(filename) {
		diffs := [4]int{}
		seen := make(map[[4]int]bool)
		last := num % 10
		for i := 0; i < 2000; i++ {
			num = nextRandom(num)
			price := num % 10
			diffs[i%4] = price - last
			last = price
			if i >= 3 {
				sequence := [4]int{
					diffs[(i-3)%4], diffs[(i-2)%4],
					diffs[(i-1)%4], diffs[i%4]}
				if !seen[sequence] {
					sequences[sequence] += price
					seen[sequence] = true
				}
			}
		}
	}
	for _, bananas := range sequences {
		if bananas > max {
			max = bananas
		}
	}
	return
}

func Part1(filename string) string {
	return strconv.Itoa(sum2000(filename))
}

func Part2(filename string) string {
	return strconv.Itoa(maxBananas(filename))
}
