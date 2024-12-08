package utils

import (
	"testing"
)

func TestAbs(t *testing.T) {
	testCases := []struct {
		x        int
		expected int
	}{
		{x: 17, expected: 17},
		{x: -17, expected: 17},
	}
	for _, tc := range testCases {
		value := Abs(tc.x)
		if value != tc.expected {
			t.Fatalf("Abs: got %d, want %d", value, tc.expected)
		}
	}
}

func TestSign(t *testing.T) {
	testCases := []struct {
		x        int
		expected int
	}{
		{x: 17, expected: 1},
		{x: -17, expected: -1},
		{x: 0, expected: 0},
	}
	for _, tc := range testCases {
		value := Sign(tc.x)
		if value != tc.expected {
			t.Fatalf("Abs: got %d, want %d", value, tc.expected)
		}
	}
}

func TestMax(t *testing.T) {
	testCases := []struct {
		x, y     int
		expected int
	}{
		{x: 17, y: 18, expected: 18},
		{x: -17, y: -18, expected: -17},
		{x: 0, y: 0, expected: 0},
	}
	for _, tc := range testCases {
		value := Max(tc.x, tc.y)
		if value != tc.expected {
			t.Fatalf("Max: got %d, want %d", value, tc.expected)
		}
	}
}
