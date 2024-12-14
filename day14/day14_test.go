package day14

import (
	"testing"
)

func TestPart1(t *testing.T) {
	testCases := []struct {
		filename   string
		lenX, lenY int
		expected   int
	}{
		{filename: "input.txt", lenX: 11, lenY: 7, expected: 12},
		{filename: "../input/day14.txt", lenX: 101, lenY: 103, expected: 232589280},
	}
	for _, tc := range testCases {
		value := safetyFactor(tc.filename, tc.lenX, tc.lenY)
		if value != tc.expected {
			t.Fatalf("part1: got %d, want %d", value, tc.expected)
		}
	}
}

func TestPart2(t *testing.T) {
	testCases := []struct {
		filename string
		expected string
	}{
		{filename: "../input/day14.txt", expected: "7569"},
	}
	for _, tc := range testCases {
		value := Part2(tc.filename)
		if value != tc.expected {
			t.Fatalf("part2: got %s, want %s", value, tc.expected)
		}
	}
}
