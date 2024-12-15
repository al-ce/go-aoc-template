package day10

import (
	"fmt"
	"strconv"
)

type Coord struct {
	row    int
	col    int
	height int
}

var PART int

type Delta struct {
	row int
	col int
}

var deltas = []Delta{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

type CoordSet map[Coord]struct{}

func (c CoordSet) add(coord Coord) {
	c[coord] = struct{}{}
}

func (c CoordSet) has(coord Coord) bool {
	_, exists := c[coord]
	return exists
}

// Make a coord from the given coordinates. Return an error if it would be OOB
func makecoord(lines []string, row, col int) (Coord, error) {
	if col <= -1 || col >= len(lines) || row <= -1 || row >= len(lines) {
		return Coord{}, fmt.Errorf("")
	}
	height, _ := strconv.Atoi(string(lines[row][col]))
	coord := Coord{row, col, height}
	return coord, nil
}

// Get all neighbors of a coord that have a height of coord.height + 1
// Add found 9s reached by single steps to the `found` set.
// Exclude duplicate trails to same trail apex for part 1.
func findTrails(lines []string, coord Coord, found CoordSet) int {
	count := 0
	for _, d := range deltas {
		neighbor, err := makecoord(lines, coord.row+d.row, coord.col+d.col)
		if err != nil || neighbor.height != coord.height+1 {
			continue
		} else if neighbor.height == 9 {
			if PART == 2 || found.has(neighbor) {
				count += 1
			}
			found.add(neighbor)
		} else {
			count += findTrails(lines, neighbor, found)
		}
	}
	return count
}

func solver(lines []string) string {
	count := 0
	for row := range lines {
		for col := range lines[row] {
			coord, _ := makecoord(lines, row, col)
			if coord.height == 0 {
				found := CoordSet{}
				count += findTrails(lines, coord, found)
			}
		}
	}
	return fmt.Sprintf("%d", count)
}

func PartOne(lines []string) string {
	PART = 1
	return solver(lines)
}

func PartTwo(lines []string) string {
	PART = 2
	return solver(lines)
}
