package day19

import (
	"testing"
)

func TestPart1(t *testing.T) {
	testCases := []struct {
		filename string
		expected string
	}{
		{filename: "input.txt", expected: "6"},
		{filename: "../input/day19.txt", expected: "300"},
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
		{filename: "input.txt", expected: "16"},
		{filename: "../input/day19.txt", expected: "624802218898092"},
	}
	for _, tc := range testCases {
		value := Part2(tc.filename)
		if value != tc.expected {
			t.Fatalf("part2: got %s, want %s", value, tc.expected)
		}
	}
}
