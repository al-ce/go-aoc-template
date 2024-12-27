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
	return "TODO"
}
