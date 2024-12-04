package day4

import (
	"testing"
)

func TestPart1(t *testing.T) {
	testCases := []struct {
		filename string
		expected string
	}{
		{filename: "input.txt", expected: "18"},
		{filename: "../input/day4.txt", expected: "2344"},
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
		{filename: "input.txt", expected: "9"},
		{filename: "../input/day4.txt", expected: "1815"},
	}
	for _, tc := range testCases {
		value := Part2(tc.filename)
		if value != tc.expected {
			t.Fatalf("part2: got %s, want %s", value, tc.expected)
		}
	}
}
