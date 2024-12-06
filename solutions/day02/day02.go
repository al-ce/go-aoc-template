package day02

import (
	"fmt"

	h "go-aoc-template/internal/helpers"
)

// Check that levels are all increasing/decreasing and change isn't too big
func checkReport(report []int) int {
	direction := h.GetSignOfDiff(report[0], report[1])
	if direction == 0 {
		return -1
	}
	for i := 0; i <= len(report)-2; i++ {
		this, next := report[i], report[i+1]
		thisDirection := h.GetSignOfDiff(this, next)
		thisDiff := h.GetAbsDiff(this, next)
		if thisDirection != direction || thisDiff > 3 {
			return -1
		}
	}
	return 0
}

func PartOne(lines []string) string {
	safeCount := len(lines)
	for _, reportStr := range lines {
		report := h.ParseIntString(reportStr, " ")
		safeCount += checkReport(report)
	}
	return fmt.Sprintf("%d", safeCount)
}

func PartTwo(lines []string) string {
	return "TODO"
}
