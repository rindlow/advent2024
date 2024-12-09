package day9

import (
	"testing"
)

func TestPart1(t *testing.T) {
	testCases := []struct {
		filename string
		expected string
	}{
		{filename: "input.txt", expected: "1928"},
		{filename: "../input/day9.txt", expected: "6356833654075"},
	}
	for _, tc := range testCases {
		value := Part1(tc.filename)
		if value != tc.expected {
			t.Fatalf("part1: got %s, want %s", value, tc.expected)
		}
	}
}

func TestPart2(t *testing.T) {
	testCases := []struct {
		filename string
		expected string
	}{
		{filename: "input.txt", expected: "2858"},
		{filename: "../input/day9.txt", expected: "6389911791746"},
	}
	for _, tc := range testCases {
		value := Part2(tc.filename)
		if value != tc.expected {
			t.Fatalf("part2: got %s, want %s", value, tc.expected)
		}
	}
}
