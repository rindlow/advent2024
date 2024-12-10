package day7

import (
	"log"
	"strconv"
	"strings"

	"github.com/rindlow/aoc-utils"
)

type Checker func(int, int, []int) bool

func existsAnswerSumMul(answer int, first int, rest []int) bool {
	if len(rest) == 0 {
		return answer == first
	}
	operand := rest[0]
	rest = rest[1:]
	sum := first + operand
	if sum <= answer && existsAnswerSumMul(answer, sum, rest) {
		return true
	}
	prod := first * operand
	return prod <= answer && existsAnswerSumMul(answer, prod, rest)
}

func existsAnswerSumMulConc(answer int, first int, rest []int) bool {
	if len(rest) == 0 {
		return answer == first
	}
	operand := rest[0]
	rest = rest[1:]

	sum := first + operand
	if sum <= answer && existsAnswerSumMulConc(answer, sum, rest) {
		return true
	}
	prod := first * operand
	if prod <= answer && existsAnswerSumMulConc(answer, prod, rest) {
		return true
	}
	concat := strconv.Itoa(first) + strconv.Itoa(operand)
	conc, err := strconv.Atoi(concat)
	if err != nil {
		log.Fatalf("ParseInt(%s): %q", concat, err)
	}
	return conc <= answer && existsAnswerSumMulConc(answer, conc, rest)
}

func sumCalibration(filename string, checker Checker) (sum int) {
	for _, line := range utils.ReadLines(filename) {
		colon := strings.Split(line, ": ")
		answer, err := strconv.Atoi(colon[0])
		if err != nil {
			log.Fatalf("Atoi(%s): %q", colon[0], err)
		}
		operands := []int{}
		for _, digits := range strings.Split(colon[1], " ") {
			num, err := strconv.Atoi(digits)
			if err != nil {
				log.Fatalf("Atoi(%s): %q", digits, err)
			}
			operands = append(operands, num)
		}
		if checker(answer, operands[0], operands[1:]) {
			sum += answer
		}
	}
	return
}

func Part1(filename string) string {
	return strconv.Itoa(sumCalibration(filename, existsAnswerSumMul))
}

func Part2(filename string) string {
	return strconv.Itoa(sumCalibration(filename, existsAnswerSumMulConc))
}
