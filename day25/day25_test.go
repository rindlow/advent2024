package day25

import (
	"testing"
)

func TestPart1(t *testing.T) {
	testCases := []struct {
		filename string
		expected string
	}{
		{filename: "input.txt", expected: "3"},
		{filename: "../input/day25.txt", expected: "3451"},
	}
	for _, tc := range testCases {
		value := Part1(tc.filename)
		if value != tc.expected {
			t.Fatalf("part1: got %s, want %s", value, tc.expected)
		}
	}
}

// func TestPart2(t *testing.T) {
// 	testCases := []struct {
// 		filename string
// 		expected string
// 	}{
// 		{filename: "input.txt", expected: "Answer"},
// 		// {filename: "../input/day25.txt", expected: "Answer"},
// 	}
// 	for _, tc := range testCases {
// 		value := Part2(tc.filename)
// 		if value != tc.expected {
// 			t.Fatalf("part2: got %s, want %s", value, tc.expected)
// 		}
// 	}
// }
