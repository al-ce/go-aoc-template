package day07

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	h "go-aoc-template/internal/helpers"
)

type SolutionFunc func([]string) string

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
	concAddOps := append([]int{ops[0] + ops[1]}, ops[2:]...)
	if evalConcatEquation(value, concAddOps) == value {
		return value
	}
	concMulOps := append([]int{ops[0] * ops[1]}, ops[2:]...)
	if evalConcatEquation(value, concMulOps) == value {
		return value
	}

	concConcOps := append([]int{concat(ops[0], ops[1])}, ops[2:]...)
	return evalConcatEquation(value, concConcOps)
}

func partOneSolver(equation string) int {

		split := strings.Split(equation, ": ")
		value, _ := strconv.Atoi(split[0])
		operands := h.ParseIntString(split[1], " ")
		if result := evalEquation(value, operands); result == value { return value }
	return 0
}

func partTwoSolver(equation string) int {
	split := strings.Split(equation, ": ")
	value, _ := strconv.Atoi(split[0])
	operands := h.ParseIntString(split[1], " ")
	if result := evalEquation(value, operands); result == value {
		return value
	} else if result := evalConcatEquation(value, operands); result == value {
		return value
	}
	return 0
}

func parallelize(lines []string, solver func(equation string) int) int {
	values := make(chan int, len(lines))

	var wg sync.WaitGroup

	for _, equation := range lines {

		wg.Add(1)
		go func() {
			defer wg.Done()
			values <- solver(equation)
		}()

	}
	go func() {
		wg.Wait()
		close(values)
	}()

	sum := 0
	for value := range values {
		sum += value
	}
	return sum
}

func PartOne(lines []string) string {
	sum := parallelize(lines, partOneSolver)
	return fmt.Sprint(sum)
}

func PartTwo(lines []string) string {
	sum := parallelize(lines, partTwoSolver)
	return fmt.Sprint(sum)
}
