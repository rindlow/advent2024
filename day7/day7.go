package day7

import (
	"log"
	"strconv"
	"strings"

	"github.com/rindlow/aoc-utils"
)

type Checker func(int64, int64, []int64) bool

func existsAnswerSumMul(answer int64, first int64, rest []int64) bool {
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

func existsAnswerSumMulConc(answer int64, first int64, rest []int64) bool {
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
	concat := strconv.FormatInt(first, 10) + strconv.FormatInt(operand, 10)
	conc, err := strconv.ParseInt(concat, 10, 64)
	if err != nil {
		log.Fatalf("ParseInt(%s): %q", concat, err)
	}
	return conc <= answer && existsAnswerSumMulConc(answer, conc, rest)
}

func sumCalibration(filename string, checker Checker) (sum int64) {
	for _, line := range utils.ReadLines(filename) {
		colon := strings.Split(line, ": ")
		answer, err := strconv.ParseInt(colon[0], 10, 64)
		if err != nil {
			log.Fatalf("ParseInt(%s): %q", colon[0], err)
		}
		operands := []int64{}
		for _, digits := range strings.Split(colon[1], " ") {
			num, err := strconv.ParseInt(digits, 10, 64)
			if err != nil {
				log.Fatalf("ParseInt(%s): %q", digits, err)
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
	return strconv.FormatInt(sumCalibration(filename, existsAnswerSumMul), 10)
}

func Part2(filename string) string {
	return strconv.FormatInt(sumCalibration(filename, existsAnswerSumMulConc), 10)
}
