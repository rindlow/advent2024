package day18

import (
	"testing"
)

func TestPart1(t *testing.T) {
	testCases := []struct {
		filename    string
		nBytes, max int
		expected    int
	}{
		{filename: "input.txt", nBytes: 12, max: 6, expected: 22},
		{filename: "../input/day18.txt", nBytes: 1024, max: 70, expected: 268},
	}
	for _, tc := range testCases {
		value := runShortestPath(tc.filename, tc.nBytes, tc.max)
		if value != tc.expected {
			t.Fatalf("part1: got %d, want %d", value, tc.expected)
		}
	}
}

func TestPart2(t *testing.T) {
	testCases := []struct {
		filename    string
		nBytes, max int
		expected    string
	}{
		{filename: "input.txt", nBytes: 12, max: 6, expected: "6,1"},
		{filename: "../input/day18.txt", nBytes: 1024, max: 70, expected: "64,11"},
	}
	for _, tc := range testCases {
		value := firstBlocking(tc.filename, tc.nBytes, tc.max)
		if value != tc.expected {
			t.Fatalf("part2: got %s, want %s", value, tc.expected)
		}
	}
}
