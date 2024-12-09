package day06

import (
	"fmt"
	"os"
	"os/exec"
	"slices"
	"sync"
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
	time.Sleep(20 * time.Millisecond)
}

type Coord struct {
	y int
	x int
}

type Guard struct {
	loc     Coord
	delta   Coord
	char    rune
	lab     [][]rune
	visited map[Coord]Coord
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
	g.visited[g.loc] = g.delta
}

// Move to the next location
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
	visited := map[Coord]Coord{}
	g := Guard{start, directions[0], '⮙', lab, visited}
	g.lab[g.loc.y][g.loc.x] = g.char //  mark the initial location
	return g
}

// Get the sum of visited locations
func (g *Guard) sumOfVisits() int {
	visitedCount := 0
	for range g.visited {
		visitedCount++
	}
	return visitedCount
}

// Move the guard along its path until it exits the lab or gets stuck in a loop
func (g *Guard) patrol() int {
	for !g.atLabEdge() {
		if dir, exists := g.visited[g.loc]; exists && dir == g.delta {
			return 1
		}
		// printLab(g.lab)
		g.visit()
		if g.lookAhead() == '#' {
			g.turnRight()
		}
		g.move()
	}
	g.visit() // mark edge visit
	return 0
}

func PartOne(lines []string) string {
	g := NewGuard(lines)
	g.patrol()

	// printLab(guard.lab)
	// println()
	// fmt.Println(guard.loc)

	return fmt.Sprintf("%d", g.sumOfVisits())
}

func PartTwo(lines []string) string {
	// Patrol once to get visited locations
	g_prime := NewGuard(lines)
	g_prime.patrol()

	results := make(chan int, len(g_prime.visited))

	var wg sync.WaitGroup

	start := findStart(lines)

	// Launch goroutine for each obstacle placement along guard's path
	for v, location := range g_prime.visited {
		if location == start { // don't place obstacle ON the guard
			continue
		}
		i, j := v.y, v.x
		wg.Add(1)
		go func(i, j int) {
			defer wg.Done()

			g := NewGuard(lines)
			g.lab[i][j] = '#'
			count := g.patrol()

			results <- count
		}(i, j)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	totalCount := 0
	for count := range results {
		totalCount += count
	}

	return fmt.Sprint(totalCount)
}
