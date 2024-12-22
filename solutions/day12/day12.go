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

func getAreaAndSides(lines []string, row, col int, visited map[Coord]int) (int, int) {
	area := 1
	sides := 4
	visited[Coord{row, col}] = 1
	char := lines[row][col]
	otherSides := 0

	// If element above is the same type
	if row > 0 && lines[row-1][col] == char {
		sides-- // remove side above

		if col == 0 || (lines[row-1][col-1] != char && lines[row][col-1] != char) {
			sides-- // remove left side already accounted for
		}
		if col == len(lines[0])-1 || (lines[row-1][col+1] != char && lines[row][col+1] != char){
			sides-- // remove right side already accounted for
		}

		if _, exists := visited[Coord{row - 1, col}]; !exists {
			a, p := getAreaAndSides(lines, row-1, col, visited)
			area += a
			otherSides += p
		}
	}

	// If element below is the same type
	if row < len(lines)-1 && lines[row+1][col] == char {
		sides-- // remove side below
		if _, exists := visited[Coord{row + 1, col}]; !exists {
			a, p := getAreaAndSides(lines, row+1, col, visited)
			area += a
			otherSides += p
		}
	}

	// If element to the left is the same type
	if col > 0 && lines[row][col-1] == char {
		sides-- // remove side to the left

		if row == 0 || (lines[row-1][col-1] != char && lines[row-1][col] != char){
			sides-- // remove side above already accounted for
		}
		if row == len(lines)-1 || (lines[row+1][col-1] != char && lines[row+1][col] != char){
			sides-- // remove side below already accounted for
		}

		if _, exists := visited[Coord{row, col - 1}]; !exists {
			a, p := getAreaAndSides(lines, row, col-1, visited)
			area += a
			otherSides += p
		}
	}

	// If element to the right is the same type
	if col < len(lines[0])-1 && lines[row][col+1] == char {
		sides--
		if _, exists := visited[Coord{row, col + 1}]; !exists {
			a, p := getAreaAndSides(lines, row, col+1, visited)
			area += a
			otherSides += p
		}
	}

	fmt.Println(row, col, sides, otherSides, area)
	sides += otherSides
	return area, sides
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
	println()
	visited := make(map[Coord]int)
	price := 0

	for row, line := range lines {
		for col := range line {
			if _, exists := visited[Coord{row, col}]; !exists {
				a, p := getAreaAndSides(lines, row, col, visited)
				price += a * p
			}
		}
	}

	return fmt.Sprintf("%d", price)
}
