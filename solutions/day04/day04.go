package day04

import (
	"fmt"

	h "go-aoc-template/internal/helpers"
)

const (
	xmas = "XMAS"
	MAS  = "MAS"
)

type Direction struct {
	y int
	x int
}

type Coords struct {
	row int
	col int
}

var ordinals = []Direction{
	{+1, +1}, // NE
	{+1, -1}, // SE
	{-1, +1}, // NW
	{-1, -1}, // SW
}

var cardinals = []Direction{
	{0, -1}, // N
	{0, +1}, // S
	{+1, 0}, // E
	{-1, 0}, // W
}

func isDirectionInBounds(lines []string, word string, c Coords, dir Direction) bool {
	min, max := len(word)-1, len(lines)-len(word)
	return !(!h.InMatrix(c.row, c.col, lines) ||
		(dir.y == -1 && c.row < min) ||
		(dir.y == +1 && c.row > max) ||
		(dir.x == -1 && c.col < min) ||
		(dir.x == +1 && c.col > max))
}

func findWord(lines []string, word string, c Coords, dir Direction) int {
	if !isDirectionInBounds(lines, word, c, dir) {
		return 0
	}
	for i, char := range word {
		next := Coords{c.row + dir.y*i, c.col + dir.x*i}
		curr := rune(lines[next.row][next.col])
		if curr != char {
			return 0
		}
	}
	return 1
}

func PartOne(lines []string) string {
	directions := append(ordinals, cardinals...)
	sum := 0
	for rowIdx, row := range lines {
		for colIdx := range len(row) {
			c := Coords{rowIdx, colIdx}
			for _, dir := range directions {
				sum += findWord(lines, xmas, c, dir)
			}
		}
	}
	return fmt.Sprintf("%d", sum)
}

func PartTwo(lines []string) string {
	sum := 0
	for rowIdx, row := range lines {
		for colIdx := range len(row) {
			c := Coords{rowIdx, colIdx}
			if string(lines[c.row][c.col]) != "A" {
				continue
			}
			sub := 0
			for _, dir := range ordinals {
				newCoords := Coords{c.row + dir.y, c.col + dir.x}
				delta := Direction{-dir.y, -dir.x}
				sub += findWord(lines, MAS, newCoords, delta)
				if sub >= 2 {
					sum += 1
					break
				}
			}
		}
	}
	return fmt.Sprintf("%d", sum)
}
