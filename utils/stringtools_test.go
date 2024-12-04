package utils

import (
	"slices"
	"testing"
)

func TestReadLines(t *testing.T) {
	testCases := []struct {
		filename string
		expected []string
	}{
		{filename: "testfile.txt", expected: []string{"foobar", "gazonk"}},
	}
	for _, tc := range testCases {
		value := ReadLines(tc.filename)
		if !slices.Equal(value, tc.expected) {
			t.Fatalf("Abs: got %v, want %v", value, tc.expected)
		}
	}
}

func TestStringToNumbers(t *testing.T) {
	testCases := []struct {
		x        string
		expected []int
	}{
		{x: "17", expected: []int{17}},
		{x: "6 17 23", expected: []int{6, 17, 23}},
	}
	for _, tc := range testCases {
		value := StringToNumbers(tc.x)
		if !slices.Equal(value, tc.expected) {
			t.Fatalf("Abs: got %v, want %v", value, tc.expected)
		}
	}
}
