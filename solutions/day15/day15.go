package day15

import (
	"fmt"
	"strings"
	"time"
)

const PRINT = false

var moveMap = map[byte]Coord{
	'v': {0, 1},
	'^': {0, -1},
	'<': {-1, 0},
	'>': {1, 0},
}

type Coord struct {
	x int
	y int
}

type Warehouse struct {
	lines [][]byte
	limit int
}

func NewWarehouse(lines []string) Warehouse {
	var _lines [][]byte
	for _, line := range lines {
		_lines = append(_lines, []byte(line))
	}
	return Warehouse{_lines, len(lines)}
}

func (w Warehouse) get(c Coord) byte {
	return w.lines[c.y][c.x]
}

func (w *Warehouse) set(c Coord, char byte) {
	w.lines[c.y][c.x] = char
}

func (w Warehouse) print() {
	fmt.Println()
	for _, row := range w.lines {
		for _, char := range row {
			fmt.Print(string(char))
		}
		fmt.Println()
	}
}

func (c Coord) getNext(move byte) Coord {
	movement := moveMap[move]
	return Coord{c.x + movement.x, c.y + movement.y}
}

func (w Warehouse) findStart() Coord {
	for y := range len(w.lines) {
		for x := range len(w.lines) {
			if w.lines[y][x] == '@' {
				return Coord{x, y}
			}
		}
	}
	return Coord{-1, -1}
}

func PartOne(lines []string) string {
	input := strings.Split(strings.Join(lines, "\n"), "\n\n")
	_warehouse, directions := input[0], input[1]
	directions = strings.ReplaceAll(directions, "\n", "")
	w := NewWarehouse(strings.Split(_warehouse, "\n"))

	robot := w.findStart()
	var delta byte

	for i := 0; i <= len(directions)-1; i++ {

		delta = directions[i]

		if PRINT {
			fmt.Print("\nmove: ", string(delta))
			w.print()
			time.Sleep(time.Millisecond * 250)
		}

		// Attempt to move and push boxes

		// Traverse any boxes to find potential empty space
		next := robot.getNext(delta)
		cursor := next
		for w.get(cursor) == 'O' {
			cursor = cursor.getNext(delta)
		}

		// If we haven't run into a wall
		if w.get(cursor) != '#' {
			w.set(robot, '.')
			// Move box to empty spot, or empty spot to itself
			w.set(cursor, w.get(next))
			// Move the robot to the next spot
			w.set(next, '@')
			// Update robot's coords
			robot = next
		}
	}

	// Calculate sum of GPS coordinates
	total := 0
	for y := range w.limit {
		for x := range w.limit {
			if w.get(Coord{x, y}) == 'O' {
				total += y*100 + x
			}
		}
	}

	return fmt.Sprintf("%d", total)
}

func PartTwo(lines []string) string {
	return "TODO"
}
