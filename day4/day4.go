package day4

import (
	"strconv"

	"github.com/rindlow/aoc-utils"
)

type WordSearch []string

func (data WordSearch) readDirection(row, col, dRow, dCol, n int) string {
	if n == 0 || row < 0 || row >= len(data) || col < 0 || col >= len(data[0]) {
		return ""
	}
	return string(data[row][col]) + data.readDirection(row+dRow, col+dCol, dRow, dCol, n-1)
}

func countXmas(filename string) (count int) {
	data := WordSearch(utils.ReadLines(filename))
	for rowNo, row := range data {
		for colNo, letter := range row {
			if letter == 'X' {
				for dRow := -1; dRow < 2; dRow++ {
					for dCol := -1; dCol < 2; dCol++ {
						if dRow != 0 || dCol != 0 {
							if data.readDirection(rowNo+dRow, colNo+dCol, dRow, dCol, 3) == "MAS" {
								count += 1
							}
						}
					}
				}
			}
		}
	}
	return
}

func countMasX(filename string) (count int) {
	data := WordSearch(utils.ReadLines(filename))
	for rowNo, row := range data {
		for colNo, letter := range row {
			if letter == 'A' {
				var mas = 0
				for dRow := -1; dRow < 2; dRow += 2 {
					for dCol := -1; dCol < 2; dCol += 2 {
						if data.readDirection(rowNo-dRow, colNo-dCol, dRow, dCol, 3) == "MAS" {
							mas += 1
						}
					}
				}
				if mas == 2 {
					count += 1
				}
			}
		}
	}
	return
}

func Part1(filename string) string {
	return strconv.Itoa(countXmas(filename))
}

func Part2(filename string) string {
	return strconv.Itoa(countMasX(filename))
}
