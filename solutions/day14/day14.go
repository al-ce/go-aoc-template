package day14

import (
	"fmt"
	"regexp"
	"strconv"

	h "go-aoc-template/internal/helpers"
)

type Coord struct {
	x int
	y int
}
type Robot struct {
	p Coord
	v Coord
}

const (
	WIDTH  = 101
	HEIGHT = 103
	TIME   = 100
)

func makeRobot(robot string) Robot {
	pattern := regexp.MustCompile(`p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)`)
	matches := pattern.FindStringSubmatch(robot)
	px, _ := strconv.Atoi(matches[1])
	py, _ := strconv.Atoi(matches[2])
	vx, _ := strconv.Atoi(matches[3])
	vy, _ := strconv.Atoi(matches[4])
	return Robot{
		Coord{px, py},
		Coord{vx, vy},
	}
}

func PartOne(lines []string) string {
	q := []int{0, 0, 0, 0}
	robots := make(map[Coord]struct{})
	for _, r := range lines {
		robot := makeRobot(r)

		newX := h.Mod(robot.p.x+robot.v.x*TIME, WIDTH)
		newY := h.Mod(robot.p.y+robot.v.y*TIME, HEIGHT)
		robots[Coord{newX, newY}] = struct{}{}

		if newX < WIDTH/2 {
			if newY < HEIGHT/2 {
				q[0]++
			} else if newY > HEIGHT/2 {
				q[1]++
			}
		} else if newX > WIDTH/2 {
			if newY < HEIGHT/2 {
				q[2]++
			} else if newY > HEIGHT/2 {
				q[3]++
			}
		}

	}

	return fmt.Sprintf("%d", h.SliceProd(q))
}

func PartTwo(lines []string) string {
	robots := make(map[Coord]struct{})
	for _, r := range lines {
		robot := makeRobot(r)

		// Found by trial and error. Saw that 33+101n (width) 87+101n (height)
		// gave clustered robots, showing their correct x and y positions
		// respectively. Then brute force checked every number in range 101*103
		// checking for a sequence of 10 robots in a row (upper frame of the
		// tree pattern), printing n along the way.
		newX := h.Mod(robot.p.x+robot.v.x*(7709), WIDTH)
		newY := h.Mod(robot.p.y+robot.v.y*(7709), HEIGHT)
		robots[Coord{newX, newY}] = struct{}{}
	}

	for i := range HEIGHT {
		for j := range WIDTH {
			if _, exists := robots[Coord{j, i}]; exists {
				if (i < 80 && i > 55) && (j < 65 && j > 40) {
					fmt.Print("\033[32m🌲\033[0m")
				} else {
					fmt.Print("\033[31m⨻ \033[0m")
				}
			} else {
				fmt.Print("..")
			}
		}
		println()
	}

	return fmt.Sprintf("%d", 0)
}
