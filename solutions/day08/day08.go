package day08

import "fmt"

// import (
//     h "go-aoc-template/internal/helpers"
// )

type Coord struct {
	y int
	x int
}

type (
	AntennaeTypes map[byte]Antennae // Map of antennae by character, holding all their coords
	Antennae      []Coord           // All the coords of antennae by character, holding their siblings
)

// Check that antinode is not off the map
func validateAntinode(antinode Coord, rowsMax, colsMax int) bool {
	return antinode.y >= 0 && antinode.y <= rowsMax &&
		antinode.x >= 0 && antinode.x <= colsMax
}

func PartOne(lines []string) string {
	antennae := AntennaeTypes{}
	antinodes := map[Coord]byte{}
	for i, row := range lines {
		for j := range row {
			char := lines[i][j]

			// Skip if this location does not have an antenna
			if char == '.' {
				continue
			}

			this := Coord{i, j}

			// Create slice of antenna type ('0', 'A', etc.) for new types
			antenna, exists := antennae[char]
			if !exists {
				antennae[char] = Antennae{}
			}

			antennae[char] = append(antennae[char], this)

			// Find the antinodes of this antenna and all preceding antennae of
			// the same type
			for _, sibling := range antenna {
				if sibling == this {
					continue
				}

				// Calculate the slope of this antenna and its sibling
				xDiff := this.x - sibling.x
				yDiff := this.y - sibling.y

				antinodeOne := Coord{
					y: this.y + yDiff,
					x: this.x + xDiff,
				}
				antinodeTwo := Coord{
					y: sibling.y - yDiff,
					x: sibling.x - xDiff,
				}

				// Place in map to create a set (value does not matter)
				if validateAntinode(antinodeOne, len(lines)-1, len(row)-1) {
					antinodes[antinodeOne] = char
				}
				if validateAntinode(antinodeTwo, len(lines)-1, len(row)-1) {
					antinodes[antinodeTwo] = char
				}
			}
		}
	}

	return fmt.Sprint(len(antinodes))
}

func PartTwo(lines []string) string {
	antennae := AntennaeTypes{}
	antinodes := map[Coord]byte{}
	for i, row := range lines {
		for j := range row {
			char := lines[i][j]

			// Skip if this location does not have an antenna
			if char == '.' {
				continue
			}

			this := Coord{i, j}

			// Create slice of antenna type ('0', 'A', etc.) for new types
			antenna, exists := antennae[char]
			if !exists {
				antennae[char] = Antennae{}
			}

			antennae[char] = append(antennae[char], this)

			// Find the antinodes of this antenna and all preceding antennae of
			// the same type
			for _, sibling := range antenna {
				if sibling == this {
					continue
				}

				curr := Coord{this.y, this.x}
				next := Coord{sibling.y, sibling.x}
				xDiff := curr.x - next.x
				yDiff := curr.y - next.y

				for {
					antinodes[curr] = char
					antinodes[next] = char
					antinode := Coord{
						y: curr.y + yDiff,
						x: curr.x + xDiff,
					}
					if !validateAntinode(antinode, len(lines)-1, len(row)-1) {
						break
					}
					curr = next
					next = antinode
				}

				curr = Coord{this.y, this.x}
				next = Coord{sibling.y, sibling.x}
				for {
					antinodes[curr] = char
					antinodes[next] = char
					antinode := Coord{
						y: next.y - yDiff,
						x: next.x - xDiff,
					}

					if !validateAntinode(antinode, len(lines)-1, len(row)-1) {
						break
					}
					next = antinode
					curr = next
				}
			}
		}
	}

	return fmt.Sprint(len(antinodes))
}
