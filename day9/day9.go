package day9

import (
	"slices"
	"strconv"

	"github.com/rindlow/aoc-utils"
)

type Item struct {
	Free   bool
	Length int
	FileID int
}

func checksum(filename string) (sum int) {
	disk := make(map[int]int)
	i := 0
	for _, line := range utils.ReadLines(filename) {
		for pos, char := range line {
			length := int(char - '0')
			for block := 0; block < length; block++ {
				if pos%2 == 0 {
					disk[i] = pos / 2
				} else {
					disk[i] = -1
				}
				i += 1
			}
		}
	}
	freePtr := 0
	for disk[freePtr] >= 0 {
		freePtr += 1
	}
	filePtr := len(disk) - 1
	for disk[filePtr] < 0 {
		filePtr -= 1
	}
	for freePtr < filePtr {
		disk[freePtr] = disk[filePtr]
		disk[filePtr] = -1
		for disk[freePtr] >= 0 {
			freePtr += 1
		}
		for disk[filePtr] < 0 {
			filePtr -= 1
		}
	}
	for pos := 0; pos < len(disk)-1; pos++ {
		if disk[pos] >= 0 {
			sum += pos * disk[pos]
		}
	}
	return
}

func checksumFile(filename string) (sum int) {
	disk := []Item{}
	fileId := 0
	for _, line := range utils.ReadLines(filename) {
		for pos, char := range line {
			length := int(char - '0')
			if pos%2 == 0 {
				disk = append(disk, Item{false, length, pos / 2})
				fileId = pos / 2
			} else {
				disk = append(disk, Item{true, length, -1})
			}
		}
	}
	filePtr := len(disk) - 1
	for disk[filePtr].Free {
		filePtr -= 1
	}
	for {
		for i := 0; i < filePtr; i++ {
			if disk[i].Free && disk[i].Length >= disk[filePtr].Length {
				rest := disk[i].Length - disk[filePtr].Length
				disk[i].Free = false
				disk[i].Length = disk[filePtr].Length
				disk[i].FileID = disk[filePtr].FileID
				disk[filePtr].Free = true
				disk[filePtr].FileID = -1
				if rest > 0 {
					disk = slices.Insert(disk, i+1, Item{true, rest, -1})
				}
				break
			}
		}

		fileId -= 1
		filePtr -= 1
		for disk[filePtr].FileID != fileId {
			filePtr -= 1
		}
		if disk[filePtr].FileID == 0 {
			break
		}
	}
	pos := 0
	for i := 0; i < len(disk)-1; i++ {
		if disk[i].Free {
			pos += disk[i].Length
		} else {
			for j := 0; j < disk[i].Length; j++ {
				sum += pos * disk[i].FileID
				pos += 1
			}
		}
	}
	return
}

func Part1(filename string) string {
	return strconv.Itoa(checksum(filename))
}

func Part2(filename string) string {
	return strconv.Itoa(checksumFile(filename))
}
