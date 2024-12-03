package day3

import (
	"log"
	"os"
	"regexp"
	"strconv"
	// "github.com/rindlow/aoc-utils"
)

func sumMultiplications(filename string) (sum int) {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("%s: %q", filename, err)
	}
	re := regexp.MustCompile(`mul\((?P<A>\d{1,3}),(?P<B>\d{1,3})\)`)
	for _, match := range re.FindAllStringSubmatch(string(data), -1) {
		a, err := strconv.Atoi(match[1])
		if err != nil {
			log.Fatalf("%s: %q", match[1], err)
		}
		b, err := strconv.Atoi(match[2])
		if err != nil {
			log.Fatalf("%s: %q", match[2], err)
		}
		sum += a * b
	}
	return
}

func sumConditionalMultiplications(filename string) (sum int) {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("%s: %q", filename, err)
	}
	var enabled = true
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|(do\(\))|(don't\(\))`)
	for _, match := range re.FindAllStringSubmatch(string(data), -1) {
		switch {
		case match[1] != "" && match[2] != "" && enabled:
			a, err := strconv.Atoi(match[1])
			if err != nil {
				log.Fatalf("%s: %q", match[1], err)
			}
			b, err := strconv.Atoi(match[2])
			if err != nil {
				log.Fatalf("%s: %q", match[2], err)
			}
			sum += a * b
		case match[3] == "do()":
			enabled = true
		case match[4] == "don't()":
			enabled = false
		}
	}
	return
}

func Part1(filename string) string {
	return strconv.Itoa(sumMultiplications(filename))
}

func Part2(filename string) string {
	return strconv.Itoa(sumConditionalMultiplications(filename))
}
