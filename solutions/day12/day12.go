package day12

import "fmt"

type Coord struct {
	row int
	col int
}

var possibleSiblings = []Coord{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}

// Beginning from the top left unvisited garden plot, calculate the area and
// perimeter of all adjacent garden plots of the same type
func getGardenAreaAndPerimeter(lines []string, row, col int, visited map[Coord]struct{}) (int, int) {
	area := 1
	peri := 4 // assume a perimeter of 4 for each plot and subtract as needed
	visited[Coord{row, col}] = struct{}{}
	for _, sibling := range possibleSiblings {
		nextRow := row + sibling.row
		nextCol := col + sibling.col
		// Ensure in bounds
		if nextRow < 0 || nextRow >= len(lines) || nextCol < 0 || nextCol >= len(lines) {
			continue
		}
		// Don't count siblings in perimeter
		if lines[row][col] == lines[nextRow][nextCol] {
			peri--
			// Visit unvisited siblings
			if _, exists := visited[Coord{nextRow, nextCol}]; !exists {
				a, p := getGardenAreaAndPerimeter(lines, nextRow, nextCol, visited)
				area += a
				peri += p
			}
		}
	}
	return area, peri
}

func getAreaAndSides(lines []string, row, col int, visited map[Coord]struct{}) (int, int) {
	area := 1
	sides := 4
	visited[Coord{row, col}] = struct{}{}
	char := lines[row][col]

	// If element above is the same type
	if row > 0 && lines[row-1][col] == char {
		sides-- // remove side above

		// If at left edge or element above already counted left side
		if col == 0 || (lines[row-1][col-1] != char && lines[row][col-1] != char) {
			sides-- // remove left side already accounted for by abolve element
		}
		// If at right edge or element above already counted right side
		if col == len(lines[0])-1 || (lines[row-1][col+1] != char && lines[row][col+1] != char) {
			sides-- // remove right side already accounted for by abolve element
		}

		if _, exists := visited[Coord{row - 1, col}]; !exists {
			a, s := getAreaAndSides(lines, row-1, col, visited)
			area += a
			sides += s
		}
	}

	// If element below is the same type
	if row < len(lines)-1 && lines[row+1][col] == char {
		sides-- // remove side below
		if _, exists := visited[Coord{row + 1, col}]; !exists {
			a, p := getAreaAndSides(lines, row+1, col, visited)
			area += a
			sides += p
		}
	}

	// If element to the left is the same type
	if col > 0 && lines[row][col-1] == char {
		sides-- // remove side to the left

		// If at top edge or element above already counted top side
		if row == 0 || (lines[row-1][col-1] != char && lines[row-1][col] != char) {
			sides-- // remove side above already accounted for
		}
		// If at bottom edge or element above already counted bottom side
		if row == len(lines)-1 || (lines[row+1][col-1] != char && lines[row+1][col] != char) {
			sides-- // remove side below already accounted for
		}

		if _, exists := visited[Coord{row, col - 1}]; !exists {
			a, p := getAreaAndSides(lines, row, col-1, visited)
			area += a
			sides += p
		}
	}

	// If element to the right is the same type
	if col < len(lines[0])-1 && lines[row][col+1] == char {
		sides--
		if _, exists := visited[Coord{row, col + 1}]; !exists {
			a, p := getAreaAndSides(lines, row, col+1, visited)
			area += a
			sides += p
		}
	}

	return area, sides
}

func PartOne(lines []string) string {
	visited := make(map[Coord]struct{})
	price := 0

	for row, line := range lines {
		for col := range line {
			if _, exists := visited[Coord{row, col}]; !exists {
				a, p := getGardenAreaAndPerimeter(lines, row, col, visited)
				price += a * p
			}
		}
	}
	return fmt.Sprintf("%d", price)
}

func PartTwo(lines []string) string {
	visited := make(map[Coord]struct{})
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
