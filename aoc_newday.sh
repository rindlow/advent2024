#!/bin/bash

if [ $# -lt 1 ]; then
    echo "USAGE: $0 dayno"
    exit 1
fi
Project="advent2024"
Module="day$1"
ModulePath="rindlow.se/$Project/$Module"
TestModule="${Module}_test"
SrcFile="$Module.go"
TestFile="$TestModule.go"

ed go.work << EOF
/utils/
i
    ./$Module
.
w
q
EOF

ed cmd/main.go << EOF
/Insert more modules above this line/
i
    "$ModulePath"
.
/Insert more days above this line/
i
		$1: {$Module.Part1, $Module.Part2},
.
w
q
EOF

mkdir $Module
cd $Module
go mod init rindlow.se/$Project/$Module
touch input.txt

cat - > $SrcFile << EOF
package $Module

import (
//	"github.com/rindlow/aoc-utils"
)

func Part1(filename string) string {
	return "Not implemented"
}

func Part2(filename string) string {
	return "Not implemented"
}
EOF

cat - > $TestFile << EOF
package $Module

import (
	"testing"
)
func TestPart1(t *testing.T) {
	testCases := []struct {
		filename string
		expected string
	}{
		{filename: "input.txt", expected: "Answer"},
		// {filename: "../input/$Module.txt", expected: "Answer"},
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
// 		// {filename: "../input/$Module.txt", expected: "Answer"},
// 	}
// 	for _, tc := range testCases {
// 		value := Part2(tc.filename)
// 		if value != tc.expected {
// 			t.Fatalf("part2: got %s, want %s", value, tc.expected)
// 		}
// 	}
// }

EOF
