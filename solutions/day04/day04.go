package day04

import "fmt"

const xmas = "XMAS"

type Direction struct {
	x int
	y int
}

var directions = []Direction{
	{0, -1},  // N
	{0, +1},  // S
	{+1, 0},  // E
	{-1, 0},  // W
	{+1, +1}, // NE
	{+1, -1}, // SE
	{-1, +1}, // NW
	{-1, -1}, // SW
}

func isDirectionValid(lines []string, rowIdx, colIdx int, dir Direction) bool {
	min, max := len(xmas)-1, len(lines)-len(xmas)
	return !((dir.y == -1 && rowIdx < min) ||
		(dir.y == +1 && rowIdx > max) ||
		(dir.x == -1 && colIdx < min) ||
		(dir.x == +1 && colIdx > max))
}

func findXmas(lines []string, rowIdx, colIdx int, dir Direction) int {
	for i, char := range xmas {
		newRow := rowIdx + dir.y*i
		newCol := colIdx + dir.x*i
		curr := rune(lines[newRow][newCol])
		if curr != char {
			return 0
		}
	}
	return 1
}

func checkCoordinate(lines []string, rowIdx, colIdx int) int {
	found := 0
	for _, dir := range directions {
		valid := isDirectionValid(lines, rowIdx, colIdx, dir)
		if !valid {
			continue
		}
		found += findXmas(lines, rowIdx, colIdx, dir)
	}
	return found
}

func PartOne(lines []string) string {
	sum := 0
	for rowIdx, row := range lines {
		for colIdx := range len(row) {
			sum += checkCoordinate(lines, rowIdx, colIdx)
		}
	}

	return fmt.Sprintf("%d", sum)
}

func PartTwo(lines []string) string {
	return "TODO"
}
