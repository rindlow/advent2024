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
	free := 0
	file := len(disk) - 1
	for disk[free] >= 0 {
		free += 1
	}
	for disk[file] < 0 {
		file -= 1
	}
	for free < file {
		disk[free] = disk[file]
		disk[file] = -1
		for disk[free] >= 0 {
			free += 1
		}
		for disk[file] < 0 {
			file -= 1
		}
	}
	for i := 0; i < len(disk)-1; i++ {
		if disk[i] >= 0 {
			sum += i * disk[i]
		}
	}
	return
}

func checksumFile(filename string) (sum int) {
	disk := []Item{}
	fileid := 0
	for _, line := range utils.ReadLines(filename) {
		for pos, char := range line {
			length := int(char - '0')
			if pos%2 == 0 {
				disk = append(disk, Item{false, length, pos / 2})
				fileid = pos / 2
			} else {
				disk = append(disk, Item{true, length, -1})
			}
		}
	}
	file := len(disk) - 1
	for disk[file].Free {
		file -= 1
	}
	for {
		for i := 0; i < file; i++ {
			if disk[i].Free && disk[i].Length >= disk[file].Length {
				rest := disk[i].Length - disk[file].Length
				disk[i].Free = false
				disk[i].Length = disk[file].Length
				disk[i].FileID = disk[file].FileID
				disk[file].Free = true
				disk[file].FileID = -1
				if rest > 0 {
					disk = slices.Insert(disk, i+1, Item{true, rest, -1})
				}
				break
			}
		}
		fileid -= 1
		file -= 1
		for disk[file].FileID != fileid {
			file -= 1
		}
		if disk[file].FileID == 0 {
			break
		}
	}
	index := 0
	for i := 0; i < len(disk)-1; i++ {
		if disk[i].Free {
			for j := 0; j < disk[i].Length; j++ {
			}
			index += disk[i].Length
		} else {
			for j := 0; j < disk[i].Length; j++ {
				sum += index * disk[i].FileID
				index += 1
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
