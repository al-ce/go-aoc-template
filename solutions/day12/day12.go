package day12

import "fmt"

type Coord struct {
	row int
	col int
}

func getSurroundingSiblings(lines []string, row, col int, visited map[Coord]int) (int, int) {
	area := 1
	peri := 4
	visited[Coord{row, col}] = 1
	char := lines[row][col]
	if row > 0 && lines[row-1][col] == char {
		peri--
		if _, exists := visited[Coord{row - 1, col}]; !exists {
			a, p := getSurroundingSiblings(lines, row-1, col, visited)
			area += a
			peri += p
		}
	}
	if row < len(lines)-1 && lines[row+1][col] == char {
		peri--
		if _, exists := visited[Coord{row + 1, col}]; !exists {
			a, p := getSurroundingSiblings(lines, row+1, col, visited)
			area += a
			peri += p
		}
	}
	if col > 0 && lines[row][col-1] == char {
		peri--
		if _, exists := visited[Coord{row, col - 1}]; !exists {
			a, p := getSurroundingSiblings(lines, row, col-1, visited)
			area += a
			peri += p
		}
	}
	if col < len(lines[0])-1 && lines[row][col+1] == char {
		peri--
		if _, exists := visited[Coord{row, col + 1}]; !exists {
			a, p := getSurroundingSiblings(lines, row, col+1, visited)
			area += a
			peri += p
		}
	}
	return area, peri
}

func PartOne(lines []string) string {
	visited := make(map[Coord]int)
	price := 0

	for row, line := range lines {
		for col := range line {
			if _, exists := visited[Coord{row, col}]; !exists {
				a, p := getSurroundingSiblings(lines, row, col, visited)
				price += a * p
			}
		}
	}
	return fmt.Sprintf("%d", price)
}

func PartTwo(lines []string) string {
	return "TODO"
}
