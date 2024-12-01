package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	testCases := []struct {
		filename string
		expected string
	}{
		{filename: "input.txt", expected: "11"},
		{filename: "../input/day1.txt", expected: "1506483"},
	}
	for _, tc := range testCases {
		value := part1(tc.filename)
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
		{filename: "input.txt", expected: "31"},
		{filename: "../input/day1.txt", expected: "23126924"},
	}
	for _, tc := range testCases {
		value := part2(tc.filename)
		if value != tc.expected {
			t.Fatalf("part2: got %s, want %s", value, tc.expected)
		}
	}
}
