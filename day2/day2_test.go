package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	testCases := []struct {
		filename string
		expected string
	}{
		{filename: "input.txt", expected: "2"},
		{filename: "../input/day2.txt", expected: "279"},
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
		{filename: "input.txt", expected: "4"},
		{filename: "../input/day2.txt", expected: "343"},
	}
	for _, tc := range testCases {
		value := part2(tc.filename)
		if value != tc.expected {
			t.Fatalf("part2: got %s, want %s", value, tc.expected)
		}
	}
}
