package day19

import (
	"strconv"
	"strings"

	"github.com/rindlow/aoc-utils"
)

func readTowelsDesigns(filename string) (towels, designs []string) {
	lines := utils.ReadLines(filename)
	towels = strings.Split(lines[0], ", ")
	designs = lines[2:]
	return
}

func possible(design string, towels *[]string, cache map[string]int) int {
	if len(design) == 0 {
		return 1
	}
	hit, ok := cache[design]
	if ok {
		return hit
	}
	p := 0
	for _, towel := range *towels {
		suffix, found := strings.CutPrefix(design, towel)
		if found {
			p += possible(suffix, towels, cache)
		}
	}
	cache[design] = p
	return p
}

func possibleDesigns(filename string) (n int) {
	towels, designs := readTowelsDesigns(filename)
	cache := make(map[string]int)
	for _, design := range designs {
		if possible(design, &towels, cache) > 0 {
			n += 1
		}
	}
	return
}

func differentDesigns(filename string) (n int) {
	towels, designs := readTowelsDesigns(filename)
	cache := make(map[string]int)
	for _, design := range designs {
		n += possible(design, &towels, cache)
	}
	return
}

func Part1(filename string) string {
	return strconv.Itoa(possibleDesigns(filename))
}

func Part2(filename string) string {
	return strconv.Itoa(differentDesigns(filename))
}
