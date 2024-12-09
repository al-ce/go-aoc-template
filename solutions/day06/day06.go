package day06

import (
	"fmt"
	"os"
	"os/exec"
	"slices"
	"time"
)

// import (
//     h "go-aoc-template/internal/helpers"
// )

const GUARD = "^"

// Find the guard's starting position
func findStart(lines []string) Coord {
	for i, row := range lines {
		for j, char := range row {
			if string(char) == GUARD {
				return Coord{i, j}
			}
		}
	}
	return Coord{-1, -1}
}

// Initialize the lab
func makeLab(lines []string) [][]rune {
	lab := make([][]rune, len(lines))
	for i, row := range lines {
		lab[i] = make([]rune, len(lab))
		for j, char := range row {
			lab[i][j] = char
		}
	}
	return lab
}

func printLab(lab [][]rune) {
	print("\033[H\033[2J")
	var cmd *exec.Cmd
	cmd = exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	println()
	for _, row := range lab {
		fmt.Println(string(row))
	}
	time.Sleep(50 * time.Millisecond)
}

type Coord struct {
	y int
	x int
}

type Guard struct {
	loc   Coord
	delta Coord
	char  rune
	lab   [][]rune
}

var directions = []Coord{
	{-1, 0}, {0, 1}, {1, 0}, {0, -1},
}

// Turn the guard 90 degrees
func (g *Guard) turnRight() {
	chars := []rune{'⮙', '⮚', '⮛', '⮘'}
	g.char = chars[(slices.Index(chars, g.char)+1)%len(chars)]
	g.delta = directions[(slices.Index(directions, g.delta)+1)%len(directions)]
}

// Mark the current location as visited
func (g *Guard) visit() {
	g.lab[g.loc.y][g.loc.x] = 'x'
}

// Mark the current location as visited and move to the next
func (g *Guard) move() {
	g.loc.x += g.delta.x
	g.loc.y += g.delta.y
	g.lab[g.loc.y][g.loc.x] = g.char
}

// Look at the next block and return the rune
func (g *Guard) lookAhead() rune {
	nextX := g.loc.x + g.delta.x
	nextY := g.loc.y + g.delta.y
	return g.lab[nextY][nextX]
}

// Check if the guard is at the lab edge (will leave on next move)
func (g *Guard) atLabEdge() bool {
	return g.loc.x == 0 || g.loc.x == len(g.lab)-1 ||
		g.loc.y == 0 || g.loc.y == len(g.lab)-1
}

// Create a new guard
func NewGuard(lines []string) Guard {
	start := findStart(lines)
	lab := makeLab(lines)
	g := Guard{start, directions[0], '⮙', lab}
	g.lab[g.loc.y][g.loc.x] = g.char //  mark the initial location
	return g
}

// Get the sum of visited locations
func (g *Guard) sumOfVisits() int {
	visitedCount := 0
	for i, row := range g.lab {
		for j := range row {
			if g.lab[i][j] == 'x' {
				visitedCount++
			}
		}
	}
	return visitedCount
}

func PartOne(lines []string) string {
	g := NewGuard(lines)
	for !g.atLabEdge() {
		// printLab(guard.lab)
		g.visit()
		if g.lookAhead() == '#' {
			g.turnRight()
		}
		g.move()
	}
	g.visit() // mark the visit at the lab edge

	// printLab(guard.lab)
	// println()
	// fmt.Println(guard.loc)

	return fmt.Sprintf("%d", g.sumOfVisits())
}

func PartTwo(lines []string) string {
	return "TODO"
}
