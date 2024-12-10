package day07

import (
	"fmt"
	"strconv"
	"strings"

	h "go-aoc-template/internal/helpers"
)

func evalEquation(value int, ops []int) int {
	if len(ops) == 1 {
		return ops[0]
	}
	last := ops[len(ops)-1]
	adds := evalEquation(value-last, ops[:len(ops)-1]) + last
	if adds == value {
		return adds
	} else {
		muls := evalEquation(value/last, ops[:len(ops)-1]) * last
		return muls
	}
}

func concat(v1, v2 int) int {
	v, _ := strconv.Atoi(fmt.Sprintf("%d%d", v1, v2))
	return v
}

func evalConcatEquation(value int, ops []int) int {
	if len(ops) == 1 {
		return ops[0]
	}
	concAddOps := append([]int{ops[0]+ops[1]}, ops[2:]...)
	if evalConcatEquation(value, concAddOps) == value {
		return value
	}
	concMulOps := append([]int{ops[0]*ops[1]}, ops[2:]...)
	if evalConcatEquation(value, concMulOps) == value {
		return value
	}

	concConcOps := append([]int{concat(ops[0], ops[1])}, ops[2:]...)
	return evalConcatEquation(value, concConcOps)
}

func PartOne(lines []string) string {
	sum := 0
	for _, equation := range lines {
		split := strings.Split(equation, ": ")
		value, _ := strconv.Atoi(split[0])
		operands := h.ParseIntString(split[1], " ")
		if result := evalEquation(value, operands); result == value {
			sum += value
		} else {
		}
	}
	return fmt.Sprint(sum)
}

func PartTwo(lines []string) string {
	sum := 0
	for _, equation := range lines {
		split := strings.Split(equation, ": ")
		value, _ := strconv.Atoi(split[0])
		operands := h.ParseIntString(split[1], " ")
		if result := evalEquation(value, operands); result == value {
			sum += value
		} else if result := evalConcatEquation(value, operands); result == value {
			sum += value
		}
	}
	return fmt.Sprint(sum)
}
