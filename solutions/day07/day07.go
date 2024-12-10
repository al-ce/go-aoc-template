package day07

import (
	"fmt"
	"strconv"
	"strings"

	h "go-aoc-template/internal/helpers"
)

func testEquation(value int, ops []int) int {
	if len(ops) == 1 {
		return ops[0]
	}
	last := ops[len(ops)-1]
	adds := testEquation(value-last, ops[:len(ops)-1]) + last
	if adds == value {
		return adds
	} else {
		muls := testEquation(value/last, ops[:len(ops)-1]) * last
		return muls
	}
}

func PartOne(lines []string) string {
	sum := 0
	for _, equation := range lines {
		split := strings.Split(equation, ": ")
		value, _ := strconv.Atoi(split[0])
		operands := h.ParseIntString(split[1], " ")
		if result := testEquation(value, operands); result == value {
			sum = sum + int(value)
		} else {
		}
	}
	return fmt.Sprint(sum)
}

func PartTwo(lines []string) string {
	return "TODO"
}
